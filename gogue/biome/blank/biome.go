package blank

import (
	area "github.com/foxyblue/gogue/gogue/biome"
	"github.com/foxyblue/gogue/gogue/biome/factory"
)

const (
	biomeName = "blank"

	defaultBiomeCreatures = 3
	defaultBiomeRooms     = 3
)

// BiomeParameters represents all the configuration options
// for the biome.
type BiomeParameters struct {
	biomeCreatures int
	biomeRooms     int
	start          *area.Coord
	end            *area.Coord
	x              int
	y              int
	maxX           int
	maxY           int
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
	Grid       area.Grid
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

	if startXY, ok := parameters["start"]; ok {
		start = startXY.(*area.Coord)
	}

	params := &BiomeParameters{
		biomeCreatures: creatures,
		biomeRooms:     rooms,
		start:          start,
		end:            end,
		maxX:           parameters["maxX"].(int) - 1,
		maxY:           parameters["maxY"].(int) - 1,
		x:              parameters["x"].(int),
		y:              parameters["y"].(int),
	}
	return params, nil
}

// New returns a constructed biome, if the linter fails it means
// we haven't implemented all the required methods on the biome
func New(params BiomeParameters) area.Biome {
	w := params.maxX - params.x
	h := params.maxY - params.y
	grid := area.NewGrid(params.x, params.y, w, h)
	return &biome{
		parameters: params,
		Grid:       *grid,
	}
}

func (b *biome) Generate() {
	g := b.Grid
	room := []*area.Coord{
		{X: 5, Y: 5},
		{X: 4, Y: 5},
		{X: 3, Y: 5},
		{X: 2, Y: 5},
	}

	for x, row := range g.Tiles {
		for y := range row {
			if area.IsIn(x, y, room) {
				row[y] = area.WallTile(x, y)
			} else {
				row[y] = area.EmptyTile(x, y)
			}
		}
	}
}

func (b *biome) GetGrid() area.Grid {
	return b.Grid
}

func (b *biome) StartLocation() *area.Coord {
	return b.parameters.start
}

func (b *biome) EndLocation() *area.Coord {
	return b.parameters.end
}
