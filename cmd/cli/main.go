package main

import (
	"log"

	"github.com/foolusion/rl"
	"github.com/nsf/termbox-go"
)

func draw(g *rl.Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(g.Player.Position.X(), g.Player.Position.Y(), '@', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	hc := humanController{}

	g := rl.NewGame(hc)

	draw(g)

	for {
		quit := g.Step()
		if quit {
			return
		}
		draw(g)
	}
}

type humanController struct{}

func (h humanController) Get() rl.Action {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEsc:
			return rl.Quit
		case termbox.KeyArrowLeft:
			return rl.MoveLeft
		case termbox.KeyArrowRight:
			return rl.MoveRight
		case termbox.KeyArrowUp:
			return rl.MoveUp
		case termbox.KeyArrowDown:
			return rl.MoveDown
		}
	case termbox.EventError:
		log.Fatal(ev.Err)
	}
	return nil
}
