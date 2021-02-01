package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Paullius/tic-tac-toe/game"
	"github.com/Paullius/tic-tac-toe/game/enum"

	"github.com/gorilla/mux"
)

func gamesHandle(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    if r.Method == "POST" {
        r.Body = http.MaxBytesReader(w, r.Body, 1048576)
        dec := json.NewDecoder(r.Body)
        dec.DisallowUnknownFields()
        newGameParams := &struct {
            GameMode enum.Mode `json:"gameMode"`
        }{}
        err := dec.Decode(&newGameParams)
        if err != nil && err != io.EOF {
            msg := "Request body has to be JSON object"
            http.Error(w, msg, http.StatusInternalServerError)
            return
        }

        g := game.CreateGame(newGameParams.GameMode)
        cache.CleanUp()
        cache[g.ID] = g

        writeStatus(g.ID, w)
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
            X    int    `json:"x"`
            Y    int    `json:"y"`
            Move string `json:"move"`
        }{}

        err := dec.Decode(&moveParams)
        if err != nil && err != io.EOF {
            msg := "Request body has to be JSON object"
            http.Error(w, msg, http.StatusInternalServerError)
            return
        }
        vars := mux.Vars(r)
        id := vars["id"]

        if g, ok := cache[id]; ok {
            move := enum.Move([]rune(moveParams.Move)[0])
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

// StartServer starts API server
func StartServer(port int) {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/v1/games", gamesHandle)
    myRouter.HandleFunc("/v1/games/{id}", gameHandle)
    myRouter.HandleFunc("/v1/games/{id}/move", gameMoveHandle)
    http.ListenAndServe(":" + strconv.Itoa(port), myRouter)
}

func writeStatus(id string, w http.ResponseWriter) {
    if g, ok := cache[id]; ok {
        gameStatus := &struct {
            GameID     string      `json:"gameID"`
            Board      [][]string  `json:"board"`
            Result     enum.Result `json:"result"`
            IsComplete bool        `json:"isComplete"`
            NextMove   string      `json:"nextMove"`
            GameMode   byte        `json:"gameMode"`
        }{
            GameID:     g.ID,
            Board:      g.Board.GetStatusBoard(),
            Result:     g.Board.GetResultsEnum(),
            IsComplete: g.Board.IsCompleted(),
            NextMove:   string(g.NextMove),
            GameMode:   byte(g.GameMode),
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
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
