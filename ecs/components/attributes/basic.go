package attributes

import "fmt"

// CharacterAttributes defines the basic attributes of a character
type CharacterAttributes struct {
	Vitality     int `json:"vitality"`     // Vitality
	QiEnergy     int `json:"qi_energy"`    // Qi/Energy
	Strength     int `json:"strength"`     // Strength
	Agility      int `json:"agility"`      // Agility
	Intelligence int `json:"intelligence"` // Intelligence
}

// NewCharacterAttributes creates and initializes a character's attributes
func NewCharacterAttributes(vitality, qiEnergy, strength, agility, intelligence int) CharacterAttributes {
	return CharacterAttributes{
		Vitality:     vitality,
		QiEnergy:     qiEnergy,
		Strength:     strength,
		Agility:      agility,
		Intelligence: intelligence,
	}
}

// Display shows the character's attributes
func (c CharacterAttributes) Display() {
	fmt.Printf("Vitality: %d\n", c.Vitality)
	fmt.Printf("Qi/Energy: %d\n", c.QiEnergy)
	fmt.Printf("Strength: %d\n", c.Strength)
	fmt.Printf("Agility: %d\n", c.Agility)
	fmt.Printf("Intelligence: %d\n", c.Intelligence)
}
