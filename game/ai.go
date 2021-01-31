package game

import (
	"math/rand"

	"github.com/Paullius/tic-tac-toe/game/enum"
)

// AI lvl1 selects random empty cell
func (g *Game) doAILvl1Move() error {

    if g.IsComplete() {
        return nil
    }
    emptyCells := g.board.GetEmptyCells()
    move := getRandomCell(emptyCells)

    return makeAIMove(g, move)
}

// AI lvl2 selects move
func (g *Game) doAILvl2Move() error {

    if g.IsComplete() {
        return nil
    }
    emptyCells := g.board.GetEmptyCells()

    var prevMove enum.Move
    if g.NextMove == enum.X {
        prevMove = enum.O
    } else {
        prevMove = enum.X
    }

    // make move if AI wins
    for _, move := range emptyCells {
        if isMoveMakeWinner(g.board, move, g.NextMove) {
            return makeAIMove(g, move)
        }
    }

    // make a move if it prevents loosing
    for _, move := range emptyCells {
        if isMoveMakeWinner(g.board, move, prevMove) {
            return makeAIMove(g, move)
        }
    }

    move := getRandomCell(emptyCells)

    return makeAIMove(g, move)
}

func makeAIMove(g *Game, move [2]int) error {
    pm := g.NextMove

    err := g.moveBoard(pm, move[0], move[1])

    return err
}

func getRandomCell(emptyCells [][2]int) [2]int {
    idx := rand.Intn(len(emptyCells) - 1)
    selected := emptyCells[idx]

    return selected
}

func isMoveMakeWinner(board *Board, move [2]int, moveType enum.Move) bool {
    board.moves[move[0]][move[1]] = moveType
    winner := board.GetWinner()
    board.moves[move[0]][move[1]] = enum.NoMove

    return winner == moveType
}
