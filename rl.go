package rl

import "time"

const player = 0

// Game holds all the state for a game.
type game struct {
	positions        map[int]V
	actors           map[int]string
	inputController  InputController
	renderController RenderController
	turn             int
	gMap             Map
	done             bool
	err              error
}

// NewGame creates a new game using the controller supplied.
func NewGame(c Controller) *game {
	return &game{
		positions: map[int]V{0: V{1, 1, 0, 0}},
		actors:    map[int]string{0: "player"},
	}
}

func (g *game) Run() error {
	var err error
	lastStep := time.Now()
	for err != nil {
		err = g.Step()
		if time.Since(lastStep) > 15*time.Millisecond {
			time.Sleep(15*time.Millisecond - time.Since(lastStep))
		}
		lastStep = time.Now()
	}
	return err
}

func (g *game) Step() error {
	g.inputController.GetInput(g)
	update(g)
	g.renderController.Render(g)
	if g.err != nil {
		return g.err
	}
	return nil
}

func update(g *game) {
}

// Player holds all the state for the player.
type Player struct {
	Position   V
	Controller Controller
}

// Controller is an interface for getting actions.
type Controller interface {
	Get() Input
}
