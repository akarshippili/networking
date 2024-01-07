import socket
import select
import threading
import sys

header_length = 10
buffer_size = 16
users = {}

def socket_msg(msg): return f"{len(msg):<{header_length}}{msg}"

def recive_msg(sock):
    
    first_msg = True
    msg_length = None
    msg = ""
    
    while True:
        sub_msg_bytes = sock.recv(buffer_size)
        sub_msg = sub_msg_bytes.decode("utf-8")
        
        if(first_msg):        
            if(len(sub_msg) == 0): 
                sock.close()
                return None
                
            first_msg = False
            msg_length = int(sub_msg[:header_length])
            msg += sub_msg[header_length:]
        else:  
            msg += sub_msg
        
        if(len(msg) == msg_length):
            data = {}
            data["msg_length"] = msg_length
            data["msg"] = msg
            return data

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server_socket:
    args = sys.argv
    server_socket.bind(("0.0.0.0", int(args[-1])))
    server_socket.listen(5)
    print(f"Starts serving at {socket.gethostname()}:{args[-1]}.")
    
    sockets = [server_socket]
    while True:
        read_sockets, write_sockets, exceptions_sockets = select.select(sockets, [], sockets)
        
        for rsock in read_sockets:
            
            # new connection
            if(rsock == server_socket):
                client_socket, client_address = rsock.accept()
                sockets.append(client_socket)
                
                new_user = recive_msg(client_socket)
                print(f"[thread-id: {threading.get_native_id()}] {new_user['msg']} successfully established connection from {client_address}")              
                users[client_socket] = new_user["msg"]
                for sock in sockets:
                    if(sock == client_socket or sock == server_socket): continue
                    sock.send(bytes(socket_msg(f"{new_user['msg']} joined the chat."), "utf-8"))
                continue
            
            # existing connections
            user = users[rsock]
            data = recive_msg(rsock)
            
            if(not data or len(data["msg"]) == 0):
                print(f"[thread-id: {threading.get_native_id()}] {user} closed connection.")              
                sockets.remove(rsock)
                rsock.close()
                del users[rsock]
                continue
            else:
                print(f"{user}> {data['msg']}")
                for sock in sockets:
                    if(sock == rsock or sock == server_socket): continue
                    sock.send(bytes(socket_msg(f"{user}> {data['msg']}"), "utf-8"))

        for exesock in exceptions_sockets:
            print(exesock)
        