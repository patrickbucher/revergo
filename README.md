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

    $ go run league/league.go -n 5

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Mini Max VI.           194        80        64        14         2      1026
           2  Mini Max IV.           143        80        47        31         2       385
           3  Conny Corner           133        80        41        29        10       131
           4  Corner Defense         132        80        43        34         3        95
           5  Corner Avoidance       115        80        36        37         7       -41
           6  Mini Max II.           108        80        35        42         3       127
           7  Max Round               98        80        32        46         2      -174
           8  Edgy Edge               87        80        28        49         3      -507
           9  Randy Random            51        80        15        59         6     -1042

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
