package initial

import (
	"math/rand"
	"time"

	"github.com/logos/ecs/entities"
	"github.com/logos/global"
)

func InitialWorld() entities.World {
	// Create a new world

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := global.WorldAttributeSetting.BlockSizeRange.Min + r.Intn(global.WorldAttributeSetting.BlockSizeRange.Max-global.WorldAttributeSetting.BlockSizeRange.Min)
	y := global.WorldAttributeSetting.BlockSizeRange.Min + r.Intn(global.WorldAttributeSetting.BlockSizeRange.Max-global.WorldAttributeSetting.BlockSizeRange.Min)
	println("World size: ", x, y)
	world := entities.NewWorld(x, y)
	world.Display()

	return world
}
