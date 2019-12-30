#!/bin/sh

rm -f *.svg *.png
for ahead in {0,54,55}
do
    for depth in {3..6}
    do
        go run ../graph/gametree.go -depth $depth -ahead $ahead > tmp.dot
        dot -Tsvg tmp.dot -o tree-depth-${depth}-ahead-${ahead}.svg
        dot -Tpng tmp.dot -o tree-depth-${depth}-ahead-${ahead}.png
        rm -f tmp.dot
    done
done
