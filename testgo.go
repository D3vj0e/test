package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		fmt.Println("input your username and password!")

		username := "admin"
		password := "1234"

		//fmt.Print("username:")
		//fmt.Scan(&username)
		//fmt.Print("password:")
		//fmt.Scan(&password)

		if username == "admin" && password == "1234" {
			fmt.Println("Your are login now!")
			print("\n\n")
			fmt.Println("your username is :", username, "\nyour password is : ", password)
			fmt.Println("Good bye")

			if username > "a" && password > "1" {
				print("\n\n")
				fmt.Println("your has a and 1 string in your username and password")
				break
			}

		} else {
			fmt.Println("Error your username or password not cerrented!")
		}
	}
}
