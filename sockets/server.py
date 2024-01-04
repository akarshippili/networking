import socket
import time
import threading


HEADER_LENGHT = 10
def getSocketMsg(msg):  return f"{len(msg):<{HEADER_LENGHT}}{msg}"

def serve(sock, addr, fn):
    print(f"[thread-id: {threading.get_native_id()}] connection established from {addr}")    
    try:
        msg = getSocketMsg("Hey!!!, welcome to the server!!!")
        sock.send(bytes(msg, 'utf-8'))
        
        count = 10
        while(count):
            time.sleep(1)
            msg = getSocketMsg(str(fn()))
            sock.send(bytes(msg, 'utf-8'))
            count -= 1 
    except Exception as ex:
        print(ex)
    finally:
        sock.close()
    
def listenAndServe():
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
            s.bind((socket.gethostname(), 1234))
            s.listen(5)
            while True:
                clientSocket, clientAddress = s.accept()
                threading.Thread(target=serve, args=(clientSocket, clientAddress, time.time)).start()
    except Exception as ex:
        print(ex)
    finally:
        s.close()

listenAndServe()