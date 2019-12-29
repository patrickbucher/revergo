package player

import (
	"math/rand"
	"revergo/board"
	"time"
)

// CorneravoidancePlayer is a player that tries to pick the corners if available,
// and avoids the three fields around the corner if possible.
type CorneravoidancePlayer struct {
	state board.State
	name  string
}

// NewCorneravoidancePlayer creates a new corneravoidance player.
func NewCorneravoidancePlayer(state board.State, name string) *Player {
	cornerdefensePlayer := CornerdefensePlayer{state, name}
	rand.Seed(time.Now().Unix())
	player := Player(&cornerdefensePlayer)
	return &player
}

// Play tries to pick a corner field if available. If not, all moves are tried
// and tested, if they are a field next to the corner fields. Those moves are
// eliminated. From the remaining moves, a random move is picked.
func (p *CorneravoidancePlayer) Play(b *board.Board) *board.Move {
	moves := b.ValidMoves(p.state)
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
	candidates := make([]*board.Move, 0)
	for _, move := range moves {
		if isNotAroundCorner(*move) {
			return move
		}
	}
	if len(candidates) > 0 {
		return candidates[rand.Intn(len(candidates))]
	}
	return moves[rand.Intn(len(moves))]
}

const max = board.Dimension - 1

var cornerNeighbours = []board.Move{
	// top left
	board.Move{0, 1},
	board.Move{1, 1},
	board.Move{1, 0},

	// top right
	board.Move{0, max - 1},
	board.Move{1, max - 1},
	board.Move{1, max},

	// bottom left
	board.Move{max - 1, 0},
	board.Move{max - 1, 1},
	board.Move{max, 1},

	// bottom right
	board.Move{max, max - 1},
	board.Move{max - 1, max - 1},
	board.Move{max - 1, max},
}

func isNotAroundCorner(move board.Move) bool {
	for _, cornerNeighbour := range cornerNeighbours {
		if move.Equal(&cornerNeighbour) {
			return true
		}
	}
	return false
}

// State returns the player's state (Black or White).
func (p *CorneravoidancePlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *CorneravoidancePlayer) Name() string {
	return p.name
}
