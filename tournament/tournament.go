package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"revergo/board"
	"revergo/game"
	"revergo/player"
	"sort"
	"strings"
	"text/tabwriter"
)

type tournamentPlayer struct {
	spawnFunc func(board.State, string) *player.Player
	name      string
}

type pairing struct {
	blackPlayer *player.Player
	whitePlayer *player.Player
}

// Tournament represents a series of games played by multiple players against
// every other player.
type Tournament struct {
	players map[string]*tournamentPlayer
}

// Result encloses all player's outcomes of the tournament played.
type Result struct {
	Stats map[string]*PlayerStatistics
}

// Table renders the result as a statistics table.
func (r *Result) Table() string {
	statsTable := make([]PlayerStatistics, 0)
	for _, entry := range r.Stats {
		statsTable = append(statsTable, *entry)
	}
	return statisticsTable(statsTable).Render()
}

// NewTournament creates a new, empty tournament.
func NewTournament() *Tournament {
	var tournament Tournament
	tournament.players = make(map[string]*tournamentPlayer, 0)
	return &tournament
}

// AddPlayer adds a new player to the tournament. If a player with the given
// name was already added to the tournament before, an error is thrown. A
// player is passed as a function that creates a new player instance, because
// the player's color (board.Black or board.White) is established for every new
// game.
func (t *Tournament) AddPlayer(name string, spawnFunc func(board.State, string) *player.Player) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("empty name for player is not allowed")
	}
	if spawnFunc == nil {
		return errors.New("spawnFunc must not be nil")
	}
	if _, ok := t.players[name]; ok {
		return fmt.Errorf("a player with the name %s already joined the tournament", name)
	}
	newPlayer := tournamentPlayer{spawnFunc, name}
	t.players[name] = &newPlayer
	return nil
}

// Play executes all the games of a tournament and returns the result thereof.
// Every player is paired of with every other player twice with different
// colors assigned in order to ensure there's no first mover advantage. If the
// tournament is played without adding players before, an error is thrown.
func (t *Tournament) Play(rounds int) (*Result, error) {
	if len(t.players) < 2 {
		return nil, errors.New("need at least two players to play a tournament")
	}
	pairings := t.pairUp()
	results := make(map[string]*PlayerStatistics, 0)
	for i, round := range pairings {
		fmt.Println("play game", i+1)
		black := round.blackPlayer
		white := round.whitePlayer
		g := game.NewGame(black, white)
		r := g.Play()
		blackStat, ok := results[(*black).Name()]
		if !ok {
			blackStat = &PlayerStatistics{(*black).Name(), 0, 0, 0, 0, 0}
			results[(*black).Name()] = blackStat
		}
		whiteStat, ok := results[(*white).Name()]
		if !ok {
			whiteStat = &PlayerStatistics{(*white).Name(), 0, 0, 0, 0, 0}
			results[(*white).Name()] = whiteStat
		}
		if r.Winner == board.Black {
			(*blackStat).Wins++
			(*blackStat).Points += 3
			(*whiteStat).Losses++
		} else if r.Winner == board.White {
			(*whiteStat).Wins++
			(*whiteStat).Points += 3
			(*blackStat).Losses++
		} else if r.Winner == board.Empty {
			(*whiteStat).Ties++
			(*whiteStat).Ties++
			(*blackStat).Points++
			(*blackStat).Points++
		}
		(*blackStat).Diff += r.Difference
		(*whiteStat).Diff -= r.Difference
	}
	tournamentResult := Result{results}
	return &tournamentResult, nil
}

func (t *Tournament) pairUp() []pairing {
	pairings := make([]pairing, 0)
	players := make([]*tournamentPlayer, 0)
	for _, tournamentPlayer := range t.players {
		players = append(players, tournamentPlayer)
	}
	for index, p1 := range players {
		for _, p2 := range players[index+1:] {
			first := pairing{p1.spawnFunc(board.Black, p1.name), p2.spawnFunc(board.White, p2.name)}
			second := pairing{p2.spawnFunc(board.Black, p2.name), p1.spawnFunc(board.White, p1.name)}
			pairings = append(pairings, first)
			pairings = append(pairings, second)
		}
	}
	return pairings
}

// PlayerStatistics describes the player's statistics for a tournament. The
// points are handed out as follows: Win: 3, Tie: 1, Loss: 0.
type PlayerStatistics struct {
	Name   string
	Wins   int
	Losses int
	Ties   int
	Diff   int
	Points int
}

type statisticsTable []PlayerStatistics

func (s statisticsTable) Len() int           { return len(s) }
func (s statisticsTable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s statisticsTable) Less(i, j int) bool { return s[i].Points < s[j].Points }

// Render renders the statistics table in descending of the number of points scored.
func (s statisticsTable) Render() string {
	const headFormat = "%8s\t%-16s\t%8s\t%8s\t%8s\t%8s\t%8s\n"
	const rowFormat = "%8d\t%-16s\t%8d\t%8d\t%8d\t%8d\t%8d\n"
	var sep16 = strings.Repeat("-", 16)
	var sep8 = strings.Repeat("-", 8)
	sort.Sort(sort.Reverse(s))
	buf := bytes.NewBufferString("")
	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, headFormat, "Rank", "Player", "Points", "Won", "Lost", "Tied", "Diff")
	fmt.Fprintf(tw, headFormat, sep8, sep16, sep8, sep8, sep8, sep8, sep8)
	for rank, stats := range s {
		fmt.Fprintf(tw, rowFormat, rank+1, stats.Name, stats.Points, stats.Wins, stats.Losses,
			stats.Ties, stats.Diff)
	}
	tw.Flush()
	return buf.String()
}
