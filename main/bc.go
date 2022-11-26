package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Date: ")
}

func main() {
	log.Println("now")
	fmt.Println("Hello")
}
