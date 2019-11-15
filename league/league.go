package main

import (
	"fmt"
	"log"
	"revergo/player"
	"revergo/tournament"
)

func main() {
	t := tournament.NewTournament()
	t.AddPlayer("Randy Random", player.NewRandomPlayer)
	t.AddPlayer("Conny Corner", player.NewCornerPlayer)
	r, err := t.Play(10)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(r.Table())
}
