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

func NewBlock(wood, fire, earth, metal, water float64, aura int64) (Block, error) {
	if wood < 0 || fire < 0 || earth < 0 || metal < 0 || water < 0 {
		return Block{}, fmt.Errorf("attributes values cannot be negative")
	}

	if aura == 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if global.WorldAttributeSetting == nil {
			return Block{}, fmt.Errorf("global.WorldAttributeSetting is not initialized")
		}
		aura = int64(global.WorldAttributeSetting.AuraRange.Min + r.Intn(global.WorldAttributeSetting.AuraRange.Max-global.WorldAttributeSetting.AuraRange.Min))
	}

	return Block{
		WuXing: attributes.NewWuXing(wood, fire, earth, metal, water),
		Aura:   aura,
	}, nil
}

func (b Block) Display() {
	b.WuXing.Display()
	fmt.Printf("Aura: %d\n", b.Aura)
}
