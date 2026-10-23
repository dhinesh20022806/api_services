// Package env is used config environment variables
package env

import (
	"github.com/joeshaw/envdecode"
)

type config struct {
	Database struct {
		URL         string `env:"DATABASE_URL,required"`
		MaxIdleConn int    `env:"DATABASE_MAX_IDLE_CONNS,default=2,strict"`
		MaxOpenConn int    `env:"DATABASE_MAX_OPEN_CONNS,default=2,strict"`
	}
}

// Config is used access env variable
var Config config

func init() {

	Config = config{}

	err := envdecode.Decode(&Config)

	if err != nil {
		panic("env not loaded!")
	}

}
