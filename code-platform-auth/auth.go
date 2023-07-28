package oauth

import (
	"fmt"

	"github.com/opensourceways/app-cla-stat/oauth2"
)

const AuthApplyToLogin = "login"

// key is the purpose that authorization applies to
var Auth = map[string]*codePlatformAuth{}

func Initialize(cfg *Config) error {
	f := func(purpose string, ac *authConfig) {
		cpa := &codePlatformAuth{
			clients: map[string]AuthInterface{},
		}

		for _, item := range ac.Configs {
			cpa.clients[item.Platform] = &authClient{
				c: oauth2.NewOauth2Client(item.Oauth2Config),
			}
		}

		Auth[purpose] = cpa
	}

	f(AuthApplyToLogin, &cfg.Login)

	return nil
}

type AuthInterface interface {
	PasswordCredentialsToken(username, password string) (string, error)
}

// codePlatformAuth
type codePlatformAuth struct {
	clients map[string]AuthInterface
}

func (this *codePlatformAuth) GetAuthInstance(platform string) (AuthInterface, error) {
	if c, ok := this.clients[platform]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("Failed to get oauth instance: unknown platform: %s", platform)
}

// authClient
type authClient struct {
	c oauth2.Oauth2Interface
}

func (this *authClient) PasswordCredentialsToken(username, password string) (string, error) {
	token, err := this.c.PasswordCredentialsToken(username, password)
	if err != nil {
		return "", fmt.Errorf("Get token failed: %s", err.Error())
	}

	return token.AccessToken, nil
}
