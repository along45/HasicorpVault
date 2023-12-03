sername = "admin"
password = "secret123"

def login():
    user = input("Enter username: ")
    passwd = input("Enter password: ")
    if user == username and passwd == password:
        print("Access granted")
    else:
        print("Access denied")
