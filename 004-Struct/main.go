package main

import "log"

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
}

func (u *User) getFullName() string {
	return u.firstName + " " + u.lastName
}

func main() {
	var myUser1 User
	myUser1.firstName = "Kandarpa"
	myUser1.lastName = "Malipeddi"

	myUser2 := User{
		firstName: "Manisha",
		lastName:  "Ghule",
	}

	log.Println(myUser1, myUser2)

	log.Println(myUser1.getFullName())
	log.Println(myUser2.getFullName())

}
