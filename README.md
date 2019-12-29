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

    $ go run league/league.go -n 7

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Corner Defense         228       126        73        44         9       693
           2  Conny Corner           228       126        75        48         3       687
           3  Mini Max IV.           210       126        70        56         0      1090
           4  Mini Max III.          205       126        68        57         1       737
           5  Max Round              194       126        64        60         2      -263
           6  Corner Avoidance       189       126        61        59         6        97
           7  Edgy Edge              187       126        62        63         1      -219
           8  Mini Max I.            172       126        56        66         4     -1034
           9  Mini Max II.           152       126        49        72         5      -372
          10  Randy Random           108       126        35        88         3     -1416

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
