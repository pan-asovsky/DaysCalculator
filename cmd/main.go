package main

import (
	in "github.com/pan-asovsky/DaysCalculator/internal"
	"log"
)

func main() {

	router := in.GetRouter()

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run():", err)
	}
}
