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

	g := &rl.Game{
		Player: rl.Player{
			Position: rl.Vector{10, 10, 0, 0},
		},
	}

	draw(g)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowLeft:
				rl.Move(g, [4]int{-1, 0, 0, 0})
			case termbox.KeyArrowRight:
				rl.Move(g, [4]int{1, 0, 0, 0})
			case termbox.KeyArrowUp:
				rl.Move(g, [4]int{0, -1, 0, 0})
			case termbox.KeyArrowDown:
				rl.Move(g, [4]int{0, 1, 0, 0})
			}
		case termbox.EventError:
			log.Fatal(err)
		}
		draw(g)
	}
}
