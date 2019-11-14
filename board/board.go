package board

// EmptyBoard creates a new Board with every field set to Empty.
func EmptyBoard() *Board {
	var board Board
	board = make([][]Field, Dimension)
	for r := 0; r < Dimension; r++ {
		board[r] = make([]Field, Dimension)
		for c := 0; c < Dimension; c++ {
			board[r][c] = Empty
		}
	}
	return &board
}

// InitialBoard creates a new Board with InitialPositions to their respective
// values, and all other fields left to Empty.
func InitialBoard() *Board {
	var board = EmptyBoard()
	for field, positions := range InitialPositions {
		for _, position := range positions {
			(*board)[position.Row][position.Col] = field
		}
	}
	return board
}

// Equal tests if the other board has the same dimensions and all the same
// field values at the corresponding positions.
func (b *Board) Equal(other *Board) bool {
	if len(*b) != len(*other) {
		return false
	}
	for r := 0; r < len(*b); r++ {
		if len((*b)[r]) != len((*other)[r]) {
			return false
		}
		for c := 0; c < len(*b); c++ {
			if (*b)[r][c] != (*other)[r][c] {
				return false
			}
		}
	}
	return true
}
