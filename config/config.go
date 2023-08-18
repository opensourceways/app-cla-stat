package config

import (
	"github.com/opensourceways/app-cla-stat/accesstoken/domain"
	"github.com/opensourceways/app-cla-stat/accesstoken/infrastructure/symmetricencryptionimpl"
	"github.com/opensourceways/app-cla-stat/common/infrastructure/mongodb"
	"github.com/opensourceways/app-cla-stat/signing/domain/dp"
	"github.com/opensourceways/app-cla-stat/signing/infrastructure/repositoryimpl"
	"github.com/opensourceways/app-cla-stat/util"
)

func Load(path string) (cfg Config, err error) {
	if err = util.LoadFromYaml(path, &cfg); err != nil {
		return
	}

	cfg.setDefault()

	err = cfg.validate()

	return
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

type domainConfig struct {
	AccessToken     domain.Config `json:"access_token"`
	DomainPrimitive dp.Config     `json:"domain_primitive"`
}

type mongodbConfig struct {
	DB mongodb.Config `json:"db" required:"true"`

	repositoryimpl.Config
}

type Config struct {
	Domain    domainConfig                   `json:"domain"          required:"true"`
	Mongodb   mongodbConfig                  `json:"mongodb"         required:"true"`
	Symmetric symmetricencryptionimpl.Config `json:"symmetric"       required:"true"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Domain.AccessToken,
		&cfg.Domain.DomainPrimitive,
		&cfg.Mongodb.DB,
		&cfg.Mongodb.Config,
		&cfg.Symmetric,
	}
}

func (cfg *Config) setDefault() {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) validate() error {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
