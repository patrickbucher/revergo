package player

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"revergo/board"
	"time"
)

// StdinPlayer is a player for interactive game play that gets its moves from
// standard input.
type StdinPlayer struct {
	state  board.State
	name   string
	symbol rune
}

var black = 'x'
var white = 'o'
var empty = '-'

// NewStdinPlayer creates a new stdin player.
func NewStdinPlayer(state board.State, name string) *Player {
	symbol := empty
	if state == board.Black {
		symbol = black
	} else {
		symbol = white
	}
	stdinPlayer := StdinPlayer{state, name, symbol}
	rand.Seed(time.Now().Unix())
	player := Player(&stdinPlayer)
	return &player
}

var inputFormat = regexp.MustCompile("^[a-h]{1}[1-8]{1}$")
var rowLabels = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
var colLabels = []rune{'1', '2', '3', '4', '5', '6', '7', '8'}

// Play shows the current board on stdout, requests for a move from stdin,
// applies the move and shows the result on stdout.
func (p *StdinPlayer) Play(b *board.Board) *board.Move {
	rendering, _ := b.Render(black, white, empty, rowLabels, colLabels)
	fmt.Println(rendering)
	fmt.Println(p.standing(b))
	fmt.Println()
	moves := b.ValidMoves(p.state)
	if len(moves) == 0 {
		return nil
	}
	validInput := false
	var move *board.Move
	for !validInput {
		fmt.Printf("Player '%c': ", p.symbol)
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading input from stdin: %v\n", err)
			continue
		}
		if !inputFormat.Match([]byte(input)) {
			fmt.Fprintf(os.Stderr, "input format must be: [row][col] (such as 'a1' or 'c5')\n")
			continue
		}
		row := indexOf(rune([]byte(input)[0]), rowLabels)
		col := indexOf(rune([]byte(input)[1]), colLabels)
		move = &board.Move{row, col}
		if elementOf(move, moves) {
			validInput = true
		} else {
			fmt.Fprintln(os.Stderr, "invalid move")
		}
	}
	result, err := b.Play(move, p.state)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to apply move %v\n", move)
		return nil
	}
	rendering, _ = result.Render(black, white, empty, rowLabels, colLabels)
	fmt.Println(rendering)
	fmt.Println(p.standing(result))
	fmt.Println()
	time.Sleep(time.Second * 2)
	return move
}

func (p *StdinPlayer) standing(b *board.Board) string {
	opponent := board.Other(p.state)
	ourPoints := b.NumberOfFields(p.state)
	oppPoints := b.NumberOfFields(opponent)
	const format = "Standing: %c %d:%d %c"
	if p.state == board.Black {
		return fmt.Sprintf(format, black, ourPoints, oppPoints, white)
	}
	return fmt.Sprintf(format, black, oppPoints, ourPoints, white)
}

func indexOf(element rune, list []rune) int {
	for i, e := range list {
		if e == element {
			return i
		}
	}
	return -1
}

func elementOf(move *board.Move, moves []*board.Move) bool {
	for _, m := range moves {
		if move.Equal(m) {
			return true
		}
	}
	return false
}

// State returns the player's state (Black or White).
func (p *StdinPlayer) State() board.State {
	return p.state
}

// Name returns the player's name.
func (p *StdinPlayer) Name() string {
	return p.name
}
