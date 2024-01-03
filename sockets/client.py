import socket


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((socket.gethostname(), 1234))


buffer_size = 16
HEADER_LENGHT = 10

while True:
    
    first_msg = True
    msg_length = None
    msg = ""
    
    while True:
        sub_msg_bytes = s.recv(buffer_size)
        sub_msg = sub_msg_bytes.decode("utf-8")
        
        if(first_msg):
            
            if(len(sub_msg) == 0): 
                s.close()
                exit(0)
                
            first_msg = False
            msg_length = int(sub_msg[:HEADER_LENGHT])
            msg += sub_msg[HEADER_LENGHT:]
        else:  
            msg += sub_msg
        
        if(len(msg) == msg_length):
            print(f"Message: {msg}, Lenght: {msg_length}.")
            final_msg, msg_length, msg = True, None, ""
            break

