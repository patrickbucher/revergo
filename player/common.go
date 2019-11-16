package player

import "revergo/board"

var Corners = []*board.Move{
	&board.Move{Row: 0, Col: 0},
	&board.Move{Row: 0, Col: board.Dimension - 1},
	&board.Move{Row: board.Dimension - 1, Col: 0},
	&board.Move{Row: board.Dimension - 1, Col: board.Dimension - 1},
}
