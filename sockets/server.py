import socket


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((socket.gethostname(), 1234))
s.listen(5)



while True:
    clientSocket, clientAddress = s.accept()
    print(f"connection established from {clientAddress}")
    clientSocket.send(bytes("Hey!!!, welcomr to the server!!!", 'utf-8'))
    clientSocket.close()