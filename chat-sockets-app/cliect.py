import socket
import sys
import threading
import time

header_length = 10
buffer_size = 16

def socket_msg(msg):  return bytes(f"{len(msg):<{header_length}}{msg}", "utf-8")
user_name = input(f"UserName: ")
empty_msg = socket_msg("")

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

def handel_recv(sock):
    while True:
        time.sleep(2)
        data = recive_msg(sock)
        print(data['msg'])

def handel_send(sock):
    while True:
        msg = socket_msg(input(f"> "))
        
        if(msg == empty_msg): 
            print("entered a empty message")
            sock.close()
            sys.exit(0)
        
        print(sock)
        sock.send(msg)



with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    try:
        s.setblocking(True)
        s.connect((socket.gethostname(), 4200))
        s.send(socket_msg(user_name))
        recv_thread = threading.Thread(target=handel_recv, args=(s,))
        send_thread = threading.Thread(target=handel_send, args=(s,))
        
        print(s)        
        send_thread.start()
        recv_thread.start()
            
    except Exception as ex:
        print(ex)
    finally:
        s.close()