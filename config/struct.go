package config

type Config struct {
	Host        string `json:"host" yaml:"host,omitempty"`
	Port        int    `json:"port" yaml:"port,omitempty"`
	Environment string `json:"env" yaml:"env,omitempty"`
	SecretKey   string `json:"secret" yaml:"secret"`
	DB_ADAPTER  string `json:"db_adapter" yaml:"db_adapter"`
	DB_CON_STR  string `json:"db_con_str" yaml:"db_con_str"`
}
