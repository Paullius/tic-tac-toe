var example1 = new Vue({
    el: '#game',
    data: {
        counter: 0,
        gameid: ""
    },
    methods: {
        startGame() {
            axios
                .post('http://localhost:8080/v1/games').then(response => {
                    this.gameid = response.data.GameID

                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        getGameStatus() {
            axios
                .get('http://localhost:8080/v1/games/' + this.gameid).then(response => {
                    alert(response.data)
                    console.log(response.data)
                    //  this.gameid = response.data.GameID

                })
                .catch(function (error) {
                    console.log(error);
                })
        },
        moveX() {
            axios
                .post('http://localhost:8080/v1/games/' + this.gameid + "/move", {
                    'X': 2,
                    'Y': 0,
                    'Move': 'X'
                }).then(response => {
                    alert(response.data)
                    console.log(response.data)
                    //  this.gameid = response.data.GameID

                })
                .catch(function (error) {
                    console.log(error);
                })
        }, moveO() {
            axios
                .post('http://localhost:8080/v1/games/' + this.gameid + "/move", {
                    'X': 2,
                    'Y': 0,
                    'Move': 'O'
                }, {
                    headers: {
                        'Content-Type': 'application/json',
                        'Access-Control-Allow-Origin' : '*'
                    }
                }).then(response => {
                    alert(response.data)
                    console.log(response.data)
                    //  this.gameid = response.data.GameID

                })
                .catch(function (error) {
                    console.log(error);
                })
        }
    },
})
