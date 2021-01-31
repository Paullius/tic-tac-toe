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

	var prevMoveWins *[2]int
	for _, move := range emptyCells {
		// make move if AI wins
		if isMoveMakeWinner(g.board, move, g.NextMove) {
			return makeAIMove(g, move)
        }
        // save move to prevent loosing
		if isMoveMakeWinner(g.board, move, prevMove) {
			prevMoveWins = &move
		}
	}

    // found move that prevents loosing
	if prevMoveWins != nil {
		return makeAIMove(g, *prevMoveWins)
	}

	// random move if no win or loosing
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
