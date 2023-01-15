import net from 'net'

const server = net.createServer(socket => {
    console.log(`Server started, started listening ${socket.remoteAddress}:${socket.remotePort}`)
    socket.write("Hello client, pls feel free to send messages")

    socket.on("data", (data) => console.log(`recived data ${data.toString()}`))
})

server.listen(8080, "localhost")

