package board

// Field represents a field value for a player.
type Field int

const (
	// Dimension is the number of rows/colums of the board.
	Dimension int = 8

	// Empty is an field not occupied yet (initial position)
	Empty Field = 0

	// Black is the value for the black player's positions.
	Black Field = 1

	// White is the value for the white player's positions.
	White Field = 2
)

// Direction is a descriptive label for a point of the compass.
type Direction int

// North, NorthEast, etc. are the eight possible compass direction labels.
const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

// Shift describes a shift both in the x axis (column) and y axis (row).
type Shift struct {
	Row int
	Col int
}

// Shifts are the neighbourhood offsets between fields in compass directions.
var Shifts = map[Direction]Shift{
	North:     {-1, 0},
	NorthEast: {-1, 1},
	East:      {0, 1},
	SouthEast: {1, 1},
	South:     {1, 0},
	SouthWest: {1, -1},
	West:      {0, -1},
	NorthWest: {-1, -1},
}

// Board is a two-dimensional array of fields.
type Board [][]Field

// InitialPositions describes the initial positions of both players.
var InitialPositions = map[Field][]Shift{
	Black: []Shift{
		Shift{Dimension / 2, Dimension/2 - 1},
		Shift{Dimension/2 - 1, Dimension / 2},
	},
	White: []Shift{
		Shift{Dimension/2 - 1, Dimension/2 - 1},
		Shift{Dimension / 2, Dimension / 2},
	},
}
