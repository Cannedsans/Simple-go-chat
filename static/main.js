const ws = new WebSocket("ws://localhost:8080/ws")

ws.onmessage = (event) => {
    const msg = JSON.parse(event.data)
    const item = document.createElement("li")

    item.textContent = msg.username + ": " + msg.message
    saveMessage(item.textContent)
    document.getElementById("messages").appendChild(item)
}

document.getElementById("form").onsubmit = (event) => {
    event.preventDefault()

    const username = document.getElementById("username").value
    const message = document.getElementById("message").value

    ws.send(JSON.stringify({ username, message }))

    document.getElementById("message").value = ""
}
