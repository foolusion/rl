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

	var g game

	termbox.SetCell(g.playerx, g.playery, '@', termbox.ColorDefault, termbox.ColorDefault)
	if err := termbox.Flush(); err != nil {
		log.Fatal(err)
	}

	for !g.done {
		// get input
		var action string
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			action = handleEvent(&g, ev)
		case termbox.EventError:
			break
		}

		// update
		switch action {
		case "MOVE_UP":
			g.playery -= 1
		case "MOVE_DOWN":
			g.playery += 1
		case "MOVE_RIGHT":
			g.playerx += 1
		case "MOVE_LEFT":
			g.playerx -= 1
		case "QUIT":
			g.done = true
		default:
			continue
		}

		// render
		if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
			log.Fatal(err)
		}
		termbox.SetCell(g.playerx, g.playery, '@', termbox.ColorDefault, termbox.ColorDefault)
		if err := termbox.Flush(); err != nil {
			log.Fatal(err)
		}

	}
}

type game struct {
	playerx int
	playery int
	done    bool
}

func handleEvent(g *game, ev termbox.Event) string {
	switch ev.Ch {
	case 0:
		return handleKey(g, ev.Key)
	case 'h', 'H':
		return "MOVE_LEFT"
	case 'j', 'J':
		return "MOVE_DOWN"
	case 'k', 'K':
		return "MOVE_UP"
	case 'l', 'L':
		return "MOVE_RIGHT"
	default:
		return ""
	}

}

func handleKey(g *game, k termbox.Key) string {
	switch k {
	case termbox.KeyArrowUp:
		return "MOVE_UP"
	case termbox.KeyArrowDown:
		return "MOVE_DOWN"
	case termbox.KeyArrowRight:
		return "MOVE_RIGHT"
	case termbox.KeyArrowLeft:
		return "MOVE_LEFT"
	case termbox.KeyEsc:
		return "QUIT"
	default:
		return ""
	}
}
