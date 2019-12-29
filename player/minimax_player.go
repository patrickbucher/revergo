package player

import (
	"math"
	"revergo/board"
)

// MinimaxPlayer is a player that tries to optimize the outcome for the whole
// game by applying the Minimax algorithm.
type MinimaxPlayer struct {
	state board.State
	name  string
	depth int
}

// DefaultDepth is the default maximum recursion depth for the Minimax
// algorithm.
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

const (
	worstPossible = math.MinInt32
	bestPossible  = math.MaxInt32
)

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
	depth := int(math.Min(float64(b.TurnsLeft()), float64(p.depth)))
	outcome := minimax(b, p.state, board.Other(p.state), depth, worstPossible, bestPossible)
	return outcome.move
}

type outcome struct {
	diff int
	move *board.Move
}

func minimax(b *board.Board, self, other board.State, depth, ourBest, oppBest int) outcome {
	validMoves := b.ValidMoves(self)
	bestOutcome := outcome{(board.Dimension * board.Dimension) * -1, nil}
	alpha := worstPossible
	beta := bestPossible
	for _, move := range validMoves {
		result, err := b.Play(move, self)
		if err != nil {
			panic("applied invalid move")
		}
		m := &board.Move{Row: move.Row, Col: move.Col}
		diff := 0
		if depth > 1 {
			// players switched: other <-> self, alpha <-> beta
			minimaxResult := minimax(result, other, self, depth-1, -beta, -alpha)
			// invert outcome: our move with opponent's weakest counter move
			diff = minimaxResult.diff * -1
		} else {
			diff, _ = b.Outcome(self, other)
		}
		if diff > alpha {
			alpha = diff
			bestOutcome = outcome{diff, m}
		}
		if -diff < beta {
			beta = -diff
		}
		if ourBest != worstPossible && alpha > ourBest {
			return bestOutcome
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
