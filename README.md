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

    $ go run league/league.go -n 2

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Mini Max V.             73        32        24         7         1       641
           2  Mini Max III.           69        32        23         9         0       551
           3  Corner Avoidance        56        32        18        12         2        73
           4  Conny Corner            49        32        16        15         1        -2
           5  Corner Defense          43        32        14        17         1       -14
           6  Edgy Edge               42        32        14        18         0      -144
           7  Randy Random            36        32        12        20         0      -318
           8  Max Round               34        32        11        20         1      -285
           9  Mini Max I.             27        32         9        23         0      -502

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
