package tracing

import (
	"github.com/canonical/identity-platform-login-ui/internal/logging"
)

type Config struct {
	JaegerEndpoint string
	Logger         logging.LoggerInterface
}

func NewConfig(endpoint string, logger logging.LoggerInterface) *Config {
	c := new(Config)

	c.JaegerEndpoint = endpoint
	c.Logger = logger

	return c
}
