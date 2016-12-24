package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.SetCell(0, 0, '@', termbox.ColorDefault, termbox.ColorDefault)
	if err := termbox.Flush(); err != nil {
		log.Fatal(err)
	}

	termbox.PollEvent()
}
