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
	numberOfRounds := flag.Int("n", 1, "number of rounds to play")
	flag.Parse()
	if *numberOfRounds < 1 {
		fmt.Fprintf(os.Stderr, "unable to play %d rounds\n", *numberOfRounds)
		os.Exit(1)
	}
	playerBlack := player.NewStdinPlayer(board.Black, "Standard Input")
	playerWhite := player.NewMinimaxPlayerSpawnFunc(6)(board.White, "Mini Max VI.")
	playerBlackWins, playerWhiteWins, ties, diff := 0, 0, 0, 0
	for i := 0; i < *numberOfRounds; i++ {
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
	printResults((*playerBlack).Name(), (*playerWhite).Name(), playerBlackWins, playerWhiteWins,
		ties, diff)
}

func printResults(blackName, whiteName string, blackWins, whiteWins, ties, diff int) {
	const headFormat = "%-16s\t%8s\t%8s\t%8s\t%8s\n"
	const rowFormat = "%-16s\t%8d\t%8d\t%8d\t%8d\n"
	sep16 := strings.Repeat("-", 16)
	sep8 := strings.Repeat("-", 8)
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, headFormat, "Player", "Won", "Lost", "Tied", "Diff")
	fmt.Fprintf(tw, headFormat, sep16, sep8, sep8, sep8, sep8)
	fmt.Fprintf(tw, rowFormat, blackName, blackWins, whiteWins, ties, diff)
	fmt.Fprintf(tw, rowFormat, whiteName, whiteWins, blackWins, ties, diff*-1)
	tw.Flush()
}
