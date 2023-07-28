package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/opensourceways/app-cla-stat/accesstoken/domain"
	"github.com/opensourceways/app-cla-stat/accesstoken/domain/symmetricencryption"
	"github.com/opensourceways/app-cla-stat/staterror"
)

var invalidToken = staterror.New(staterror.ErrorCodeAccessTokenInvalid)

type AccessTokenService interface {
	Add(payload []byte) (string, error)
	Validate(string) (p []byte, err error)
}

func NewAccessTokenService(
	encrypt symmetricencryption.Encryption,
) AccessTokenService {
	return &accessTokenService{
		encrypt: encrypt,
	}
}

// accessTokenService
type accessTokenService struct {
	encrypt symmetricencryption.Encryption
}

func (s *accessTokenService) Add(payload []byte) (string, error) {
	token := domain.NewAccessToken(payload)

	v, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	bytes, err := s.encrypt.Encrypt(v)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (s *accessTokenService) Validate(old string) ([]byte, error) {
	ds, err := base64.StdEncoding.DecodeString(old)
	if err != nil {
		return nil, invalidToken
	}

	v, err := s.encrypt.Decrypt(ds)
	if err != nil {
		return nil, invalidToken
	}

	var token domain.AccessToken

	if err := json.Unmarshal(v, &token); err != nil {
		return nil, invalidToken
	}

	if err := token.Validate(); err != nil {
		return nil, err
	}

	return token.Payload, nil
}
