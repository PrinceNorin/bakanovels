package config

type Config struct {
	Host        string     `json:"host" yaml:"host,omitempty"`
	Port        int        `json:"port" yaml:"port,omitempty"`
	Environment string     `json:"env" yaml:"env,omitempty"`
	SecretKey   string     `json:"secret" yaml:"secret"`
	DB_ADAPTER  string     `json:"db_adapter" yaml:"db_adapter"`
	DB_CON_STR  string     `json:"db_con_str" yaml:"db_con_str"`
	I18n        I18nConfig `json:"i18n" yaml:"i18n"`
}

type I18nConfig struct {
	Directory       string `json:"directory" yaml:"directory"`
	DefaultLanguage string `json:"default_language" yaml:"default_language"`
}
