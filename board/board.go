package board

// EmptyBoard creates a new Board with every field set to Empty.
func EmptyBoard() *Board {
	var board Board
	board = make([][]State, Dimension)
	for r := 0; r < Dimension; r++ {
		board[r] = make([]State, Dimension)
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

// Copy copies the original board and returns the copy.
func (b *Board) Copy() *Board {
	copied := EmptyBoard()
	for r := 0; r < Dimension; r++ {
		for c := 0; c < Dimension; c++ {
			(*copied)[r][c] = (*b)[r][c]
		}
	}
	return copied
}

// ValidMoves determines all valid moves for the given player field.
func (b *Board) ValidMoves(playerState State) []*Move {
	validMoves := make([]*Move, 0)
	empty := b.emptyFields()
	otherState := other(playerState)
	neighbourships := b.adjacentOf(empty, otherState) // other's empty neighbours
	for _, candidate := range neighbourships {
		orig := candidate.origin.Copy()
		shift := candidate.shift
		if containsMove(validMoves, orig) {
			continue
		}
		for p, err := orig.Apply(shift); err != ShiftError; p, err = p.Apply(shift) {
			// follow empty field in shift direction to opponent until own field found
			pState := (*b)[p.Row][p.Col]
			if pState == playerState {
				validMoves = append(validMoves, orig)
				break
			}
			if pState == Empty {
				// empty field interrupts chain
				break
			}
		}
	}
	return validMoves
}

func containsMove(moves []*Move, move *Move) bool {
	for _, candidate := range moves {
		if candidate.Equal(move) {
			return true
		}
	}
	return false
}

type neighbourship struct {
	origin    *Move
	shift     *Shift
	neighbour *Move
}

func (b *Board) adjacentOf(fields []*Move, withState State) []*neighbourship {
	adjacents := make([]*neighbourship, 0)
	for _, field := range fields {
		for _, shift := range Shifts {
			shifted, err := field.Apply(&shift)
			if err != nil && err == ShiftError {
				// out of bounds, no longer try this shift direction
				continue
			}
			if (*b)[shifted.Row][shifted.Col] == withState {
				newNeighbourship := neighbourship{field, &Shift{shift.Row, shift.Col}, shifted}
				adjacents = append(adjacents, &newNeighbourship)
			}
		}
	}
	return adjacents
}

func (b *Board) emptyFields() []*Move {
	empty := make([]*Move, 0)
	for r := 0; r < len(*b); r++ {
		for c := 0; c < len((*b)[r]); c++ {
			if (*b)[r][c] == Empty {
				empty = append(empty, &Move{r, c})
			}
		}
	}
	return empty
}

func other(this State) State {
	if this == Black {
		return White
	}
	if this == White {
		return Black
	}
	return Empty
}
