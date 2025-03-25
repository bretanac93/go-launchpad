package server

import (
	"fmt"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

func (c Config) Addr() string {
	return fmt.Sprintf("0.0.0.0:%d", c.Port)
}
