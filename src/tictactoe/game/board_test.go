package game

import (
	"reflect"
	"testing"
)

func TestBoard_GetWinner(t *testing.T) {
    tests := []struct {
        name string
        gb   *Board
        want *Move
    }{
        // TODO: Add test cases.
        {
            name: "No winners",
            gb:   &Board{moves: [3][3]*Move{}},
        },
        {
            name: "Winner: X - vertical 1",
            gb: &Board{moves: [3][3]*Move{
                {&Move{Type: 'X'}, nil, nil},
                {&Move{Type: 'X'}, nil, nil},
                {&Move{Type: 'X'}, nil, nil},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - vertical 2",
            gb: &Board{moves: [3][3]*Move{
                {nil, &Move{Type: 'X'}, nil},
                {nil, &Move{Type: 'X'}, nil},
                {nil, &Move{Type: 'X'}, nil},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - vertical 3",
            gb: &Board{moves: [3][3]*Move{
                {nil, nil, &Move{Type: 'X'}},
                {nil, nil, &Move{Type: 'X'}},
                {nil, nil, &Move{Type: 'X'}},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - horizontal 1",
            gb: &Board{moves: [3][3]*Move{
                {&Move{Type: 'X'}, &Move{Type: 'X'}, &Move{Type: 'X'}},
                {nil, nil, nil},
                {nil, nil, nil},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - horizontal 2",
            gb: &Board{moves: [3][3]*Move{
                {nil, nil, nil},
                {&Move{Type: 'X'}, &Move{Type: 'X'}, &Move{Type: 'X'}},
                {nil, nil, nil},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - horizontal 3",
            gb: &Board{moves: [3][3]*Move{
                {nil, nil, nil},
                {nil, nil, nil},
                {&Move{Type: 'X'}, &Move{Type: 'X'}, &Move{Type: 'X'}},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - diagonal 1",
            gb: &Board{moves: [3][3]*Move{
                {&Move{Type: 'X'}, nil, nil},
                {nil,  &Move{Type: 'X'}, nil},
                {nil,nil, &Move{Type: 'X'}},
            }},
            want: &Move{Type: 'X'},
        },
        {
            name: "Winner: X - diagonal 2",
            gb: &Board{moves: [3][3]*Move{
                {nil,nil, &Move{Type: 'X'}},
                {nil,  &Move{Type: 'X'}, nil},
                {&Move{Type: 'X'}, nil, nil},
            }},
            want: &Move{Type: 'X'},
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
