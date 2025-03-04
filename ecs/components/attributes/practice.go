package attributes

import "fmt"

// PracticeAttributes defines the special attributes of a character
type PracticeAttributes struct {
	SpiritualRoot    string `json:"spiritual_root"`    // Spiritual Root (e.g., Gold, Wood, Water, Fire, Earth, etc.)
	BodyConstitution int    `json:"body_constitution"` // Body Constitution
	Comprehension    int    `json:"comprehension"`     // Comprehension
	LuckFortune      int    `json:"luck_fortune"`      // Luck/Fortune
	Willpower        int    `json:"willpower"`         // Willpower
	CultivationLevel int    `json:"cultivation_level"` // Cultivation Level
}

// NewSpecialAttributes creates and initializes a character's special attributes
func NewSpecialAttributes(spiritualRoot string, bodyConstitution, comprehension, luckFortune, willpower, cultivationLevel int) PracticeAttributes {
	return PracticeAttributes{
		SpiritualRoot:    spiritualRoot,
		BodyConstitution: bodyConstitution,
		Comprehension:    comprehension,
		LuckFortune:      luckFortune,
		Willpower:        willpower,
		CultivationLevel: cultivationLevel,
	}
}

// Display shows the character's special attributes
func (s PracticeAttributes) Display() {
	fmt.Printf("Spiritual Root: %s\n", s.SpiritualRoot)
	fmt.Printf("Body Constitution: %d\n", s.BodyConstitution)
	fmt.Printf("Comprehension: %d\n", s.Comprehension)
	fmt.Printf("Luck/Fortune: %d\n", s.LuckFortune)
	fmt.Printf("Willpower: %d\n", s.Willpower)
	fmt.Printf("Cultivation Level: %d\n", s.CultivationLevel)
}
