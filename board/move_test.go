package board

import "testing"

type moveShift struct {
	move        *Move
	shift       Shift
	expected    *Move
	expectedErr error
}

var moveShiftTests = []moveShift{
	{&Move{0, 0}, Shifts[SouthEast], &Move{1, 1}, nil},
	{&Move{0, 0}, Shifts[NorthWest], nil, ShiftError},
}

func TestApply(t *testing.T) {
	for _, test := range moveShiftTests {
		got, err := test.move.Apply(&test.shift)
		if got == nil && test.expected != nil {
			t.Errorf("expected result %v, got %v", test.expected, got)
		}
		if got != nil && !got.Equal(test.expected) {
			t.Errorf("expected result %v, got %v", test.expected, got)
		}
		if err != test.expectedErr {
			t.Errorf("expected error %v, got %v", test.expectedErr, err)
		}
	}
}

type movePair struct {
	moveA *Move
	moveB *Move
	equal bool
}

var movePairTests = []movePair{
	{&Move{0, 0}, &Move{0, 0}, true},
	{&Move{0, 0}, &Move{0, 1}, false},
	{&Move{0, 1}, &Move{0, 0}, false},
	{&Move{1, 1}, &Move{1, 1}, true},
	{&Move{0, 1}, &Move{1, 1}, false},
	{&Move{0, 1}, nil, false},
}

func TestEqual(t *testing.T) {
	for _, test := range movePairTests {
		got := test.moveA.Equal(test.moveB)
		if got != test.equal {
			t.Errorf("expected %v.Equal(%v) to be %v, was %v",
				test.moveA, test.moveB, test.equal, got)
		}
	}
}
