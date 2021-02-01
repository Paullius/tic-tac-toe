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
    ID        string
    StartTime time.Time
    Board     *Board
    NextMove  enum.Move
    GameMode  enum.Mode
    mu        sync.Mutex
}

// CreateGame creates new game
func CreateGame(gm enum.Mode) *Game {
    id := uuid.New().String()
    g := &Game{ID: id, NextMove: enum.X, GameMode: gm, StartTime: time.Now().UTC()}
    gb := &Board{}
    gb.Init()
    g.Board = gb

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
    err := g.Board.Move(pm, x, y)
    if err == nil {
        if pm == enum.X {
            g.NextMove = enum.O
        } else {
            g.NextMove = enum.X
        }
    }

    return err
}
