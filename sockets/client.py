import socket


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((socket.gethostname(), 1234))

final_msg = ""

while True:
    msgBytes = s.recv(1)
    msg = msgBytes.decode("utf-8")
    
    if(len(msg) <= 0): break

    print(msg)
    final_msg += msg

print(final_msg)

