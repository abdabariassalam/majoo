package config

type (
	// Config .
	Config struct {
		// Base
		Base struct {
			Port string `yaml:"port"`
		} `yaml:"base"`

		// Mysql.
		Mysql struct {
			DSN string `yaml:"dsn"`
		} `ymal:"mysql"`

		// JWT.
		JWT struct {
			JwtExpiration int    `yaml:"jwt_expiration"`
			JwtSecret     string `yaml:"jwt_secret"`
		} `ymal:"jwt"`
	}
)
