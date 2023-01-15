import dgram from 'dgram';

const server = dgram.createSocket('udp4')

server.on('error', (err)=>{
    console.log(`server error : ${err}`);
    server.close(() => console.log(`server closed`))
})

server.on('message', (message, info) => {
    console.log(`recived a datagram ${message} from ${info.address}:${info.port}`);
})

server.on('listening', () => {
    const address = server.address();
    console.log(`server listening ${address.address}:${address.port}`);
});


server.bind({
    address: "localhost",
    port: 5500
})