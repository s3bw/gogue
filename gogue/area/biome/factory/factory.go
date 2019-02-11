package factory

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/area/biome"
)

// biomeFactories stores an internal mapping between domains
// and their respective factories.
var biomeFactories = make(map[string]BiomeFactory)

// BiomeFactory is the factory interface for creating interfaces.
type BiomeFactory interface {
	Create(parameters map[string]interface{}) (biome.Biome, error)
}

func Register(name string, factory BiomeFactory) {
	biomeFactories[name] = factory
}

func Create(name string, parameters map[string]interface{}) (biome.Biome, error) {
	biomeFactory, ok := biomeFactories[name]
	if !ok {
		return nil, InvalidBiomeError{name}
	}
	return biomeFactory.Create(parameters)
}

type InvalidBiomeError struct {
	Name string
}

func (err InvalidBiomeError) Error() string {
	return fmt.Sprintf("Biome not registered: %s", err.Name)
}
