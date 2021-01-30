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
		// io.WriteString(w, "New Game Created: "+g.ID)

		newGame := &struct {
			GameID string
		}{
			GameID: g.ID,
		}
		js, err := json.Marshal(newGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
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
		if g, ok := cache[id]; ok {
			// io.WriteString(w, "Game ID: "+g.ID)
			// status := getStatus(g)
			// io.WriteString(w, "\nStatus:\n")
			// io.WriteString(w, status)
			// io.WriteString(w, "\nResults:\n")
			// result := g.GetResults()
			// io.WriteString(w, result)

			gameStatus := &struct {
				GameID string
				Status [][]string
                Result string
                NextMove string
			}{
				GameID: g.ID,
				Status: g.Status(),
                Result: g.GetResults(),
                NextMove: string(g.NextMove),
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
	w.WriteHeader(http.StatusOK)

}

func gameMoveHandle(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)
		// bodyString := string(body)
		// fmt.Println(bodyString)

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
			// io.WriteString(w, "Game ID: "+g.ID)
			move := &game.Move{Type: []rune(moveParams.Move)[0]}
            err = g.Move(move, moveParams.X, moveParams.Y)
            if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// status := getStatus(g)
			// io.WriteString(w, "\nStatus:\n")
			// io.WriteString(w, status)
			// io.WriteString(w, "\nResults:\n")
			// result := g.GetResults()
			// io.WriteString(w, result)
			gameStatus := &struct {
				GameID string
				Status [][]string
				Result string
                NextMove string
			}{
				GameID: g.ID,
				Status: g.Status(),
				Result: g.GetResults(),
                NextMove: string(g.NextMove),
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
