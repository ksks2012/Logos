package unit

import (
	"fmt"

	"github.com/logos/ecs/components/attributes"
)

// Unit represents a complete character with all attributes
type Unit struct {
	Name       string
	Character  attributes.CharacterAttributes
	Practice   attributes.PracticeAttributes
	Combat     attributes.CombatAttributes
	Experience attributes.Experience
}

// NewUnit creates and initializes a new unit with all attributes
func NewUnit(name string, charAttr attributes.CharacterAttributes, pracAttr attributes.PracticeAttributes, combAttr attributes.CombatAttributes) Unit {
	return Unit{
		Name:      name,
		Character: charAttr,
		Practice:  pracAttr,
		Combat:    combAttr,
	}
}

// Display shows the unit's full attributes
func (u Unit) Display() {
	fmt.Printf("Unit Name: %s\n", u.Name)
	fmt.Println("Basic Attributes:")
	fmt.Printf("  Vitality: %d, Qi/Energy: %d, Strength: %d, Agility: %d, Intelligence: %d\n",
		u.Character.Vitality, u.Character.QiEnergy, u.Character.Strength, u.Character.Agility, u.Character.Intelligence)
	fmt.Println("Practice Attributes:")
	fmt.Printf("  Spiritual Root: %s, Body Constitution: %d, Comprehension: %d, Luck/Fortune: %d, Willpower: %d, Cultivation Level: %d\n",
		u.Practice.SpiritualRoot, u.Practice.BodyConstitution, u.Practice.Comprehension, u.Practice.LuckFortune, u.Practice.Willpower, u.Practice.CultivationLevel)
	fmt.Println("Combat Attributes:")
	fmt.Printf("  Attack: %d, Defense: %d, Critical Rate: %.2f%%, Accuracy: %.2f%%, Evasion: %.2f%%\n",
		u.Combat.Attack, u.Combat.Defense, u.Combat.CriticalRate*100, u.Combat.Accuracy*100, u.Combat.Evasion*100)
}
