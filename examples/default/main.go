package main

import (
	"log"

	"github.com/PapayaJuice/goose"
)

// This example simply opens a new Goose game window and displays the default
// splash screen. This is a great way to ensure Goose is working on your
// system.
func main() {
	err := goose.Run(nil)
	if err != nil {
		log.Fatal(err)
	}
}
