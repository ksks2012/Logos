package unit

import (
	"github.com/logos/ecs/components/attributes"
)

type Block struct {
	WuXing attributes.WuXingAttributes
}

func NewBlock(wood, fire, earth, metal, water float64) Block {
	return Block{
		WuXing: attributes.NewWuXing(wood, fire, earth, metal, water),
	}
}

func (b Block) Display() {
	b.WuXing.Display()
}
