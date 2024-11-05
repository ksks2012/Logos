package entities

import (
	"github.com/logos/ecs/entities/unit"
)

type World struct {
	Blocks []unit.Block
}

func NewWorld(x, y int) World {
	blocks := make([]unit.Block, x*y)
	return World{
		Blocks: blocks,
	}
}

func (w World) Display() {
	for _, b := range w.Blocks {
		b.Display()
	}
}
