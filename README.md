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
           1  Mini Max IX.           579       280       188        77        15      2580
           2  Mini Max V.            554       280       184        94         2      4067
           3  Mini Max VIII.         548       280       181        94         5      2532
           4  Mini Max IV.           491       280       159       107        14       635
           5  Mini Max X.            469       280       156       123         1      1889
           6  Mini Max VI.           457       280       149       121        10      -329
           7  Conny Corner           442       280       145       128         7       382
           8  Corner Defense         441       280       145       129         6       380
           9  Mini Max III.          385       280       124       143        13      -108
          10  Max Round              352       280       114       156        10     -1708
          11  Mini Max II.           338       280       105       152        23     -1581
          12  Mini Max VII.          326       280        98       150        32       496
          13  Edgy Edge              324       280       105       166         9     -2375
          14  Mini Max I.            304       280       101       178         1     -3131
          15  Randy Random           211       280        67       203        10     -3729

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
