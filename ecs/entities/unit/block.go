package unit

import (
	"fmt"

	"github.com/logos/ecs/components/attributes"
)

type Block struct {
	WuXing attributes.WuXingAttributes
	Aura   int64
}

func NewBlock(wood, fire, earth, metal, water float64, aura int64) Block {
	return Block{
		WuXing: attributes.NewWuXing(wood, fire, earth, metal, water),
		Aura:   aura,
	}
}

func (b Block) Display() {
	b.WuXing.Display()
	fmt.Printf("Aura: %d\n", b.Aura)
}
