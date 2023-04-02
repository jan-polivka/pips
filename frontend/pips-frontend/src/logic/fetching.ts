
export var getMatch = (async (): Promise<String> => {
    let ret = await fetch('http://localhost:8080/match')
        .then(
            async (response) => {
                if (response.status !== 200) {
                    console.log("there was an issue: "
                        + response.status)
                    return
                }
                response.json().then(function (data) {
                    console.log(data);
                    ret = data
                    return data
                });
            }
        )
    return "test"
})