package board

import (
	"errors"
	"fmt"
)

// Move represents a player setting a stone to the given field (row/column).
type Move struct {
	Row int
	Col int
}

// ErrorShift occurs when a Shift cannot be applied to a Move because the
// resulting position would be out of bounds.
var ErrorShift = errors.New("shifted move is out of bounds")

// Apply creates a new move by applying the given shift to the move. A
// ErrorShift is returned if the resulting position would be out of bounds.
func (m *Move) Apply(shift *Shift) (*Move, error) {
	move := Move{m.Row + shift.Row, m.Col + shift.Col}
	if move.Row < 0 || move.Row >= Dimension || move.Col < 0 || move.Col >= Dimension {
		return nil, ErrorShift
	}
	return &move, nil
}

// Equal checks whether the Row and Col fields of both fields are equal.
func (m *Move) Equal(other *Move) bool {
	if other == nil {
		return false
	}
	return m.Row == other.Row && m.Col == other.Col
}

// String returns the string representation of a move as Move(Row, Col).
func (m *Move) String() string {
	return fmt.Sprintf("Move(%d, %d)", m.Row, m.Col)
}

// Copy creates a copy of the given move.
func (m *Move) Copy() *Move {
	copied := Move{m.Row, m.Col}
	return &copied
}
