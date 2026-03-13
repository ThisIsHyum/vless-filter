package config

import (
	"bufio"
	"flag"
	"log"
	"net"
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
	host := flag.String("host", "", "IP or Domain of server")
	port := flag.Int("port", 80, "port of server")
	interval := flag.Duration("interval", 30*time.Minute, "interval of subs updates")
	timeout := flag.Duration("timeout", 3*time.Second, "timeout of one node")
	workers := flag.Int("workers", 200, "amount of workers")
	subsPath := flag.String("subs", "sub_urls.txt", "path to sub_urls.txt")

	flag.Parse()

	subs, err := getSubUrls(*subsPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return Config{
		Host:     *host,
		Port:     strconv.Itoa(*port),
		Interval: *interval, Timeout: *timeout,
		Workers: *workers,
		SubUrls: subs,
	}

}

func (c Config) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
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
