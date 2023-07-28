package dp

import "strings"

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	MaxLengthOfName     int      `json:"max_length_of_name"        required:"true"`
	MaxLengthOfEmail    int      `json:"max_length_of_email"       required:"true"`
	MaxLengthOfCorpName int      `json:"max_length_of_corp_name"   required:"true"`
	SupportedLanguages  []string `json:"supported_languages"`
	supportedLanguages  map[string]string
}

func (cfg *Config) SetDefault() {
	if len(cfg.SupportedLanguages) == 0 {
		cfg.SupportedLanguages = []string{"Chinese", "English"}
	}
}

func (cfg *Config) Validate() error {
	v := map[string]string{}
	for _, item := range cfg.SupportedLanguages {
		v[strings.ToLower(item)] = item
	}

	cfg.supportedLanguages = v

	return nil
}

func (cfg *Config) getLanguage(v string) string {
	return cfg.supportedLanguages[strings.ToLower(v)]
}
