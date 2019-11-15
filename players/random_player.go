package player

import (
	"math/rand"
	"revergo/board"
	"time"
)

// RandomPlayer is a player that picks a random valid move.
type RandomPlayer struct {
	state board.State
}

// NewRandomPlayer creates a new random player.
func NewRandomPlayer(state board.State) *RandomPlayer {
	randomPlayer := RandomPlayer{state}
	rand.Seed(time.Now().Unix())
	return &randomPlayer
}

// Play picks a random move from the board and returns it.
func (p *RandomPlayer) Play(board *board.Board) *board.Move {
	moves := board.ValidMoves(p.state)
	pick := rand.Intn(len(moves))
	return moves[pick]
}
