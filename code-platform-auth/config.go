package oauth

import "github.com/opensourceways/app-cla-stat/oauth2"

type Config struct {
	Login authConfig `json:"login" required:"true"`
}

type authConfig struct {
	Configs []platformConfig `json:"platforms" required:"true"`
}

type platformConfig struct {
	Platform string `json:"platform" required:"true"`

	oauth2.Oauth2Config
}
