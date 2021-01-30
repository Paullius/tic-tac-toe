package game

import (
	"github.com/google/uuid"
)

// Game is tic-tac-toe instance
type Game struct {
    ID    string
    board *Board
}

// CreateGame is for creating new game
func CreateGame() *Game {
    id := uuid.New().String()
    g := &Game{ID: id}
    gb := &Board{}
    gb.Init()
    g.board = gb

    return g
}

// Move is for player move
func (g *Game) Move(pm *Move, x, y int) error {
    return g.board.Move(pm, x, y)
}

// Status returs game status
func (g *Game) Status() [][]rune {
    status := make([][]rune, len(g.board.moves))
    for x, row := range g.board.moves {
        status[x] = make([]rune, len(row))
        for y, player := range row {
            if player == nil {
                status[x][y] = ' '
            } else {
                status[x][y] = player.Type
            }
        }
    }

    return status
}

// GetResults is restuls
func (g *Game) GetResults() string {
    winner := g.board.GetWinner()
    if winner != nil {
        return "Winner is " + string(winner.Type)
    }

    if g.board.IsComplete() {
        return "Draw"
    }

    return "In Progress"
}
