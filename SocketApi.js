const SocketApi = (socket) => {
    return {
        checkData(event = "", data = {}) {
            switch (event) {
                case "connection":
                    if (!data.hasOwnProperty("type")) {
                        console.error("Provide exact data about connection: `trying to connect`, `connected`, `disconnected`")
                        return false
                    }
                    return true
                case "message":
                    if (!data.hasOwnProperty("type") || !data.hasOwnProperty("text")) {
                        console.error("For messages use objects like",
                            {type: "client|server", text: "", sender: "(optional)"}
                        );
                        return false
                    }
                    if (data.type === "server" && data.hasOwnProperty("sender")) {
                        console.error("For server messages no `sender` allowed")
                        return false
                    }
                    return true
                case "default":
                    if (Object.keys(data).length !== 0) {
                        console.error("No data allowed to send in `default` event")
                        return false
                    }
                    console.warn("Sending `default` to WS")
                    return true
                default:
                    console.error("Undefined event: ", event)
                    return false
            }
        },
        emit(eventName = "default", data) {
            if (!this.checkData(eventName, data)) {
                console.error("Data is not verified")
                return
            }

            console.log(data)
            socket.send(JSON.stringify({
                event: eventName,
                data: JSON.stringify(data),
            }))
        }
    }
}
