package main

import (
	"log"

	"github.com/thisishyum/vless-filter/client"
	"github.com/thisishyum/vless-filter/config"
	"github.com/thisishyum/vless-filter/server"
)

func main() {
	cfg := config.New()
	cl := client.New(cfg.Interval, cfg.Timeout, cfg.Workers, cfg.SubUrls)
	srv := server.NewServer(cl, cfg.Addr())
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
