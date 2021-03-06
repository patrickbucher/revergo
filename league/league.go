package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"revergo/player"
	"revergo/tournament"
)

func main() {
	numberOfRounds := flag.Int("n", 1, "number of rounds to play")
	flag.Parse()
	if *numberOfRounds < 1 {
		fmt.Fprintf(os.Stderr, "unable to play %d rounds\n", *numberOfRounds)
		os.Exit(1)
	}
	t := tournament.NewTournament()
	t.AddPlayer("Randy Random", player.NewRandomPlayer)
	t.AddPlayer("Conny Corner", player.NewCornerPlayer)
	t.AddPlayer("Corner Defense", player.NewCornerdefensePlayer)
	t.AddPlayer("Corner Avoidance", player.NewCorneravoidancePlayer)
	t.AddPlayer("Edgy Edge", player.NewEdgePlayer)
	t.AddPlayer("Max Round", player.NewMaxroundPlayer)
	t.AddPlayer("Mini Max II.", player.NewMinimaxPlayerSpawnFunc(2))
	t.AddPlayer("Mini Max IV.", player.NewMinimaxPlayerSpawnFunc(4))
	t.AddPlayer("Mini Max VI.", player.NewMinimaxPlayerSpawnFunc(6))
	r, err := t.Play(*numberOfRounds)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(r.Table())
}
