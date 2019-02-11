package blank

import (
	area "github.com/foxyblue/gogue/gogue/area/biome"
	"github.com/foxyblue/gogue/gogue/area/biome/factory"
)

const (
	biomeName = "blank"

	defaultBiomeCreatures = 3
	defaultBiomeRooms     = 3
)

// BiomeParameters represents all the configuration options
// for the biome.
type BiomeParameters struct {
	Grid           *area.Grid
	biomeCreatures int
	biomeRooms     int
	start          *area.Coord
	end            *area.Coord
}

func init() {
	factory.Register(biomeName, &blankBiomeFactory{})
}

// blankBiomeFactory implements the factory.biomeFactory interface
type blankBiomeFactory struct{}

func (factory *blankBiomeFactory) Create(parameters map[string]interface{}) (area.Biome, error) {
	return fromParameters(parameters)
}

type biome struct {
	parameters BiomeParameters
}

func fromParameters(parameters map[string]interface{}) (area.Biome, error) {
	params, err := fromParametersImpl(parameters)
	if err != nil || params == nil {
		return nil, err
	}
	return New(*params), nil
}

func fromParametersImpl(parameters map[string]interface{}) (*BiomeParameters, error) {
	var (
		creatures = defaultBiomeCreatures
		rooms     = defaultBiomeRooms
		start     = &area.Coord{X: 10, Y: 10}
		end       = &area.Coord{X: 20, Y: 20}
	)
	var x, y int

	if parameters != nil {
		if startXY, ok := parameters["start"]; ok {
			start = startXY.(*area.Coord)
		}
		if maxX, ok := parameters["maxX"]; ok {
			x = maxX.(int)
		}
		if maxY, ok := parameters["maxY"]; ok {
			y = maxY.(int)
		}
	}

	grid := area.NewGrid(x, y)
	params := &BiomeParameters{
		biomeCreatures: creatures,
		biomeRooms:     rooms,
		start:          start,
		end:            end,
		Grid:           grid,
	}
	return params, nil
}

func New(params BiomeParameters) area.Biome {
	// This will succeed once the interface is implemented
	return &biome{parameters: params}
}

func (b *biome) StartLocation() *area.Coord {
	return b.parameters.start
}

func (b *biome) EndLocation() *area.Coord {
	return b.parameters.end
}

func (b *biome) Generate() {
}

func (b *biome) Draw() {}
