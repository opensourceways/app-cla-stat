package oauth2

import (
	"context"
	"fmt"

	liboauth2 "golang.org/x/oauth2"
)

type Oauth2Interface interface {
	PasswordCredentialsToken(username, password string) (*liboauth2.Token, error)
}

type client struct {
	cfg *liboauth2.Config
}

func (this *client) PasswordCredentialsToken(username, password string) (*liboauth2.Token, error) {
	token, err := this.cfg.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve token: %v", err)
	}
	return token, nil
}

// Oauth2Config
type Oauth2Config struct {
	ClientID     string   `json:"client_id" required:"true"`
	ClientSecret string   `json:"client_secret" required:"true"`
	AuthURL      string   `json:"auth_url" required:"true"`
	TokenURL     string   `json:"token_url" required:"true"`
	RedirectURL  string   `json:"redirect_url" required:"true"`
	Scope        []string `json:"scope" required:"true"`
}

func NewOauth2Client(cfg Oauth2Config) Oauth2Interface {
	return &client{
		cfg: buildOauth2Config(cfg),
	}
}

func buildOauth2Config(cfg Oauth2Config) *liboauth2.Config {
	return &liboauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scope,
		Endpoint: liboauth2.Endpoint{
			AuthURL:  cfg.AuthURL,
			TokenURL: cfg.TokenURL,
		},
		RedirectURL: cfg.RedirectURL,
	}
}
