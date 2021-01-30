package api

import (
	"encoding/json"
	"io"
	"net/http"

	"../game"
	"github.com/gorilla/mux"
)

var cache map[string]*game.Game = map[string]*game.Game{}

func gamesHandle(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "test")
    if r.Method == "POST" {
        g := game.CreateGame()
        cache[g.ID] = g
        io.WriteString(w, "New game created: "+g.ID)
    } else if r.Method == "GET" {
        io.WriteString(w, "Games:\n\n")
        for id := range cache {
            io.WriteString(w, "Game ID: "+id)
        }
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
            io.WriteString(w, "\nResults:\n")
            result := g.GetResults()
            io.WriteString(w, result)

        } else {
            io.WriteString(w, "No Game Found")
        }

    } else if r.Method == "PUT" {
        //TODO: move
    }
    w.WriteHeader(http.StatusOK)

}

func gameMoveHandle(w http.ResponseWriter, r *http.Request) {

    r.Body = http.MaxBytesReader(w, r.Body, 1048576)
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    moveParams := &struct {
        X    int
        Y    int
        Move string
    }{}
    err := dec.Decode(&moveParams)
    if err != nil {
        msg := "Request body has to be JSON object"
        http.Error(w, msg, http.StatusInternalServerError)
    }

    vars := mux.Vars(r)
    id := vars["id"]
    if g, ok := cache[id]; ok {
        io.WriteString(w, "Game ID: "+g.ID)
        move := &game.Move{Type: []rune(moveParams.Move)[0]}
        g.Move(move,moveParams.X,moveParams.Y)
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
    myRouter.HandleFunc("/v1/games", gamesHandle)
    myRouter.HandleFunc("/v1/games/{id}", gameHandle)
    myRouter.HandleFunc("/v1/games/{id}/move", gameMoveHandle)
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
