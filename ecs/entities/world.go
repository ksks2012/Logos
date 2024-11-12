package entities

import (
	"fmt"

	"github.com/logos/ecs/entities/unit"
)

type World struct {
	Blocks    []unit.Block
	Length    int
	Width     int
	TotalAura int64
}

func NewWorld(x, y int) World {
	blocks := make([]unit.Block, x*y)

	var auraCount int64 = 0
	for i := range blocks {
		blocks[i] = unit.NewBlock(0, 0, 0, 0, 0, 0)
		auraCount += blocks[i].Aura
	}
	return World{
		Blocks:    blocks,
		Length:    x,
		Width:     y,
		TotalAura: auraCount,
	}
}

func (w World) Display() {
	for i, b := range w.Blocks {
		fmt.Printf("%d\n", i)
		b.Display()
	}
	fmt.Printf("Length: %d\n", w.Length)
	fmt.Printf("Width: %d\n", w.Width)
	fmt.Printf("Total Aura: %d\n", w.TotalAura)
}
