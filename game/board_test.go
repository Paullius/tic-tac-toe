package game

import (
	"reflect"
	"testing"

	"github.com/Paullius/tic-tac-toe/game/enum"
)

func TestBoard_GetWinner(t *testing.T) {
    tests := []struct {
        name string
        gb   *Board
        want enum.Move
    }{
        // TODO: Add test cases.
        {
            name: "No winners",
            gb:   &Board{moves: [3][3]enum.Move{}},
        },
        {
            name: "Winner: X - vertical 1",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.X, enum.NoMove, enum.NoMove},
                {enum.X, enum.NoMove, enum.NoMove},
                {enum.X, enum.NoMove, enum.NoMove},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - vertical 2",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.NoMove, enum.X, enum.NoMove},
                {enum.NoMove, enum.X, enum.NoMove},
                {enum.NoMove, enum.X, enum.NoMove},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - vertical 3",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.NoMove, enum.NoMove, enum.X},
                {enum.NoMove, enum.NoMove, enum.X},
                {enum.NoMove, enum.NoMove, enum.X},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - horizontal 1",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.X, enum.X, enum.X},
                {enum.NoMove, enum.NoMove, enum.NoMove},
                {enum.NoMove, enum.NoMove, enum.NoMove},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - horizontal 2",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.NoMove, enum.NoMove, enum.NoMove},
                {enum.X, enum.X, enum.X},
                {enum.NoMove, enum.NoMove, enum.NoMove},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - horizontal 3",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.NoMove, enum.NoMove, enum.NoMove},
                {enum.NoMove, enum.NoMove, enum.NoMove},
                {enum.X, enum.X, enum.X},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - diagonal 1",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.X, enum.NoMove, enum.NoMove},
                {enum.NoMove,  enum.X, enum.NoMove},
                {enum.NoMove,enum.NoMove, enum.X},
            }},
            want: enum.X,
        },
        {
            name: "Winner: X - diagonal 2",
            gb: &Board{moves: [3][3]enum.Move{
                {enum.NoMove,enum.NoMove, enum.X},
                {enum.NoMove,  enum.X, enum.NoMove},
                {enum.X, enum.NoMove, enum.NoMove},
            }},
            want: enum.X,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.gb.GetWinner(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Board.GetWinner() = %v, want %v", got, tt.want)
            }
        })
    }
}
