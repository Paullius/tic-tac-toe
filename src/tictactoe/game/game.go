package game

import (
	"errors"

	"github.com/google/uuid"
)
const (
    // X move
    X = 'X'
    // O move
    O = 'O'
)

// Game is tic-tac-toe instance
type Game struct {
    ID       string
    board    *Board
    NextMove rune
    gameMode Mode
}

// CreateGame is for creating new game
func CreateGame() *Game {
    id := uuid.New().String()
    g := &Game{ID: id, NextMove: X}
    gb := &Board{}
    gb.Init()
    g.board = gb

    return g
}

// Move is for player move
func (g *Game) Move(pm *Move, x, y int) error {
    if pm.Type != g.NextMove {
        return errors.New("Invalid move")
    }
    err := g.board.Move(pm, x, y)
    if err == nil {
        if pm.Type == X {
            g.NextMove = O
        } else {
            g.NextMove = X
        }
    }

    return err
}

// StatusBoard returs game status
func (g *Game) StatusBoard() [][]string {
    status := make([][]string, len(g.board.moves))
    for x, row := range g.board.moves {
        status[x] = make([]string, len(row))
        for y, player := range row {
            if player == nil {
                status[x][y] = " "
            } else {
                status[x][y] = string(player.Type)
            }
        }
    }

    return status
}

// GetResultsEnum is for getting game results enum
func (g *Game) GetResultsEnum() ResultEnum {
    winner := g.board.GetWinner()
    if winner != nil {
        if winner.Type == X {
            return WinX // WIN - X
        }
        return WinO // WIN - O
    }

    if g.board.IsComplete() {
        return Draw //"DRAW"
    }

    return InProgress //"INPROGRESS"
}

// IsComplete is game complete
func (g *Game) IsComplete() bool {
    return g.board.IsComplete() || g.board.GetWinner() != nil
}
