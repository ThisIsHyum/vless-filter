package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/thisishyum/vless-filter/client"
)

type Server struct {
	addr   string
	client *client.Client
}

func NewServer(client *client.Client, addr string) Server {
	return Server{
		client: client,
		addr:   addr,
	}
}

func (s *Server) Run() error {
	go s.client.Cycle()

	mux := http.NewServeMux()
	mux.HandleFunc("/subs", s.getSubs)
	return http.ListenAndServe(
		s.addr,
		s.logger(mux),
	)
}

func (s *Server) getSubs(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	strMaxLatency := r.URL.Query().Get("max_latency")

	var limit int
	if strLimit != "" {
		l, err := strconv.Atoi(strLimit)
		if err != nil {
			http.Error(w, "limit is not number", http.StatusBadRequest)
			return
		}
		limit = l
	}

	var maxLatency time.Duration
	if strMaxLatency != "" {
		ml, err := time.ParseDuration(strMaxLatency)
		if err != nil {
			http.Error(w, "max_latency format is not valid", http.StatusBadRequest)
			return
		}
		maxLatency = ml
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(s.client.GetFilteredLinks(limit, maxLatency))
}
