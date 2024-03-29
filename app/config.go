package app

import (
	"fmt"

	"github.com/iarie/rechallenge/data"
	"github.com/iarie/rechallenge/internal"
)

const (
	defaultPort    = 8080
	defaultPackerV = "V1"
)

type Config struct {
	Port          int
	Packer        internal.Packer
	InventoryRepo internal.Repository
}

func (ac *Config) Addr() string {
	return fmt.Sprintf(":%v", ac.Port)
}

func NewConfig(opts ...ConfigOption) *Config {
	cfg := &Config{
		Port: defaultPort,
	}

	UsePacker(defaultPackerV)(cfg)

	for _, o := range opts {
		o(cfg)
	}

	return cfg
}

type ConfigOption func(*Config)

func UsePort(p int) ConfigOption {
	return func(c *Config) { c.Port = p }
}

func UsePacker(v string) ConfigOption {
	fn := getPackerFuncByVersion(v)

	return func(c *Config) { c.Packer = internal.PackerFunc(fn) }
}

func UseInventoryRepo(r internal.Repository) ConfigOption {
	return func(c *Config) { c.InventoryRepo = r }
}

func getPackerFuncByVersion(v string) func(int, []data.Package) (data.Order, error) {
	switch v {
	case "V1":
		return internal.PackerV1
	case "V2":
		return internal.PackerV2
	default:
		panic(fmt.Sprintf("Unknown Packer version %v", v))
	}
}
