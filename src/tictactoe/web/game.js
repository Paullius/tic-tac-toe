var defaultData = {
    gameid: "",
    board: [[" ", " ", " "], [" ", " ", " "], [" ", " ", " "]],
    nextMove: "X",
    result: "",
    newGamemode: 0,
    isComplete: false
};
function setStatus(data, respone) {
    data.gameid = respone.GameID;
    data.board = respone.Board
    data.nextMove = respone.NextMove;
    data.result = respone.Result
    data.isComplete = respone.IsComplete;
    data.gamemode = respone.GameMode;
}
var app = new Vue({
    el: '#game',
    data: Object.assign({}, defaultData),
    methods: {
        select(position) {
            if (this.isComplete === true) return
            axios
                .post('http://localhost:8080/v1/games/' + this.gameid + "/move", {
                    'X': position[0],
                    'Y': position[1],
                    'Move': this.nextMove
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
                .post('http://localhost:8080/v1/games', { 'GameMode': parseInt(this.newGamemode) }).then(response => {
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
