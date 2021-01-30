

GET /v1/game/{GUID} - get game status
POST /v1/game - create game
POST /v1/game/{GUID}/move - create move

DELETE /v1/game/{GUID} - deltes game?


https://en.wikipedia.org/wiki/Tic-tac-toe


go get "github.com/google/uuid"
go get "github.com/gorilla/mux"

Examples:
http://localhost:8080/v1/game

http://localhost:8080/v1/game/e99ff85c-1e6d-46be-8b0e-2c48163a2884/0/0/x

