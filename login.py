for i in range(5):
    try:
        print("welcome to my login!")

        username = input("username: ")
        password = input("password: ")

        if username == "admin" and password == "1234":
            print("your are login now!\n") 
            print("Thank for comming! {}\n".format(username))
            print("good bey!\n\n")
            break

        elif "a" in username and "1" in password:
            print("your have a in username and 1 in password")
            break
        else:
            print("your username and password not cerrented!")
            print("please try again")
    except KeyboardInterrupt:
        E = KeyboardInterrupt
        print("\nYour are close program now come back later\n")
        print(E)
        break

print("hi")
