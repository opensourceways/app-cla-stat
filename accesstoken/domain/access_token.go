package domain

import (
	"time"

	"github.com/opensourceways/app-cla-stat/staterror"
)

func Now() int64 {
	return time.Now().Unix()
}

func NewAccessToken(payload []byte) AccessToken {
	return AccessToken{
		Expiry:  config.AccessTokenExpiry + Now(),
		Payload: payload,
	}
}

type AccessToken struct {
	Expiry  int64  `json:"expiry"`
	Payload []byte `json:"payload"`
}

func (at *AccessToken) Validate() error {
	if at.Expiry >= Now() {
		return staterror.New(staterror.ErrorCodeAccessTokenExpiry)
	}

	return nil
}
