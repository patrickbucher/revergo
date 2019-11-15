package player

import "revergo/board"

// Player describes a player, which is capable of playing a move to a given board.
type Player interface {
	// Play returns the player's chosen move to the given board.
	Play(board *board.Board) *board.Move
}
