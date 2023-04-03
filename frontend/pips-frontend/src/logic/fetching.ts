
export var getMatch = (async (): Promise<String | void> => {
    let fetched = await fetch('http://localhost:8080/match')
    let json = await fetched.json()
    return json
})