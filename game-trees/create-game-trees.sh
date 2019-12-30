#!/bin/sh

for depth in {2..6}
do
    go run ../graph/gametree.go -depth $depth > tmp.dot
    dot -Tsvg tmp.dot -o tree-${depth}.svg
    dot -Tpng tmp.dot -o tree-${depth}.png
    rm -f tmp.dot
done
