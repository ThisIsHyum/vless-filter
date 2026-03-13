package config

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Host, Port        string
	Interval, Timeout time.Duration
	Workers           int
	SubUrls           []string
}

func New() Config {
	hostFlag := flag.String("host", "127.0.0.1", "IP or Domain of server")
	portFlag := flag.Int("port", 80, "Port of server")
	intervalFlag := flag.Duration("interval", 30*time.Minute, "Interval of subscriptions updates")
	timeoutFlag := flag.Duration("timeout", 3*time.Second, "Timeout per node")
	workersFlag := flag.Int("workers", 200, "Number of concurrent workers")
	subsPathFlag := flag.String("subs", "sub_urls.txt", "Path to sub_urls.txt")

	flag.Parse()

	host := *hostFlag
	port := *portFlag
	interval := *intervalFlag
	timeout := *timeoutFlag
	workers := *workersFlag
	subsPath := *subsPathFlag

	if envHost := os.Getenv("HOST"); envHost != "" {
		host = envHost
	}
	if envPort := os.Getenv("PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}
	if envInterval := os.Getenv("INTERVAL"); envInterval != "" {
		if d, err := time.ParseDuration(envInterval); err == nil {
			interval = d
		}
	}
	if envTimeout := os.Getenv("TIMEOUT"); envTimeout != "" {
		if d, err := time.ParseDuration(envTimeout); err == nil {
			timeout = d
		}
	}
	if envWorkers := os.Getenv("WORKERS"); envWorkers != "" {
		if w, err := strconv.Atoi(envWorkers); err == nil {
			workers = w
		}
	}
	if envSubs := os.Getenv("SUBS_PATH"); envSubs != "" {
		subsPath = envSubs
	}

	subs, err := getSubUrls(subsPath)
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		Host:     host,
		Port:     strconv.Itoa(port),
		Interval: interval,
		Timeout:  timeout,
		Workers:  workers,
		SubUrls:  subs,
	}
}

func (c Config) Addr() string {
	return c.Host + ":" + c.Port
}

func getSubUrls(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}
	return urls, scanner.Err()
}
