package config

import (
	"bytes"
	"fmt"
)

type PostgresRW struct {
	Read  Postgres `yaml:"read"`
	Write Postgres `yaml:"write"`
}

type Postgres struct {
	Host                 string `yaml:"host"`
	Port                 int    `yaml:"port"`
	User                 string `yaml:"user"`
	Password             string `yaml:"password"`
	Database             string `yaml:"database"`
	MaxOpenCons          uint   `yaml:"max_open_cons"`
	MaxIdleCons          uint   `yaml:"max_idle_cons"`
	Log                  uint   `yaml:"log"`
	PreferSimpleProtocol bool   `yaml:"prefer_simple_protocol"`
	SSLMode              string `yaml:"ssl_mode"`
	Timezone             string `yaml:"timezone"`
}

func (r *Postgres) DSN() string {
	dsn := bytes.NewBufferString(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		r.Host,
		r.Port,
		r.User,
		r.Password,
		r.Database,
		r.SSLMode,
		r.Timezone,
	))

	return dsn.String()
}
