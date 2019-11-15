package player

import (
	"log"
	"revergo/board"
)

// MaxroundPlayer is a player that tries to optimize the outcome for the
// current move.
type MaxroundPlayer struct {
	state board.State
	name  string
}

// NewMaxroundPlayer creates a new maxround player.
func NewMaxroundPlayer(state board.State, name string) *Player {
	maxroundPlayer := MaxroundPlayer{state, name}
	player := Player(&maxroundPlayer)
	return &player
}

// Play picks the move that, applied to the current board, yields the best
// improvement in the outcome.
func (p *MaxroundPlayer) Play(b *board.Board) *board.Move {
	candidates := b.ValidMoves(p.state)
	if len(candidates) == 0 {
		return nil
	}
	if len(candidates) == 1 {
		return candidates[0]
	}
	var bestMove *board.Move
	opponent := board.Other(p.state)
	diff, _ := b.Outcome(p.state, opponent)
	for _, candidate := range candidates {
		result, err := b.Play(candidate, p.state)
		if err != nil {
			log.Printf("applying move %v: %v", candidate, err)
			return nil
		}
		resultDiff, _ := result.Outcome(p.state, opponent)
		if resultDiff > diff {
			bestMove = candidate
			diff = resultDiff
		}
	}
	return bestMove
}

// State returns the player's state (Black or White).
func (p *MaxroundPlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *MaxroundPlayer) Name() string {
	return p.name
}
