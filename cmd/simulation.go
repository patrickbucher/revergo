package main

import (
	"flag"
	"fmt"
	"os"
	"revergo/board"
	"revergo/game"
	"revergo/player"
	"strings"
	"text/tabwriter"
)

func main() {
	numberOfGames := flag.Int("n", 1, "number of games")
	flag.Parse()
	if *numberOfGames < 1 {
		fmt.Fprintf(os.Stderr, "unable to play %d rounds\n", *numberOfGames)
		os.Exit(1)
	}
	playerBlack := player.NewRandomPlayer(board.Black)
	playerWhite := player.NewRandomPlayer(board.White)
	playerBlackWins, playerWhiteWins, ties, diff := 0, 0, 0, 0
	for i := 0; i < *numberOfGames; i++ {
		game := game.NewGame(playerBlack, playerWhite)
		result := game.Play()
		diff += result.Difference
		if result.Winner == board.Black {
			playerBlackWins++
		} else if result.Winner == board.White {
			playerWhiteWins++
		} else {
			ties++
		}
	}
	printResults(playerBlackWins, playerWhiteWins, ties, diff)
}

func printResults(blackWins, whiteWins, ties, diff int) {
	const headFormat = "%-16s\t%5s\t%5s\t%5s\t%5s\n"
	const rowFormat = "%-16s\t%5d\t%5d\t%5d\t%5d\n"
	sep16 := strings.Repeat("-", 16)
	sep5 := strings.Repeat("-", 5)
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, headFormat, "Player", "Won", "Lost", "Tied", "Diff")
	fmt.Fprintf(tw, headFormat, sep16, sep5, sep5, sep5, sep5)
	fmt.Fprintf(tw, rowFormat, "Black", blackWins, whiteWins, ties, diff)
	fmt.Fprintf(tw, rowFormat, "White", whiteWins, blackWins, ties, diff*-1)
	tw.Flush()
}
