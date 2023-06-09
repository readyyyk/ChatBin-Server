const ROOM = "TROOM"
const socket = new SocketApi(`ws://localhost:3001/${ROOM}/ws`)

// connection handler and emitter
socket.addDataChecker("connection", ()=>true)
socket.emit("connection", {detail: "trying to connect"})

socket.on("connection", e=> {
    console.log(e)
    if(e.detail === "trying to connect")
        document.body.innerHTML = document.body.innerHTML +
            "<p>smb trying to connect...<p>"
    else if (e.detail === "connected")
        document.body.innerHTML = document.body.innerHTML +
            `<p style="color:green"><b>${e.name}</b> connected</p>`
    else if (e.detail === "disconnected")
        document.body.innerHTML = document.body.innerHTML +
            `<p style="color:red"><i>${e.name}</i> disconnected</p>`
})

// msg handler
const msg = (text, sender="") => {
    const el = document.createElement("p")
    el.innerHTML = `<p style="color:${sender?"blue":"darkblue"}">${sender||"You"}: ${text} </p>`
    return el
}
socket.addDataChecker("message", ()=>true)
window.addEventListener("keydown", e=>
    e.key==="Enter"?sendMessage():""
)
function sendMessage(){
    const msg = document.querySelector("#msg")
    socket.emit("message", {text: msg.value})
    msg.value = ""
}
document.querySelector("#sbm").onclick = sendMessage

socket.on("message", e=>{
    console.log(e)
    document.body.append(msg(e.text, e.sender))
})

// prompt for name
function enterName(){
    const name = prompt("Name to connect")
    if(!name){
        enterName()
        return
    }
    fetch(`http://localhost:3001/${ROOM}/names?name=${name}`)
        .then(data => {
            if(!data.ok){
                enterName()
                return
            }
            socket.emit("connection", {detail: "connected", name: name})
        })
}
enterName()
