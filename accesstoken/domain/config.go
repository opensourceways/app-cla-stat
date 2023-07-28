package domain

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	// AccessTokenExpiry is the one in seconds
	AccessTokenExpiry int64 `json:"access_token_expiry"`
}

func (cfg *Config) SetDefault() {
	if cfg.AccessTokenExpiry <= 0 {
		cfg.AccessTokenExpiry = 3600 * 2
	}
}
