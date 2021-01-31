package game

import (
	"errors"
	"sync"
	"time"

	"github.com/Paullius/tic-tac-toe/game/enum"
	"github.com/google/uuid"
)

// Game is tic-tac-toe game
type Game struct {
    ID       string
    StartTime time.Time
    board    *Board
    NextMove enum.Move
    GameMode enum.Mode
    mu       sync.Mutex
}

// CreateGame creates new game
func CreateGame(gm enum.Mode) *Game {
    id := uuid.New().String()
    g := &Game{ID: id, NextMove: enum.X, GameMode: gm, StartTime: time.Now().UTC()}
    gb := &Board{}
    gb.Init()
    g.board = gb

    return g
}

// Move does player move
func (g *Game) Move(pm enum.Move, x, y int) error {
    g.mu.Lock()
    defer g.mu.Unlock()
    if pm != g.NextMove {
        return errors.New("Invalid move")
    }
    err := g.moveBoard(pm, x, y)
    // AI move
    if err == nil {
        if g.GameMode == enum.PvAIv1 {
            err = g.doAILvl1Move()
        } else if g.GameMode == enum.PvAIv2 {
            err = g.doAILvl2Move()
        }
    }

    return err
}

func (g *Game) moveBoard(pm enum.Move, x, y int) error {
    err := g.board.Move(pm, x, y)
    if err == nil {
        if pm == enum.X {
            g.NextMove = enum.O
        } else {
            g.NextMove = enum.X
        }
    }

    return err
}

// GetStatusBoard gets game board matrix
func (g *Game) GetStatusBoard() [][]string {
    status := make([][]string, len(g.board.moves))
    for x, row := range g.board.moves {
        status[x] = make([]string, len(row))
        for y, move := range row {
            if move == enum.NoMove {
                status[x][y] = " "
            } else {
                status[x][y] = string(move)
            }
        }
    }

    return status
}

// GetResultsEnum gets game result enum
func (g *Game) GetResultsEnum() enum.Result {
    winMove := g.board.GetWinner()
    if winMove != enum.NoMove {
        if winMove == enum.X {
            return enum.WinX
        }
        return enum.WinO
    }

    if g.board.IsComplete() {
        return enum.Draw
    }

    return enum.InProgress
}

// IsComplete checks if the game is complete
func (g *Game) IsComplete() bool {
    return g.board.IsComplete() || g.board.GetWinner() != enum.NoMove
}
