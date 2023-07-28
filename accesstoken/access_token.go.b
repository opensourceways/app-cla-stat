package accesstoken

import (
	"github.com/opensourceways/app-cla-stat/accesstoken/domain"
	"github.com/opensourceways/app-cla-stat/accesstoken/domain/service"
	"github.com/opensourceways/app-cla-stat/models"
)

func NewAccessTokenAdapter(
	s service.AccessTokenService,
) *accessTokenAdatper {
	return &accessTokenAdatper{s: s}
}

type accessTokenAdatper struct {
	s service.AccessTokenService
}

func (adapter *accessTokenAdatper) Add(payload []byte) (string, models.IModelError) {
	v, err := adapter.s.Add(payload)
	if err != nil {
		return "", models.NewModelError(models.ErrSystemError, err)
	}

	return v, nil
}

func (adapter *accessTokenAdatper) Validate(old string) (
	payload []byte, merr models.IModelError,
) {
	payload, err := adapter.s.Validate(old)
	if err != nil {
		v := ""
		if code, ok := err.(errorCode); ok {
			v = code.ErrorCode()
		}

		switch v {
		case domain.ErrorCodeAccessTokenInvalid:
			merr = models.NewModelError(models.ErrInvalidToken, err)

		case domain.ErrorCodeAccessTokenExpiry:
			merr = models.NewModelError(models.ErrExpiredToken, err)

		default:
			merr = models.NewModelError(models.ErrSystemError, err)
		}
	}

	return
}

type errorCode interface {
	ErrorCode() string
}
