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
