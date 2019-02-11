package factory

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/area/domain"
)

// domainFactories stores an internal mapping between domains
// and their respective factories.
var domainFactories = make(map[string]DomainFactory)

// DomainFactory is the factory interface for creating interfaces.
type DomainFactory interface {
	Create(parameters map[string]interface{}) (domain.Domain, error)
}

func Register(name string, factory DomainFactory) {
	domainFactories[name] = factory
}

func Create(name string, parameters map[string]interface{}) (domain.Domain, error) {
	domainFactory, ok := domainFactories[name]
	if !ok {
		return nil, InvalidDomainError{name}
	}
	return domainFactory.Create(parameters)
}

type InvalidDomainError struct {
	Name string
}

func (err InvalidDomainError) Error() string {
	return fmt.Sprintf("Domain not registered: %s", err.Name)
}
