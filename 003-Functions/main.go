package main

import (
	"fmt"
	"log"
)

func print_string(str string) {
	fmt.Println("Input string : ", str)
}

func format_ourput_string(str string) string {
	output_str := "Formatter string : " + str
	return output_str
}

func format_string_with_error(str string) (string, error) {
	if str == "" {
		err := fmt.Errorf("No input string provided !!!")
		return str, err
	}

	return "Formatted string : " + str, nil
}

func main() {
	print_string("Hello All !!!")
	log.Println(format_ourput_string("Hello There !!!"))
	log.Println(format_string_with_error("Hello World !!!"))
	str, err := format_string_with_error("")
	if err == nil {
		log.Println(str)
	} else {
		log.Println(err)
	}

}
