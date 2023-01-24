package app

import (
	in "github.com/pan-asovsky/DaysCalculator/internal/transport"
	"log"
)

func Run() {
	router := in.GetRouter()

	err := router.Run()
	if err != nil {
		log.Fatalln("Fatal error at router.Run():", err)
	}
}
