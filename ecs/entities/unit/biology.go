package unit

import (
	"fmt"

	"github.com/logos/ecs/components/attributes"
)

// Unit represents a complete character with all attributes
type Unit struct {
	ID         int64                          `json:"id"`
	Name       string                         `json:"name"`
	Character  attributes.CharacterAttributes `json:"character"`
	Practice   attributes.PracticeAttributes  `json:"practice"`
	Combat     attributes.CombatAttributes    `json:"combat"`
	Experience attributes.Experience          `json:"experience"`
	Relations  attributes.RelationAttributes  `json:"relations"`
	LocationID int64                          `json:"location_id"`
	IsAlive    bool                           `json:"is_alive"`
}

// TODO: Add attributes for temporary status effects

// NewUnit creates and initializes a new unit with all attributes
func NewUnit(id int64, name string, charAttr attributes.CharacterAttributes, pracAttr attributes.PracticeAttributes, combAttr attributes.CombatAttributes, LocationID int64) Unit {
	return Unit{
		ID:         id,
		Name:       name,
		Character:  charAttr,
		Practice:   pracAttr,
		Combat:     combAttr,
		LocationID: LocationID,
		IsAlive:    true,
	}
}

// SetLocationID sets the unit's location ID
func (u *Unit) SetLocationID(locationID int64) {
	u.LocationID = locationID
}

// Display shows the unit's full attributes
func (u Unit) Display() {
	fmt.Printf("Unit Name: %s\n", u.Name)
	fmt.Printf("Location ID: %d\n", u.LocationID)
	fmt.Printf("Is Alive: %t\n", u.IsAlive)
	fmt.Println("Basic Attributes:")
	fmt.Printf("  Vitality: %d, Qi/Energy: %d, Strength: %d, Agility: %d, Intelligence: %d\n",
		u.Character.Vitality, u.Character.QiEnergy, u.Character.Strength, u.Character.Agility, u.Character.Intelligence)
	fmt.Println("Practice Attributes:")
	fmt.Printf("  Spiritual Root: %s, Body Constitution: %d, Comprehension: %d, Luck/Fortune: %d, Willpower: %d, Cultivation Level: %d\n",
		u.Practice.SpiritualRoot, u.Practice.BodyConstitution, u.Practice.Comprehension, u.Practice.LuckFortune, u.Practice.Willpower, u.Practice.CultivationLevel)
	fmt.Println("Combat Attributes:")
	fmt.Printf("  Attack: %d, Defense: %d, Critical Rate: %.2f%%, Accuracy: %.2f%%, Evasion: %.2f%%\n",
		u.Combat.Attack, u.Combat.Defense, u.Combat.CriticalRate, u.Combat.Accuracy, u.Combat.Evasion)
}
