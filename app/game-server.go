package main

import (
	"github.com/Paullius/tic-tac-toe/api"
)

const serverPort = 8080

func main() {
    api.StartServer(serverPort)
}
