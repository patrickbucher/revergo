# revergo

Reversi implementation in Go for educational purposes (AI algorithms, golang, testing, simulations).

## TODO

- Board
    - [x] determine valid moves
    - [x] apply player's move
    - [x] calculate outcome
    - [x] render board as string
- Players
    - [x] `STDIN` Player: Plays interactively from `STDIN`.
    - [x] Random Player: Picks a random move for every round.
    - [x] Corner Player: Goes for corners first, then for a random move.
    - [x] Corner/Edge Player: Goes for corners first, then for edges, then for a random move.
    - [x] Minimax Player: Applies Minimax algorithm (concurrently) with a limited step.
    - [x] Minimax Player: shake search tree with alpha-beta pruning
    - [ ] other ideas: Minimax player with corner strategy, …
- Modes
    - [x] Simulation with two players, playing a lot of rounds against one another
    - [x] Tournament with multiple players, playing against each other twice (first and second leg)

## League

Run a league with 10 rounds (all players play 20 rounds against one another
with colors mixed in order to ensure there's no first mover advantage):

    $ go run league/league.go -n 10

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Conny Corner           310       160       102        54         4       953
           2  Mini Max VI.           301       160        96        51        13      1576
           3  Corner Defense         289       160        93        57        10       545
           4  Corner Avoidance       260       160        84        68         8       181
           5  Mini Max III.          245       160        80        75         5      1046
           6  Max Round              214       160        70        86         4      -294
           7  Edgy Edge              201       160        66        91         3      -651
           8  Mini Max I.            177       160        54        91        15     -1676
           9  Randy Random           130       160        42       114         4     -1680

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
