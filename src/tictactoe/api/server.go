package api

import (
	"io"
	"net/http"
	"strconv"
	"unicode"

	"../game"
	"github.com/gorilla/mux"
)

var cache map[string]*game.Game = map[string]*game.Game{}

func gamesHandle(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "test")
    if r.Method == "POST" {
        g := game.CreateGame()
        cache[g.ID] = g
    } else if r.Method == "GET" {
        io.WriteString(w, "Need to show games\n\n")
        g := game.CreateGame()
        cache[g.ID] = g
        io.WriteString(w, "New game created: "+g.ID)
    }
    w.WriteHeader(http.StatusOK)

}

func gameHandle(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        vars := mux.Vars(r)
        id := vars["id"]
        if g, ok := cache[id]; ok {
            io.WriteString(w, "Game ID: "+g.ID)
            status := getStatus(g)
            io.WriteString(w, "\nStatus:\n")
            io.WriteString(w, status)

        } else {
            io.WriteString(w, "No Game Found")
        }

    } else if r.Method == "PUT" {
        //TODO: move
    }
    w.WriteHeader(http.StatusOK)

}

func gameMoveHandle(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]
    if g, ok := cache[id]; ok {
        move := []rune(vars["move"])[0]
        x, _ := strconv.Atoi(vars["x"])
        y, _ := strconv.Atoi(vars["y"])
        // fmt.Println("MOVE: ", string(move), x, y)
        move = unicode.ToUpper(move)
        pm := &game.Move{Type: move}
        if err := g.Move(pm, x, y); err != nil {
            io.WriteString(w, "ERROR: "+err.Error())
        }
        io.WriteString(w, "Game ID: "+g.ID)
        status := getStatus(g)
        io.WriteString(w, "\nStatus:\n")
        io.WriteString(w, status)
        io.WriteString(w, "\nResults:\n")
        result := g.GetResults()
        io.WriteString(w, result)

    } else {
        io.WriteString(w, "No Game Found")
    }

}

// StartServer is for starting API server
func StartServer() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/v1/game", gamesHandle)
    myRouter.HandleFunc("/v1/game/{id}", gameHandle)
    myRouter.HandleFunc("/v1/game/{id}/{x}/{y}/{move}", gameMoveHandle)
    http.ListenAndServe(":8080", myRouter)
}

func getStatus(g *game.Game) string {
    status := g.Status()
    st := ""
    for _, row := range status {
        st += "|"
        for _, move := range row {
            st += string(move) + "|"
        }
        st += "\n"
    }

    return st
}
