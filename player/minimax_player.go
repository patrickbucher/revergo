package player

import (
	"revergo/board"
)

// MinimaxPlayer is a player that tries to optimize the outcome for the whole
// game by applying the Minimax algorithm.
type MinimaxPlayer struct {
	state board.State
	name  string
	depth int
}

const DefaultDepth = 3

// NewMinimaxPlayerSpawnFunc creates a spawn func for the minimax player. Since
// the tournament API isn't aware of a depth parameter, this function creates a
// closure that satisfies the interface. If a depth < 1 is given, depth falls
// back to DefaultDepth.
func NewMinimaxPlayerSpawnFunc(depth int) func(board.State, string) *Player {
	if depth < 1 {
		depth = DefaultDepth
	}
	spawnFunc := func(state board.State, name string) *Player {
		minimaxPlayer := MinimaxPlayer{state, name, depth}
		player := Player(&minimaxPlayer)
		return &player
	}
	return spawnFunc
}

// NewMinimaxPlayer creates a new minimax player with DefaultDepth.
func NewMinimaxPlayer(state board.State, name string) *Player {
	minimaxPlayer := MinimaxPlayer{state, name, DefaultDepth}
	player := Player(&minimaxPlayer)
	return &player
}

// Play applies the Minimax algorithm to find the move that yields the best
// outcome over the next n moves, with n determined by the depth parameter of
// the constructor function.
func (p *MinimaxPlayer) Play(b *board.Board) *board.Move {
	candidates := b.ValidMoves(p.state)
	if len(candidates) == 0 {
		return nil
	}
	if len(candidates) == 1 {
		return candidates[0]
	}
	outcome := minimax(b, p.state, board.Other(p.state), p.depth)
	return outcome.move
}

type outcome struct {
	diff int
	move *board.Move
}

func minimax(b *board.Board, self, other board.State, depth int) outcome {
	validMoves := b.ValidMoves(self)
	outcomes := make([]outcome, 0)
	for _, move := range validMoves {
		result, err := b.Play(move, self)
		if err != nil {
			panic("applied invalid move")
		}
		if depth > 1 {
			// players switched: other <-> self
			result := minimax(result, other, self, depth-1)
			outcomes = append(outcomes, outcome{result.diff, move.Copy()})
		} else {
			diff, _ := b.Outcome(self, other)
			outcomes = append(outcomes, outcome{diff, move.Copy()})
		}
	}
	bestOutcome := outcome{(board.Dimension * board.Dimension) * -1, nil}
	for _, result := range outcomes {
		if result.diff > bestOutcome.diff {
			bestOutcome = result
		}
	}
	return bestOutcome
}

// State returns the player's state (Black or White).
func (p *MinimaxPlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *MinimaxPlayer) Name() string {
	return p.name
}
