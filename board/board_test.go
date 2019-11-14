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
