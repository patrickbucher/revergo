package game

import (
	"log"
	"revergo/board"
	player "revergo/player"
)

// Game represents a game played between two players.
type Game struct {
	playerBlack *player.Player
	playerWhite *player.Player
	board       *board.Board
}

// NewGame creates a new game with the two given players playerBlack and
// playerWhite playing.
func NewGame(playerBlack, playerWhite *player.Player) *Game {
	if (*playerBlack).State() == (*playerWhite).State() {
		panic("players must have different state values")
	}
	board := board.InitialBoard()
	game := Game{playerBlack, playerWhite, board}
	return &game
}

// Result represents the result of a game, with a Winner indicated by
// board.State (Empy for a tie), and a Difference calculated as:
// (black fields) - (white fields)
type Result struct {
	Winner     board.State
	Difference int
}

// Play lets the two players taking turns and modifies the board's state
// accordingly until the game is finished. If the game is stuck, i.e. both
// player were unable to pick a valid move in two successive rounds, the game
// ends prematurely, with the winner being the player with more fields
// captured. If all moves can be played to the end, a result indicating the
// winner and the difference is returned.
func (g *Game) Play() *Result {
	var result Result
	currentPlayer := g.playerWhite
	finished := false
	stuckCount := 0
	for !finished {
		if currentPlayer == g.playerWhite {
			currentPlayer = g.playerBlack
		} else {
			currentPlayer = g.playerWhite
		}
		move := (*currentPlayer).Play(g.board.Copy())
		if move == nil {
			stuckCount++
			if stuckCount == 2 {
				log.Println("game is stuck")
				diff, _ := g.board.Outcome(board.Black, board.White)
				winner := board.Empty
				if diff > 0 {
					winner = board.Black
				} else if diff < 0 {
					winner = board.White
				}
				result := Result{winner, diff}
				return &result
			}
		}
		newBoard, err := g.board.Play(move, (*currentPlayer).State())
		if err == board.ErrorInvalidMove {
			// TODO: store in statistics later, for now just skip
			continue
		}
		g.board = newBoard
		stuckCount = 0
		diff, done := g.board.Outcome(board.Black, board.White)
		if done {
			if diff > 0 {
				result.Winner = board.Black
			} else if diff < 0 {
				result.Winner = board.White
			} else {
				result.Winner = board.Empty
			}
			result.Difference = diff
			finished = true
		}
	}
	return &result
}
