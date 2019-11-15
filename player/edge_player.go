package player

import (
	"math/rand"
	"revergo/board"
	"time"
)

// EdgePlayer is a player that tries to pick the edge lines if available, and
// otherwise picks a random move.
type EdgePlayer struct {
	state board.State
	name  string
}

// NewEdgePlayer creates a new edge player.
func NewEdgePlayer(state board.State, name string) *Player {
	edgePlayer := EdgePlayer{state, name}
	rand.Seed(time.Now().Unix())
	player := Player(&edgePlayer)
	return &player
}

// Play tries to pick an edge field if available, and otherwise just picks a random move.
func (p *EdgePlayer) Play(board *board.Board) *board.Move {
	moves := board.ValidMoves(p.state)
	if len(moves) == 0 {
		return nil
	}
	for _, move := range moves {
		if isEdge(move) {
			return move
		}
	}
	pick := rand.Intn(len(moves))
	return moves[pick]
}

func isEdge(move *board.Move) bool {
	return move.Row == 0 || move.Row == board.Dimension-1 ||
		move.Col == 0 || move.Col == board.Dimension-1
}

// State returns the player's state (Black or White).
func (p *EdgePlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *EdgePlayer) Name() string {
	return p.name
}
