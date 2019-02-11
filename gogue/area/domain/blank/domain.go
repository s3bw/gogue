package blank

import (
	areaDomain "github.com/foxyblue/gogue/gogue/area/domain"
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
	Grid            *areaDomain.Grid
	domainCreatures int
	domainRooms     int
	start           *areaDomain.Coord
	end             *areaDomain.Coord
}

func init() {
	factory.Register(domainName, &blankDomainFactory{})
}

// blankDomainFactory implements the factory.domainFactory interface
type blankDomainFactory struct{}

func (factory *blankDomainFactory) Create(parameters map[string]interface{}) (areaDomain.Domain, error) {
	return fromParameters(parameters)
}

type domain struct {
	parameters DomainParameters
}

func fromParameters(parameters map[string]interface{}) (areaDomain.Domain, error) {
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
		start     = &areaDomain.Coord{X: 10, Y: 10}
		end       = &areaDomain.Coord{X: 20, Y: 20}
	)
	var x, y int

	if parameters != nil {
		if startXY, ok := parameters["start"]; ok {
			start = startXY.(*areaDomain.Coord)
		}
		if maxX, ok := parameters["maxX"]; ok {
			x = maxX.(int)
		}
		if maxY, ok := parameters["maxY"]; ok {
			y = maxY.(int)
		}
	}

	grid := areaDomain.NewGrid(x, y)
	params := &DomainParameters{
		domainCreatures: creatures,
		domainRooms:     rooms,
		start:           start,
		end:             end,
		Grid:            grid,
	}
	return params, nil
}

func New(params DomainParameters) areaDomain.Domain {
	// This will succeed once the interface is implemented
	return &domain{parameters: params}
}

func (d *domain) StartLocation() *areaDomain.Coord {
	return d.parameters.start
}

func (d *domain) EndLocation() *areaDomain.Coord {
	return d.parameters.end
}

func (d *domain) Generate() {
}

func (d *domain) Draw() {}
