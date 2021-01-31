var defaultData = {
    gameid: "",
    board: [[" ", " ", " "], [" ", " ", " "], [" ", " ", " "]],
    nextMove: "X",
    result: "",
    isComplete: false
};
function setStatus(data,respone){
    data.gameid = respone.GameID;
    data.board = respone.Status
    data.nextMove = respone.NextMove;
    data.result = respone.Result
    data.isComplete = respone.IsComplete;
}
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
                    setStatus(this,response.data)
                    console.log(response.data)
                    // this.board = response.data.Status;
                    // this.nextMove = response.data.NextMove;
                    // this.result = response.data.Result;
                    // this.isComplete = response.data.IsComplete;

                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        startGame() {
            axios
                .post('http://localhost:8080/v1/games').then(response => {
                    setStatus(this,response.data)

                    // this.gameid = response.data.GameID;
                    // this.board = defaultData.board;
                    // this.nextMove = defaultData.nextMove;
                    // this.result = defaultData.result
                    // this.isComplete = response.data.IsComplete;
                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        getGameStatus() {
            if (this.gameid == "") return
            axios
                .get('http://localhost:8080/v1/games/' + this.gameid).then(response => {
                    setStatus(this,response.data)
                    console.log(response.data)

                    // this.board = response.data.Status;
                    // this.nextMove = response.data.NextMove;
                    // this.result = response.data.Result;
                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    },
    beforeMount() {
        this.startGame();

        setInterval(function () {
            //this.getGameStatus();
        }.bind(this), 1000);
    }
})
