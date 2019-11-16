package board

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

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
	otherState := Other(playerState)
	neighbourships := b.adjacentOf(empty, otherState) // other's empty neighbours
	for _, candidate := range neighbourships {
		orig := candidate.origin.Copy()
		shift := candidate.shift
		if containsMove(validMoves, orig) {
			continue
		}
		for p, err := orig.Apply(shift); err != ErrorShift; p, err = p.Apply(shift) {
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

// ErrorInvalidMove is the result of an invalid move played on the board.
var ErrorInvalidMove = errors.New("the move is invalid")

// Play applies a move for the given player and returns a new board with the
// move played. If the move is not valid, ErrorInvalidMove is returned.
func (b *Board) Play(move *Move, player State) (*Board, error) {
	validMoves := b.ValidMoves(player)
	if !containsMove(validMoves, move) {
		return nil, ErrorInvalidMove
	}
	board := b.Copy()
	(*board)[move.Row][move.Col] = player
	opponent := Other(player)
	for _, shift := range Shifts {
		chain := make([]*Move, 0)
		for p, err := move.Apply(&shift); err != ErrorShift; p, err = p.Apply(&shift) {
			if (*board)[p.Row][p.Col] == opponent {
				// opponent's field: mark for take-over
				chain = append(chain, p)
			} else if (*board)[p.Row][p.Col] == player {
				// own field at the end of a row: apply take-over
				for _, captured := range chain {
					(*board)[captured.Row][captured.Col] = player
				}
				break
			} else if (*board)[p.Row][p.Col] == Empty {
				break
			}
		}
	}
	return board, nil
}

// Outcome determines the current outcome of the game as the difference between
// player's and opponent's fields. If the board is full, a boolean indicating
// the finished state of the game will be returned; otherwise, false is
// returned.
func (b *Board) Outcome(player State, opponent State) (int, bool) {
	emptyFields := b.numberOfFields(Empty)
	playerFields := b.numberOfFields(player)
	opponentFields := b.numberOfFields(opponent)
	return (playerFields - opponentFields), emptyFields == 0
}

// TurnsLeft returns the number of turns left, i.e. the number of Empty fields
// on the board.
func (b *Board) TurnsLeft() int {
	return b.numberOfFields(Empty)
}

// Other returns the opponent's state of the given state.
func Other(this State) State {
	if this == Black {
		return White
	}
	if this == White {
		return Black
	}
	return Empty
}

// Render renders the board using the given runes for black/white/empty fields
// and adds the given row and column labels.
func (b *Board) Render(black, white, empty rune, rowLabels, colLabels []rune) (string, error) {
	if len(rowLabels) != Dimension || len(colLabels) != Dimension {
		return "", fmt.Errorf("label slices must have length %d", Dimension)
	}
	buf := bytes.NewBufferString("")
	title := bytes.NewBufferString("+ ")
	for _, col := range colLabels {
		title.WriteRune(col)
		title.WriteRune(' ')
	}
	buf.WriteString(strings.TrimSpace(title.String()))
	buf.WriteRune('\n')
	for r, rowLabel := range rowLabels {
		line := bytes.NewBufferString("")
		line.WriteRune(rowLabel)
		line.WriteRune(' ')
		for c := 0; c < Dimension; c++ {
			val := (*b)[r][c]
			if val == Empty {
				line.WriteRune(empty)
			} else if val == Black {
				line.WriteRune(black)
			} else if val == White {
				line.WriteRune(white)
			}
			line.WriteRune(' ')
		}
		buf.WriteString(strings.TrimSpace(line.String()))
		buf.WriteRune('\n')
	}
	return strings.TrimSpace(buf.String()), nil
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
			if err != nil && err == ErrorShift {
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

func (b *Board) numberOfFields(state State) int {
	n := 0
	for r := 0; r < len(*b); r++ {
		for c := 0; c < len((*b)[r]); c++ {
			if (*b)[r][c] == state {
				n++
			}
		}
	}
	return n
}
