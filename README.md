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
    - [ ] other ideas: Minimax player with corner strategy, …
- Modes
    - [x] Simulation with two players, playing a lot of rounds against one another
    - [x] Tournament with multiple players, playing against each other twice (first and second leg)

## League

Run a league with 15 rounds (all players play 30 rounds against one another
with colors mixed in order to ensure there's no first mover advantage):

    $ go run league/league.go -n 15

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Conny Corner           413       210       135        67         8      1648
           2  Corner Defense         406       210       133        70         7      1346
           3  Mini Max III.          376       210       124        82         4      1855
           4  Mini Max II.           319       210       104        99         7       122
           5  Edgy Edge              298       210        97       106         7      -215
           6  Max Round              252       210        81       120         9     -1098
           7  Mini Max I.            232       210        76       130         4     -1621
           8  Randy Random           199       210        65       141         4     -2037

Run a league with only two rounds, but with additional minimax players with
deeper recursion levels:

        Rank  Player              Points     Games       Won      Lost      Tied      Diff
    --------  ----------------  --------  --------  --------  --------  --------  --------
           1  Mini Max V.             83        36        27         7         2       568
           2  Conny Corner            70        36        23        12         1       238
           3  Corner Defense          61        36        20        15         1        50
           4  Mini Max III.           60        36        20        16         0       148
           5  Max Round               54        36        18        18         0       -74
           6  Mini Max IV.            51        36        17        19         0        12
           7  Edgy Edge               41        36        13        21         2      -179
           8  Mini Max II.            40        36        13        22         1      -176
           9  Mini Max I.             40        36        13        22         1      -213
          10  Randy Random            36        36        12        24         0      -374

## Interactive Play

Run a game against the computer (from `stdin`; pick an opponent by modifying the source code):

    go run simulation/simulation.go 
