package main

import (
	"flag"
	"fmt"
	"os"
	"revergo/board"
	"revergo/game"
	"revergo/player"
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
	const headFormat = "%s\t%s\t%s\t%s\t%s\n"
	const rowFormat = "%s\t%3d\t%4d\t%4d\t%4d\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, headFormat, "Player", "Won", "Lost", "Tied", "Diff")
	fmt.Fprintf(tw, headFormat, "------", "---", "----", "----", "----")
	fmt.Fprintf(tw, rowFormat, "Black", blackWins, whiteWins, ties, diff)
	fmt.Fprintf(tw, rowFormat, "White", whiteWins, blackWins, ties, diff*-1)
	tw.Flush()
}
