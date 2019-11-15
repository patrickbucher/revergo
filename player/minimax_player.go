package player

import (
	"math"
	"revergo/board"
	"sync"
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
	depth := int(math.Min(float64(b.TurnsLeft()), float64(p.depth)))
	outcome := minimax(b, p.state, board.Other(p.state), depth)
	return outcome.move
}

type outcome struct {
	diff int
	move *board.Move
}

func minimax(b *board.Board, self, other board.State, depth int) outcome {
	validMoves := b.ValidMoves(self)
	ch := make(chan outcome)
	var wg sync.WaitGroup
	for _, move := range validMoves {
		result, err := b.Play(move, self)
		if err != nil {
			panic("applied invalid move")
		}
		wg.Add(1)
		m := &board.Move{Row: move.Row, Col: move.Col}
		go func() {
			defer wg.Done()
			if depth > 1 {
				// players switched: other <-> self
				result := minimax(result, other, self, depth-1)
				ch <- outcome{result.diff, m}
			} else {
				diff, _ := b.Outcome(self, other)
				ch <- outcome{diff, m}
			}
		}()
	}
	bestOutcome := outcome{(board.Dimension * board.Dimension) * -1, nil}
	go func() {
		for result := range ch {
			if result.diff > bestOutcome.diff {
				bestOutcome = result
			}
		}
	}()
	wg.Wait()
	close(ch)
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
