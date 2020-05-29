package config

type BasicAuth struct {
	Username string
	Password string
}

type Config struct {
	Auth    BasicAuth
	BaseURL string
}

func New(username string, password string, baseurl string) *Config {
	auth := BasicAuth{username, password}
	c := new(Config)

	c.Auth = auth
	c.BaseURL = baseurl
	return c
}
