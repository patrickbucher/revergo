#!/bin/sh

for depth in {2..7}
do
    go run ../graph/gametree.go -depth $depth > tmp.dot
    dot -Tsvg tmp.dot -o tree-${depth}.svg
    rm -f tmp.dot
done
