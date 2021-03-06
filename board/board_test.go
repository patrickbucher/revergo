package board

import (
	"strings"
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

func TestCopyBoard(t *testing.T) {
	original := InitialBoard()
	copied := original.Copy()
	if original == copied {
		t.Error("copy is not supposed to refer to the same memory area")
	}
	if !original.Equal(copied) {
		t.Error("copy is supposed to be equal to original")
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

func TestPlayInalidMoves(t *testing.T) {
	initial := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 1, 2, 2, 0, 0, 0},
		{0, 0, 1, 1, 1, 2, 0, 0},
		{0, 0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	illegalBlackMove := &Move{6, 6}
	_, err := initial.Play(illegalBlackMove, Black)
	if err != ErrorInvalidMove {
		t.Errorf("play %v move %v to \n%v\n should cause error",
			Black, illegalBlackMove, initial)
	}

	illegalWhiteMove := &Move{1, 4}
	_, err = initial.Play(illegalWhiteMove, White)
	if err != ErrorInvalidMove {
		t.Errorf("play %v move %v to \n%v\n should cause error",
			White, illegalWhiteMove, initial)
	}
}

func TestPlayValidMoves(t *testing.T) {
	initial := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 1, 2, 2, 0, 0, 0},
		{0, 0, 1, 1, 1, 2, 0, 0},
		{0, 0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	afterBlackMove := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 2, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 2, 0, 0},
		{0, 0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	blackMove := &Move{1, 4}
	got, err := initial.Play(blackMove, Black)
	if err != nil {
		t.Errorf("applying %d move %v to \n%v\n caused an unexpected error %v",
			Black, blackMove, initial, err)
	}
	if !got.Equal(afterBlackMove) {
		t.Errorf("applying %d move %v to \n%v\n expected to be \n%v\n but was \n%v",
			Black, blackMove, initial, afterBlackMove, got)
	}
}

func TestPlayBigMove(t *testing.T) {
	fictional := &Board{
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 0, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 2},
		{2, 1, 1, 1, 1, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
	}
	expected := &Board{
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 2, 2, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 2, 2, 2, 1, 1, 2},
		{2, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 1, 2, 1, 1, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
	}
	move := &Move{3, 3}
	got, _ := fictional.Play(move, White)
	if !got.Equal(expected) {
		t.Errorf("applying %d move %v to \n%v\n expected to be \n%v\n but was \n%v",
			White, move, fictional, expected, got)
	}
}

func TestOutcomeInGame(t *testing.T) {
	board := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 1, 2, 2, 0, 0, 0},
		{0, 2, 2, 2, 2, 2, 0, 0},
		{0, 0, 1, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	black := 2
	white := 9
	expectedDiff := black - white
	expectedFinished := false

	gotDiff, gotFinished := board.Outcome(Black, White)
	if gotDiff != expectedDiff {
		t.Errorf("expected a differnce of %d, but got %d", expectedDiff, gotDiff)
	}
	if gotFinished != expectedFinished {
		t.Errorf("expected finished to be %v, but was %v", expectedFinished, gotFinished)
	}
}

func TestOutcomeFinished(t *testing.T) {
	board := &Board{
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 2, 2, 1, 1, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 1, 2, 2, 2, 1, 1, 2},
		{2, 2, 1, 2, 1, 2, 1, 2},
		{2, 1, 1, 2, 1, 1, 2, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
	}
	white := 48
	black := 16
	expectedDiff := white - black
	expectedFinished := true

	gotDiff, gotFinished := board.Outcome(White, Black)
	if gotDiff != expectedDiff {
		t.Errorf("expected a differnce of %d, but got %d", expectedDiff, gotDiff)
	}
	if gotFinished != expectedFinished {
		t.Errorf("expected finished to be %v, but was %v", expectedFinished, gotFinished)
	}
}

func TestRender(t *testing.T) {
	const emptyField = '-'
	const blackPlayer = 'x'
	const whitePlayer = 'o'
	var rowLabels = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	var colLabels = []rune{'1', '2', '3', '4', '5', '6', '7', '8'}
	position := &Board{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	expected := strings.TrimSpace(`
+ 1 2 3 4 5 6 7 8
a - - - - - - - -
b - - - - - - - -
c - - - - - - - -
d - - - o x - - -
e - - - x o - - -
f - - - - - - - -
g - - - - - - - -
h - - - - - - - -`)

	got, err := position.Render(blackPlayer, whitePlayer, emptyField, rowLabels, colLabels)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != expected {
		t.Errorf("expected rendering \n%v\ngot rendering \n%v\n", expected, got)
	}
}
