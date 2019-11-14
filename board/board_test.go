package board

import (
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	expected := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	got := EmptyBoard()
	if !got.Equal(expected) {
		t.Errorf("expected board\n%v\n got board \n%v\n", expected, got)
	}
}

func TestInitialBoard(t *testing.T) {
	expected := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	got := InitialBoard()
	if !got.Equal(expected) {
		t.Errorf("expected board\n%v\n got board \n%v\n", expected, got)
	}
}

func TestEqualBoard(t *testing.T) {
	boardA := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 1, 1, 2, 0, 0, 0},
		{0, 0, 0, 2, 1, 1, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	boardB := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 2, 2, 2, 2, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	if !boardA.Equal(boardA) {
		t.Error("boardA is supposed to be equal to itself, but is not")
	}
	if !boardB.Equal(boardB) {
		t.Error("boardB is supposed to be equal to itself, but is not")
	}
	if boardA.Equal(boardB) {
		t.Error("boardA is not supposed to be equal to boardB, but was")
	}
	if boardB.Equal(boardA) {
		t.Error("boardB is not supposed to be equal to boardA, but was")
	}

	boardC := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 2, 2, 2, 2, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		// missing row
	}
	boardD := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 2, 2, 2, 2, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0}, // missing col
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	if boardC.Equal(boardB) {
		t.Error("boardB is not supposed to be equal to boardA, but was")
	}
	if boardD.Equal(boardB) {
		t.Error("boardB is not supposed to be equal to boardA, but was")
	}
}

var initialPlayerMoveTests = map[State][]*Move{
	Black: {
		&Move{2, 3},
		&Move{3, 2},
		&Move{4, 5},
		&Move{5, 4},
	},
	White: {
		&Move{2, 4},
		&Move{3, 5},
		&Move{4, 2},
		&Move{5, 3},
	},
}

func TestValidMovesInitialState(t *testing.T) {
	board := InitialBoard()
	testValidMoves(board, t, initialPlayerMoveTests)
}

var inGamePlayerMoveTests = map[State][]*Move{
	Black: {
		&Move{0, 0},
		&Move{1, 0},
		&Move{1, 3},
		&Move{2, 3},
	},
	White: {
		&Move{0, 2},
		&Move{1, 3},
		&Move{3, 5},
		&Move{4, 0},
		&Move{4, 2},
		&Move{5, 3},
		&Move{5, 5},
	},
}

func TestValidMovesInGameState(t *testing.T) {
	board := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 1, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0},
		{1, 1, 1, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	testValidMoves(board, t, inGamePlayerMoveTests)
}

func testValidMoves(board *Board, t *testing.T, tests map[State][]*Move) {
	for player, expected := range tests {
		got := board.ValidMoves(player)
		if !sameMoves(got, expected) {
			t.Errorf("for player %v expected valid moves %v but got %v",
				player, expected, got)
		}
	}
}

func sameMoves(a, b []*Move) bool {
	if len(a) != len(b) {
		return false
	}
	for _, moveA := range a {
		foundEqual := false
		for _, moveB := range b {
			if moveB.Equal(moveA) {
				foundEqual = true
				break
			}
		}
		if !foundEqual {
			return false
		}
	}
	return true
}
