package vless

import "net/url"

type VlessConfig struct {
	Uuid     string
	Host     string
	Port     string
	Security string
	Sni      string
}

func Parse(link string) (*VlessConfig, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	uuid := u.User.Username()
	host := u.Hostname()
	port := u.Port()

	q := u.Query()
	security := q.Get("security")
	sni := q.Get("sni")

	return &VlessConfig{
		Uuid:     uuid,
		Host:     host,
		Port:     port,
		Security: security,
		Sni:      sni,
	}, nil
}
