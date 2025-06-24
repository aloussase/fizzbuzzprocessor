package fizzbuzzprocessor

import (
	"errors"

	"go.opentelemetry.io/collector/component"
)

type Casing string

const (
	Upper Casing = "upper"
	Lower        = "lower"
)

var (
	validCasings = map[Casing]bool{
		Upper: true,
		Lower: true,
	}
)

type Config struct {
	// Casing for the output messages. Valid values are "upper" and "lower"
	casing Casing
}

func createDefaultConfig() component.Config {
	return &Config{
		casing: Lower,
	}
}

func (c *Config) Validate() error {
	if _, ok := validCasings[c.casing]; !ok {
		return errors.New("invalid value for casing: " + string(c.casing))
	}
	return nil
}
