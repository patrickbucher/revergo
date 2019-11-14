package board

import "errors"

// Move represents a player setting a stone to the given field (row/column).
type Move struct {
	Row int
	Col int
}

// ShiftError occurs when a Shift cannot be applied to a Move because the
// resulting position would be out of bounds.
var ShiftError = errors.New("shifted move is out of bounds")

// Apply creates a new move by applying the given shift to the move. A
// ShiftError is returned if the resulting position would be out of bounds.
func (m *Move) Apply(shift *Shift) (*Move, error) {
	move := Move{m.Row + shift.Row, m.Col + shift.Col}
	if move.Row < 0 || move.Row >= Dimension || move.Col < 0 || move.Col >= Dimension {
		return nil, ShiftError
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
