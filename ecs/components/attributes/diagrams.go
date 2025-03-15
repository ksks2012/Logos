package attributes

import "fmt"

type DiagramsAttributes struct {
	Sky      float64 `json:"sky"`
	Earth    float64 `json:"earth"`
	Thunder  float64 `json:"thunder"`
	Wind     float64 `json:"wind"`
	Water    float64 `json:"water"`
	Fire     float64 `json:"fire"`
	Mountain float64 `json:"mountain"`
	Marsh    float64 `json:"marsh"`
}

func NewDiagrams(sky, earth, thunder, wind, water, fire, mountain, marsh float64) DiagramsAttributes {
	return DiagramsAttributes{
		Sky:      sky,
		Earth:    earth,
		Thunder:  thunder,
		Wind:     wind,
		Water:    water,
		Fire:     fire,
		Mountain: mountain,
		Marsh:    marsh,
	}
}

func (d DiagramsAttributes) Display() {
	fmt.Printf("Sky: %.2f\n", d.Sky)
	fmt.Printf("Earth: %.2f\n", d.Earth)
	fmt.Printf("Thunder: %.2f\n", d.Thunder)
	fmt.Printf("Wind: %.2f\n", d.Wind)
	fmt.Printf("Water: %.2f\n", d.Water)
	fmt.Printf("Fire: %.2f\n", d.Fire)
	fmt.Printf("Mountain: %.2f\n", d.Mountain)
	fmt.Printf("Marsh: %.2f\n", d.Marsh)
}
