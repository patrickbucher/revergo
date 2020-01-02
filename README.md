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
           1  Mini Max V.            377       160       125        33         2      3034
           2  Mini Max III.          301       160       100        59         1      1757
           3  Corner Defense         264       160        84        64        12        98
           4  Conny Corner           258       160        84        70         6       377
           5  Edgy Edge              241       160        79        77         4        40
           6  Corner Avoidance       237       160        76        75         9       134
           7  Max Round              176       160        57        98         5     -1507
           8  Mini Max I.            150       160        49       108         3     -2276
           9  Randy Random           133       160        43       113         4     -1657

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
