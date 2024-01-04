import socket
import sys

header_length = 10
buffer_size = 16

def socket_msg(msg):  return bytes(f"{len(msg):<{header_length}}{msg}", "utf-8")

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

user_name = input(f"UserName: ")
empty_msg = socket_msg("")

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    try:
        s.connect((socket.gethostname(), 4200))
        s.send(socket_msg(user_name))

        while True:
            # msg = socket_msg(input(f"{user_name}> "))
            # if(msg == empty_msg): 
            #     s.close()
            #     sys.exit(0)
                
            # s.send(msg)
            
            data = recive_msg(s)
            print(data['msg'])
            
            
    except Exception as ex:
        print(ex)
    finally:
        s.close()