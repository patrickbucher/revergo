# revergo

Reversi implementation in Go for educational purposes (AI algorithms, golang, testing, simulations).

## TODO

- Board
    - [x] determine valid moves
    - [x] apply player's move
    - [x] calculate outcome
    - [ ] render board as string
- Players
    - [ ] `STDIN` Player: Plays interactively from `STDIN`.
    - [x] Random Player: Picks a random move for every round.
    - [x] Corner Player: Goes for corners first, then for a random move.
    - [ ] Corner/Edge Player: Goes for corners first, then for edges, then for a random move.
    - [x] Minimax Player: Applies Minimax algorithm (concurrently) with a limited step.
    - [ ] other ideas: Minimax player with corner strategy, …
- Modes
    - [x] Simulation with two players, playing a lot of rounds against one another
    - [ ] Tournament with multiple players, playing against each other twice (first and second leg)

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
