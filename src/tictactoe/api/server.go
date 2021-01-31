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
    enableCors(&w)
    if r.Method == "POST" {
        g := game.CreateGame()
        cache[g.ID] = g

        writeStatus(g.ID, w)

        // newGame := &struct {
        //     GameID string
        // }{
        //     GameID: g.ID,
        // }
        // js, err := json.Marshal(newGame)
        // if err != nil {
        //     http.Error(w, err.Error(), http.StatusInternalServerError)
        //     return
        // }
        // w.Header().Set("Content-Type", "application/json")
        // w.Write(js)
    } else if r.Method == "GET" {
        io.WriteString(w, "Games:\n\n")
        for id := range cache {
            io.WriteString(w, "Game ID: "+id)
        }
    }
    w.WriteHeader(http.StatusOK)

}

func gameHandle(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "GET" {
        vars := mux.Vars(r)
        id := vars["id"]
        writeStatus(id, w)
    }
    w.WriteHeader(http.StatusOK)
}

func gameMoveHandle(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "POST" {

        r.Body = http.MaxBytesReader(w, r.Body, 1048576)
        dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()
        moveParams := &struct {
            X    int
            Y    int
            Move string
        }{}

        err := dec.Decode(&moveParams)
        if err != nil && err != io.EOF {
            msg := "Request body has to be JSON object"
            http.Error(w, msg, http.StatusInternalServerError)
        }
        vars := mux.Vars(r)
        id := vars["id"]

        if g, ok := cache[id]; ok {
            move := &game.Move{Type: []rune(moveParams.Move)[0]}
            err = g.Move(move, moveParams.X, moveParams.Y)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            writeStatus(id, w)
        } else {
            io.WriteString(w, "No Game Found")
        }

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

func writeStatus(id string, w http.ResponseWriter) {
    if g, ok := cache[id]; ok {
        gameStatus := &struct {
            GameID     string
            Status     [][]string
            Result     string
            IsComplete bool
            NextMove   string
        }{
            GameID:     g.ID,
            Status:     g.Status(),
            Result:     g.GetResults(),
            IsComplete: g.IsComplete(),
            NextMove:   string(g.NextMove),
        }
        js, err := json.Marshal(gameStatus)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)

    } else {
        io.WriteString(w, "No Game Found")
    }
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
