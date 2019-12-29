package main

import (
	"flag"
	"fmt"
	"os"
	"revergo/board"
	"revergo/game"
	"revergo/player"
	"strings"
	"sync"
	"text/tabwriter"
)

type gameWinner int

const (
	tie   gameWinner = 0
	black gameWinner = 1
	white gameWinner = 2
)

type gameResult struct {
	winner gameWinner
	diff   int
}

func main() {
	numberOfRounds := flag.Int("n", 1, "number of rounds to play")
	flag.Parse()
	if *numberOfRounds < 1 {
		fmt.Fprintf(os.Stderr, "unable to play %d rounds\n", *numberOfRounds)
		os.Exit(1)
	}
	playerBlack := player.NewCornerPlayer(board.Black, "Conny Corner")
	playerWhite := player.NewMinimaxPlayerSpawnFunc(7)(board.White, "Mini Max VII.")
	playerBlackWins, playerWhiteWins, ties, diff := 0, 0, 0, 0
	var wg sync.WaitGroup
	ch := make(chan gameResult, 0)
	for i := 0; i < *numberOfRounds; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			game := game.NewGame(playerBlack, playerWhite)
			result := game.Play()
			winner := tie
			if result.Winner == board.Black {
				winner = black
			} else if result.Winner == board.White {
				winner = white
			}
			ch <- gameResult{winner, result.Difference}
		}()
	}
	go func() {
		for result := range ch {
			diff += result.diff
			if result.winner == black {
				playerBlackWins++
			} else if result.winner == white {
				playerWhiteWins++
			} else if result.winner == tie {
				ties++
			}
		}
		printResults((*playerBlack).Name(), (*playerWhite).Name(), playerBlackWins, playerWhiteWins,
			ties, diff)
	}()
	wg.Wait()
	close(ch)
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
