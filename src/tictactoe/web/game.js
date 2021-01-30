var defaultData = {
    gameid: "",
    board: [[" ", " ", " "], [" ", " ", " "], [" ", " ", " "]],
    nextMove: "X",
    result: ""
};
var app = new Vue({
    el: '#game',
    data: Object.assign({}, defaultData),
    methods: {
        select(position) {
            axios
                .post('http://localhost:8080/v1/games/' + this.gameid + "/move", {
                    'X': position[0],
                    'Y': position[1],
                    'Move': this.nextMove
                }).then(response => {
                    console.log(response.data)
                    this.board = response.data.Status;
                    this.nextMove = response.data.NextMove;
                    this.result = response.data.Result;

                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        startGame() {
            axios
                .post('http://localhost:8080/v1/games').then(response => {
                    this.gameid = response.data.GameID
                    this.board = defaultData.board,
                        this.nextMove = defaultData.nextMove,
                        this.result = defaultData.result
                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        getGameStatus() {
            if (this.gameid == "") return
            axios
                .get('http://localhost:8080/v1/games/' + this.gameid).then(response => {
                    console.log(response.data)
                    this.board = response.data.Status;
                    this.nextMove = response.data.NextMove;
                    this.result = response.data.Result;
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
