package conf

type DbConf struct {
	Driver       string `mapstructure:"driver"`
	Dsn          string `mapstructure:"dsn"`
	MaxIdleCount int    `mapstructure:"maxIdleCount"`
	MaxOpenCount int    `mapstructure:"maxOpenCount"`
}
