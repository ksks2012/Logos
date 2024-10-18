package attributes

import "fmt"

// CharacterAttributes defines the basic attributes of a character
type CharacterAttributes struct {
	Vitality     int // Vitality
	QiEnergy     int // Qi/Energy
	Strength     int // Strength
	Agility      int // Agility
	Intelligence int // Intelligence
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
