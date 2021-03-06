package player

import (
	"math/rand"
	"revergo/board"
	"time"
)

// CornerPlayer is a player that tries to pick the corners if available, and
// otherwise picks a random move.
type CornerPlayer struct {
	state board.State
	name  string
}

// NewCornerPlayer creates a new corner player.
func NewCornerPlayer(state board.State, name string) *Player {
	cornerPlayer := CornerPlayer{state, name}
	rand.Seed(time.Now().Unix())
	player := Player(&cornerPlayer)
	return &player
}

// Play tries to pick a corner field if available, and otherwise just picks a random move.
func (p *CornerPlayer) Play(board *board.Board) *board.Move {
	moves := board.ValidMoves(p.state)
	if len(moves) == 0 {
		return nil
	}
	if len(moves) == 1 {
		return moves[0]
	}
	for _, move := range moves {
		for _, cornerMove := range Corners {
			if move.Equal(cornerMove) {
				return move
			}
		}
	}
	pick := rand.Intn(len(moves))
	return moves[pick]
}

// State returns the player's state (Black or White).
func (p *CornerPlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *CornerPlayer) Name() string {
	return p.name
}
