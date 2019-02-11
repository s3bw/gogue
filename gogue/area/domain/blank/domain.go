package blank

import (
	"github.com/foxyblue/gogue/gogue/area/domain/factory"
)

const (
	domainName = "blank"

	defaultDomainCreatures = 3
	defaultDomainRooms     = 3
)

// DomainParameters represents all the configuration options
// for the domain.
type DomainParameters struct {
	domainCreatures int
	domainRooms     int
	start           *area.Coord
	end             *area.Coord
}

func init() {
	factory.Register(domainName, &blankDomainFactory{})
}

// blankDomainFactory implements the factory.domainFactory interface
type blankDomainFactory struct{}

func (factory *blankDomainFactory) Create(parameters map[string]interface{}) (area.Domain, error) {
	return fromParameters(parameters)
}

type domain struct {
	parameters DomainParameters
}

func fromParameters(parameters map[string]interface{}) (area.Domain, error) {
	params, err := fromParametersImpl(parameters)
	if err != nil || params == nil {
		return nil, err
	}
	return New(*params), nil
}

func fromParametersImpl(parameters map[string]interface{}) (*DomainParameters, error) {
	var (
		creatures = defaultDomainCreatures
		rooms     = defaultDomainRooms
		start     = &area.Coord{X: 10, Y: 10}
		end       = &area.Coord{X: 20, Y: 20}
	)

	if parameters != nil {
		if startXY, ok := parameters["start"]; ok {
			start = startXY.(*area.Coord)
		}
	}
	params := &DomainParameters{
		domainCreatures: creatures,
		domainRooms:     rooms,
		start:           start,
		end:             end,
	}
	return params, nil
}

func New(params DomainParameters) area.Domain {
	// This will succeed once the interface is implemented
	return &domain{parameters: params}
}

func (d *domain) StartLocation() *area.Coord {
	return d.parameters.playerStart
}

func (d *domain) EndLocation() *area.Coord {
	return d.parameters.end
}

func (d *domain) Generate(x, y, lvl int) {
}

func (d *domain) Draw() {}
