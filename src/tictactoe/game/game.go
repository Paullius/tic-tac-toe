package game

import (
	"errors"
	"math/rand"
	"sync"

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
    GameMode ModeEnum
    mu       sync.Mutex
}

// CreateGame is for creating new game
func CreateGame(gm ModeEnum) *Game {
    id := uuid.New().String()
    g := &Game{ID: id, NextMove: X, GameMode: gm}
    gb := &Board{}
    gb.Init()
    g.board = gb

    return g
}

// Move is for player move
func (g *Game) Move(pm *Move, x, y int) error {
    g.mu.Lock()
    defer g.mu.Unlock()
    if pm.Type != g.NextMove {
        return errors.New("Invalid move")
    }
    err := g.moveBoard(pm, x, y)
    // AI move
    if err == nil && g.GameMode == PvAIv1 {
        g.doAILvl1Move()
    }

    return err
}

func (g *Game) moveBoard(pm *Move, x, y int) error {
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
            return WinX
        }
        return WinO
    }

    if g.board.IsComplete() {
        return Draw
    }

    return InProgress
}

// IsComplete is game complete
func (g *Game) IsComplete() bool {
    return g.board.IsComplete() || g.board.GetWinner() != nil
}

// AI selects random empty cell
func (g *Game) doAILvl1Move() error {

    if g.IsComplete() {
        return nil
    }
    emptyCells := g.board.GetEmptyCells()

    idx := rand.Intn(len(emptyCells) - 1)
    selected := emptyCells[idx]

    pm := &Move{Type: g.NextMove}

    err := g.moveBoard(pm, selected[0], selected[1])

    return err
}
