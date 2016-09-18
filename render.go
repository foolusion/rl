package rl

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

type RenderController interface {
	Render(g *game)
	Close()
}

type termboxRenderer struct{}

func NewTermboxRenderer() (*termboxRenderer, error) {
	err := termbox.Init()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize termbox")
	}
	return &termboxRenderer{}, nil

}

func (t *termboxRenderer) Render(g *game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(g.positions[0].X(), g.positions[0].Y(), '@', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func (t *termboxRenderer) Close() {
	termbox.Close()
}
