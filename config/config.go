package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	HTTPPort string `envconfig:"HTTP_PORT" default:"8085"`

	DBType string `envconfig:"DB_TYPE" default:"postgres"`

	DBHost         string `envconfig:"DB_HOST" default:"aws-0-ap-southeast-1.pooler.supabase.com"`
	DBPort         string `envconfig:"DB_PORT" default:"6543"`
	DBUsername     string `envconfig:"DB_USERNAME" default:"postgres.ujzhcinmidvizxywyqsx"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"Barcelona2015##"`
	DBName         string `envconfig:"DB_NAME" default:"postgres"`
	DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS" default:"50"`
	DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS" default:"400"`
	DBMaxIdleTime  int    `envconfig:"DB_MAX_IDLE_TIME" default:"10"`
	DBMaxLifetime  int    `envconfig:"DB_MAX_LIFETIME" default:"60"`

	Kafka string `envconfig:"KAFKA" default:""`

	AppName    string `envconfig:"APP_NAME" default:"be-inventory"`
	KodeCabang string `envconfig:"KODE_CABANG" default:""`

	AuthServiceURL string `envconfig:"AUTH_SERVICE_URL" default:"http://localhost:8081"`
}

func Get() *Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return &cfg
}
