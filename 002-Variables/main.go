package main

import "log"

// We can't use :=  shortcut for global variables.
// global_int := 20
var global_str = "Global string"

func main() {

	var num int
	var decimal float32

	num = 10
	decimal = 20.4

	dynamic_str := "This is a dynamic string !!! "

	log.Println("Integer number : ", num)
	log.Println("Float number : ", decimal)
	log.Println("Dynamic string : ", dynamic_str)
	log.Println("Global variable : ", global_str)

}
