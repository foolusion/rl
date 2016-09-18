package rl

import termbox "github.com/nsf/termbox-go"

type Input int

const (
	InputError Input = iota
	InputN
	InputS
	InputE
	InputW
	InputNE
	InputSE
	InputNW
	InputSW
	InputA
	InputB
	InputC
	InputD
	InputStart
	InputLeft
	InputRight
)

type InputController interface {
	GetInput(*game)
}

type termboxInput struct {
	events chan termbox.Event
}

func (t *termboxInput) GetInput(g *game) {
	select {
	case ev := <-t.events:
		termboxInputHandler(g, ev)
	default:
		return
	}
}

func termboxInputHandler(g *game, e termbox.Event) Input {
	switch e.Key {
	case termbox.KeyEsc:
		return InputStart
	case termbox.KeyArrowLeft:
		return InputWest
	case termbox.KeyArrowRight:
		return InputEast
	case termbox.KeyArrowUp:
		return InputNorth
	case termbox.KeyArrowDown:
		return InputSouth
	}
}

func NewTermboxInput() *termboxInput {
	events := make(chan termbox.Event, 1)
	go func(ch chan termbox.Event) {
		ch <- termbox.PollEvent()
	}(events)
	return &termboxInput{events: events}
}
