package resourcecfg

import (
	"github.com/tanveerprottoy/jenkins-pipeline/service/internal/server/resource"
)

// Config holds the components of the current package
type Config struct {
	Handler *resource.Handler
	Service *resource.Service
}

// NewConfig initializes a new NewConfig
func NewConfig() *Config {
	c := new(Config)
	c.Service = resource.NewService()
	c.Handler = resource.NewHandler(c.Service)
	return c
}
