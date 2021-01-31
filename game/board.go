package game

import (
	"errors"

	"github.com/Paullius/tic-tac-toe/game/enum"
)

// Board is game board with player moves
type Board struct {
    moves [3][3]enum.Move
}

// Init initialize the board
func (gb *Board) Init() {
    mv := [3][3]enum.Move{}
    gb.moves = mv

}

// Move does player move
func (gb *Board) Move(pm enum.Move, x, y int) error {

    if err := gb.validateMove(pm, x, y); err != nil {
        return err
    }
    gb.moves[x][y] = pm
    return nil
}

func (gb *Board) validateMove(pm enum.Move, x, y int) error {

    if x < 0 || y < 0 || x >= len(gb.moves) || y >= len(gb.moves[0]) {
        return errors.New("invalid move - out of range")
    }

    if gb.moves[x][y] != 0 {
        return errors.New("invalid move - move already exists")
    }

    if pm != enum.X && pm != enum.O {
        return errors.New("invalid move - " + string(pm))
    }

    return nil
}

// IsComplete checks if game board is completed
func (gb *Board) IsComplete() bool {
    for _, row := range gb.moves {
        for _, move := range row {
            if move == enum.NoMove {
                return false
            }
        }
    }

    return true
}

// GetEmptyCells gets empty cell coordinates
func (gb *Board) GetEmptyCells() [][2]int {
    emptyCells := [][2]int{}
    for x, row := range gb.moves {
        for y, move := range row {
            if move == enum.NoMove {
                emptyCells = append(emptyCells, [2]int{x,y})
            }
        }
    }

    return emptyCells
}


// GetWinner gets winner move
func (gb *Board) GetWinner() enum.Move {
    l := len(gb.moves)
    var move enum.Move

    // horizontal check
    for r := 0; r < l; r++ {
        move = gb.moves[r][0]
        if move == enum.NoMove {
            continue
        }
        for c := 0; c < l; c++ {
            if gb.moves[r][c] == 0 || gb.moves[r][c] != move {
                move = enum.NoMove
                break
            }
        }
        if move != enum.NoMove {
            return move
        }
    }

    // vertical check
    for c := 0; c < l; c++ {
        move = gb.moves[0][c]
        if move == enum.NoMove {
            continue
        }
        for r := 0; r < l; r++ {
            if gb.moves[r][c] == 0 || gb.moves[r][c] != move {
                move = enum.NoMove
                break
            }
        }
        if move != enum.NoMove {
            return move
        }
    }

    move = gb.moves[0][0]
    if move != 0 {
        for d := 0; d < l; d++ {
            if gb.moves[d][d] == 0 || gb.moves[d][d] != move {
                move = enum.NoMove
                break
            }
        }
        if move != enum.NoMove {
            return move
        }
    }

    move = gb.moves[0][l-1]
    if move != enum.NoMove {
        for d := 0; d < l; d++ {
            bd := l - d - 1
            if gb.moves[d][bd] == 0 || gb.moves[d][bd] != move {
                move = enum.NoMove
                break
            }
        }
        if move != enum.NoMove {
            return move
        }
    }

    return enum.NoMove
}
