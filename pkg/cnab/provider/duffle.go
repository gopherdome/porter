package cnabprovider

import (
	"os"

	"github.com/deislabs/porter/pkg/config"

	driver "github.com/deislabs/duffle/pkg/driver"
)

type Duffle struct {
	*config.Config
}

func NewDuffle(c *config.Config) *Duffle {
	return &Duffle{
		Config: c,
	}
}

func (d *Duffle) newDriver(driverName string) (driver.Driver, error) {
	driverImpl, err := driver.Lookup(driverName)
	if err != nil {
		return driverImpl, err
	}

	// Load any driver-specific config out of the environment.
	// TODO: This should be exposed in duffle, taken from cmd/duffle/main.go prepareDriver
	if configurable, ok := driverImpl.(driver.Configurable); ok {
		driverCfg := map[string]string{}
		for env := range configurable.Config() {
			driverCfg[env] = os.Getenv(env)
		}
		configurable.SetConfig(driverCfg)
	}

	return driverImpl, err
}
