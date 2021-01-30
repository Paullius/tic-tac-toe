package game

import (
	"errors"
)

// Board is game board with player moves
type Board struct {
    moves [3][3]*Move
}

// Init is for GameBoard initialization
func (gb *Board) Init() {
    mv := [3][3]*Move{}
    gb.moves = mv

}

// Move does player move
func (gb *Board) Move(pm *Move, x, y int) error {

    if err := gb.validateMove(pm, x, y); err != nil {
        return err
    }
    gb.moves[x][y] = pm
    return nil
}

func (gb *Board) validateMove(pm *Move, x, y int) error {

    if x < 0 || y < 0 || x >= len(gb.moves) || y >= len(gb.moves[0]) {
        return errors.New("invalid move - out of range")
    }

    if gb.moves[x][y] != nil {
        return errors.New("invalid move - move already exists")
    }

    if pm.Type != rune('X') && pm.Type != rune('O') {
        return errors.New("invalid move - " + string(pm.Type))
    }

    return nil
}

// IsComplete checks if game board is completed
func (gb *Board) IsComplete() bool {
    for _, row := range gb.moves {
        for _, move := range row {
            if move == nil {
                return false
            }
        }
    }

    return true
}

// GetWinner is for geting results
func (gb *Board) GetWinner() *Move {
    l := len(gb.moves)
    var candidate *Move

    // horizontal check
    for r := 0; r < l; r++ {
        candidate = gb.moves[r][0]
        if candidate == nil {
            continue
        }
        for c := 0; c < l; c++ {
            if gb.moves[r][c] == nil || gb.moves[r][c].Type != candidate.Type {
                candidate = nil
                break
            }
        }
        if candidate != nil {
            return candidate
        }
    }

    // vertical check
    for c := 0; c < l; c++ {
        candidate = gb.moves[0][c]
        if candidate == nil {
            continue
        }
        for r := 0; r < l; r++ {
            if gb.moves[r][c] == nil || gb.moves[r][c].Type != candidate.Type {
                candidate = nil
                break
            }
        }
        if candidate != nil {
            return candidate
        }
    }

    candidate = gb.moves[0][0]
    if candidate != nil {
        for d := 0; d < l; d++ {
            if gb.moves[d][d] == nil || gb.moves[d][d].Type != candidate.Type {
                candidate = nil
                break
            }
        }
        if candidate != nil {
            return candidate
        }
    }

    candidate = gb.moves[0][l-1]
    if candidate != nil {
        for d := 0; d < l; d++ {
            bd :=l-d-1
            if gb.moves[d][bd] == nil || gb.moves[d][bd].Type != candidate.Type {
                candidate = nil
                break
            }
        }
        if candidate != nil {
            return candidate
        }
    }

    return nil
}
