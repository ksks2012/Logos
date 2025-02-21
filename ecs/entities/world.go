package entities

import (
	"fmt"
	"math/rand"

	"github.com/logos/ecs/entities/unit"
)

type World struct {
	Blocks    []unit.Block
	Length    int
	Width     int
	TotalAura int64
}

func NewWorld(x, y int) (World, error) {
	if x <= 0 || y <= 0 {
		return World{}, fmt.Errorf("invalid dimensions: x and y must be greater than 0")
	}

	blocks := make([]unit.Block, x*y)

	var auraCount int64 = 0
	var err error = nil
	for i := range blocks {
		blocks[i], err = unit.NewBlock(0, 0, 0, 0, 0, 0)
		auraCount += blocks[i].Aura
	}
	return World{
		Blocks:    blocks,
		Length:    x,
		Width:     y,
		TotalAura: auraCount,
	}, err
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

func (w World) RandomLocation() unit.Block {
	index := rand.Intn(len(w.Blocks))
	return w.Blocks[index]
}
