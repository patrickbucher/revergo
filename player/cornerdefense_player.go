package player

import (
	"log"
	"math/rand"
	"revergo/board"
	"time"
)

// CornerdefensePlayer is a player that tries to pick the corners if available,
// and tries to avoid all moves that make it possible for the opponent to grab
// a corner in the next move.
type CornerdefensePlayer struct {
	state board.State
	name  string
}

// NewCornerdefensePlayer creates a new cornerdefense player.
func NewCornerdefensePlayer(state board.State, name string) *Player {
	cornerdefensePlayer := CornerdefensePlayer{state, name}
	rand.Seed(time.Now().Unix())
	player := Player(&cornerdefensePlayer)
	return &player
}

// Play tries to pick a corner field if available. If not, all moves are tried
// and tested, if they allow the opponent to pick a corner in the next move.
// Those moves are eliminated. From the remaining moves, a random move is
// picked.
func (p *CornerdefensePlayer) Play(b *board.Board) *board.Move {
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
	opponent := board.Other(p.state)
	for _, move := range moves {
		result, err := b.Play(move, p.state)
		if err != nil {
			log.Println(err)
			break
		}
		opponentMoves := result.ValidMoves(opponent)
		for _, opponentMove := range opponentMoves {
			capturesCorner := false
			for _, cornerMove := range Corners {
				if opponentMove.Equal(cornerMove) {
					capturesCorner = true
					break
				}
			}
			if !capturesCorner {
				candidates = append(candidates, move.Copy())
			}
		}
	}
	if len(candidates) > 0 {
		return candidates[rand.Intn(len(candidates))]
	}
	return moves[rand.Intn(len(moves))]
}

// State returns the player's state (Black or White).
func (p *CornerdefensePlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *CornerdefensePlayer) Name() string {
	return p.name
}
