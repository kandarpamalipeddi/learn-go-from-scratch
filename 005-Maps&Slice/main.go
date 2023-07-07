package main

import "log"

func main() {
	log.Println("This is an Map example")

	family_map := make(map[string]string)
	family_map["mom"] = "Manisha Ghule"
	family_map["dad"] = "Kandarpa Malipeddi"
	family_map["daughter"] = "Mahika Malipeddi"
	family_map["son"] = "Dhruv Malipeddi"

	mySlice := []string{
		"Kandarpa Malipeddi",
		"Manisha Ghule",
		"Mahika Malipeddi",
		"Dhruv Malipeddi",
	}

	log.Println(family_map)
	log.Println(mySlice)
}
