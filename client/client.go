package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/thisishyum/vless-filter/vless"
)

type Client struct {
	client            *http.Client
	cached            []Node
	interval, timeout time.Duration
	workers           int
	urls              []string
	mu                sync.RWMutex
}

type Node struct {
	Link    string
	Latency time.Duration
}

func New(interval, timeout time.Duration, workers int, urls []string) *Client {
	return &Client{
		client:   http.DefaultClient,
		interval: interval,
		timeout:  timeout,
		workers:  workers,
		urls:     urls,
	}
}

func (c *Client) getLinks(url string) ([]string, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	links := strings.Split(string(bytes), "\n")
	return links, nil
}

func (c *Client) getLinksWithRetry(url string, attempts int) ([]string, error) {
	var err error
	for range attempts {
		var links []string
		links, err = c.getLinks(url)
		if err == nil {
			return links, nil
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil, err
}

func (c *Client) filter(links []string) (nodes []Node, failed int) {
	connections := []Node{}
	results := make(chan Node)
	var mu sync.Mutex

	var wg sync.WaitGroup
	sem := make(chan struct{}, c.workers)

	for _, link := range links {
		wg.Add(1)
		sem <- struct{}{}

		go func(l string) {
			defer func() { <-sem }()
			defer wg.Done()

			conn, err := c.isSuitable(l)
			if err == nil {
				results <- conn
			} else {
				mu.Lock()
				failed++
				mu.Unlock()
			}
		}(link)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for conn := range results {
		connections = append(connections, conn)
	}

	c.sortByLatency(connections)
	return connections, failed
}

func (c *Client) sortByLatency(connections []Node) {
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Latency < connections[j].Latency
	})
}

func (c *Client) isSuitable(link string) (Node, error) {
	config, err := vless.Parse(link)
	if err != nil {
		return Node{}, err
	}

	start := time.Now()
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: c.timeout},
		"tcp", config.Host+":"+config.Port,
		&tls.Config{ServerName: config.Sni},
	)
	if err != nil {
		return Node{}, err
	}

	latency := time.Since(start)
	conn.Close()

	return Node{
		Link:    link,
		Latency: latency,
	}, nil
}

func (c *Client) Cycle() {
	for {
		slog.Info("cycle iteration started")

		filtered := []Node{}
		var failed int
		for _, url := range c.urls {
			links, err := c.getLinksWithRetry(url, 3)
			if err != nil {
				slog.Error("getLinksWithRetry: ", slog.Any("error", err), slog.String("url", url))
				continue
			}
			nodes, fld := c.filter(links)
			filtered = append(filtered, nodes...)
			failed += fld
		}

		c.mu.Lock()
		c.cached = filtered
		c.mu.Unlock()

		slog.Info("cycle iteration finished",
			slog.Int("valid", len(filtered)),
			slog.Int("failed", failed))
		time.Sleep(c.interval)
	}
}

// GetFilteredLinks returns links filtered by given constraints.
//
// limit specifies the maximum number of links to return; 0 means no limit.
//
// maxLatency specifies the maximum allowed latency; 0 means no limit.
//
// Links with Latency >= maxLatency are excluded.
func (c *Client) GetFilteredLinks(limit int, maxLatency time.Duration) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()

	links := []string{}
	var count int

	for _, conn := range c.cached {
		if maxLatency != 0 && conn.Latency >= maxLatency {
			continue
		}
		links = append(links, conn.Link)
		count++
		if limit != 0 && count >= limit {
			break
		}
	}
	return []byte(strings.Join(links, "\r\n"))
}
