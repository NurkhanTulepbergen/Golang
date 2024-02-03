package cmd

import (
	"TSIS1/internal"
	"log"
)

func main() {
	log.Println("starting API server")
	internal.StartServer()

}
