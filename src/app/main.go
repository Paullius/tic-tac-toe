package main

import (
	api "../tictactoe/api"
)

func main() {
	// g := tictactoe.CreateGame()
	// pm := tictactoe.PlayerMove{}
	// err := g.Move(pm, 0, 0)

	// fmt.Println(g, err)

	api.StartServer()
}
