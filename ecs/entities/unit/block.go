package unit

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/global"
)

type Block struct {
	WuXing attributes.WuXingAttributes
	Aura   int64
}

func NewBlock(wood, fire, earth, metal, water float64, aura int64) Block {
	if aura == 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		aura = int64(global.WorldAttributeSetting.AuraRange.Min + r.Intn(global.WorldAttributeSetting.AuraRange.Max-global.WorldAttributeSetting.AuraRange.Min))
	}
	return Block{
		WuXing: attributes.NewWuXing(wood, fire, earth, metal, water),
		Aura:   aura,
	}
}

func (b Block) Display() {
	b.WuXing.Display()
	fmt.Printf("Aura: %d\n", b.Aura)
}
