package player

import (
	"math/rand"
	"revergo/board"
	"time"
)

// RandomPlayer is a player that picks a random valid move.
type RandomPlayer struct {
	state board.State
	name  string
}

// NewRandomPlayer creates a new random player.
func NewRandomPlayer(state board.State, name string) *Player {
	randomPlayer := RandomPlayer{state, name}
	rand.Seed(time.Now().Unix())
	player := Player(&randomPlayer)
	return &player
}

// Play picks a random move from the board and returns it.
func (p *RandomPlayer) Play(board *board.Board) *board.Move {
	moves := board.ValidMoves(p.state)
	if len(moves) == 0 {
		return nil
	}
	pick := rand.Intn(len(moves))
	return moves[pick]
}

// State returns the player's state (Black or White).
func (p *RandomPlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *RandomPlayer) Name() string {
	return p.name
}
