package attributes

import "fmt"

type WuXingAttributes struct {
	Wood  float64 `json:"wood"`
	Fire  float64 `json:"fire"`
	Earth float64 `json:"earth"`
	Metal float64 `json:"metal"`
	Water float64 `json:"water"`
}

func NewWuXing(wood, fire, earth, metal, water float64) WuXingAttributes {
	return WuXingAttributes{
		Wood:  wood,
		Fire:  fire,
		Earth: earth,
		Metal: metal,
		Water: water,
	}
}

func (w WuXingAttributes) Display() {
	fmt.Printf("Wood: %.2f\n", w.Wood)
	fmt.Printf("Fire: %.2f\n", w.Fire)
	fmt.Printf("Earth: %.2f\n", w.Earth)
	fmt.Printf("Metal: %.2f\n", w.Metal)
	fmt.Printf("Water: %.2f\n", w.Water)
}
