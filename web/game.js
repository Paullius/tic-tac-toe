var apiServer = "http://localhost:8080"

var defaultData = {
    gameid: "",
    board: [[" ", " ", " "], [" ", " ", " "], [" ", " ", " "]],
    nextMove: "X",
    result: "",
    newGamemode: 0,
    isComplete: false
};
function setStatus(data, respone) {
    data.gameid = respone.gameID;
    data.board = respone.board
    data.nextMove = respone.nextMove;
    data.result = respone.result
    data.isComplete = respone.isComplete;
    data.gamemode = respone.gameMode;
}
var app = new Vue({
    el: '#game',
    data: Object.assign({}, defaultData),
    methods: {
        select(position) {
            if (this.isComplete === true) return
            axios
                .post(apiServer + '/v1/games/' + this.gameid + "/move", {
                    'x': position[0],
                    'y': position[1],
                    'move': this.nextMove
                }).then(response => {
                    setStatus(this, response.data)
                    // console.log(response.data)
                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        startGame() {
            axios
                .post(apiServer + '/v1/games', { 'GameMode': parseInt(this.newGamemode) }).then(response => {
                    setStatus(this, response.data)
                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        getGameStatus() {
            if (this.gameid == "" || this.isComplete === true) return
            axios
                .get('http://localhost:8080/v1/games/' + this.gameid).then(response => {
                    setStatus(this, response.data)
                    console.log(response.data)
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    },
    beforeMount() {
        this.startGame();

        setInterval(function () {
            this.getGameStatus();
        }.bind(this), 1000);
    }
})
