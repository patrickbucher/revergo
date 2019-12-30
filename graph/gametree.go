package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"revergo/board"
	"time"
)

type Node struct {
	Value    int
	Level    int
	MoveNr   int
	MovePath string
	Move     string
	Children []*Node
}

func (n *Node) Graphviz(name string) io.Reader {
	buf := bytes.NewBufferString(fmt.Sprintf("digraph %s {\n", name))

	// first, define all the nodes
	nodes := make(map[string]int)
	n.MapChildren(nodes)
	for key, value := range nodes {
		node := fmt.Sprintf("\t\"%s\" [label=\"%d\"];\n", key, value)
		buf.WriteString(node)
	}

	// second, define the relationships
	relationships := make(map[string][]*Node)
	n.ChildRelationships(relationships)
	for parent, children := range relationships {
		for _, child := range children {
			relationship := fmt.Sprintf("\t\"%s\" -> \"%s\" [label=\" %s \"];\n",
				parent, child.Identifier(), child.Move)
			buf.WriteString(relationship)
		}
	}

	buf.WriteString("}")
	return buf
}

func (n *Node) ChildRelationships(m map[string][]*Node) {
	m[n.Identifier()] = make([]*Node, 0)
	for _, child := range n.Children {
		if child == nil {
			continue
		}
		m[n.Identifier()] = append(m[n.Identifier()], child)
		child.ChildRelationships(m)
	}
}

func (n *Node) MapChildren(m map[string]int) {
	if n == nil {
		return
	}
	key := n.Identifier()
	m[key] = n.Value
	for _, child := range n.Children {
		child.MapChildren(m)
	}
}

func (n *Node) Identifier() string {
	return n.MovePath
}

func main() {
	depth := flag.Int("depth", 1, "depth of the tree")
	ahead := flag.Int("ahead", 0, "number of (random) moves to play ahead")
	flag.Parse()
	if *depth <= 0 {
		log.Fatal("depth must be positive")
	}
	if *ahead < 0 {
		log.Fatal("ahead must not be negative ")
	}
	b := board.InitialBoard()

	playerLastMove := board.Black
	playerNextMove := board.White
	if *ahead > 0 {
		rand.Seed(time.Now().Unix())
		for player := board.Black; *ahead > 0; *ahead-- {
			moves := b.ValidMoves(player)
			m := moves[rand.Intn(len(moves))]
			b, _ = b.Play(m, player)
			playerLastMove = player
			if player == board.Black {
				player = board.White
			} else {
				player = board.Black
			}
			playerNextMove = player
		}
	}

	root := buildGameTree(b, playerLastMove, playerNextMove, *depth, 0, 0, nil, "")
	graph := root.Graphviz("gametree")
	io.Copy(os.Stdout, graph)
}

func buildGameTree(b *board.Board, player, opponent board.State, depth, level, moveNr int,
	move *board.Move, movePath string) *Node {
	if depth <= 0 {
		return nil
	}
	diff, _ := b.Outcome(1, 2) // always view from same player's perspective
	movePath = fmt.Sprintf("%s-%d", movePath, moveNr)
	node := Node{
		Value:    diff,
		Level:    level,
		MoveNr:   moveNr,
		MovePath: movePath,
		Children: make([]*Node, 0),
	}
	if move != nil {
		node.Move = fmt.Sprintf("%d/%d", move.Row, move.Col)
	}
	validMoves := b.ValidMoves(player)
	for moveNr, move := range validMoves {
		newBoard, err := b.Play(move, player)
		if err != nil {
			log.Fatalf("play move %v on board %v: %v", move, b, err)
		}
		child := buildGameTree(newBoard, opponent, player, depth-1, level+1, moveNr, move, movePath)
		node.Children = append(node.Children, child)
	}
	return &node
}
