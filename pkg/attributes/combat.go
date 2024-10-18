package attributes

import "fmt"

// CombatAttributes defines the combat attributes of a character
type CombatAttributes struct {
	Attack       int     // Attack
	Defense      int     // Defense
	CriticalRate float64 // Critical Rate
	Accuracy     float64 // Accuracy
	Evasion      float64 // Evasion
}

// NewCombatAttributes creates and initializes a character's combat attributes
func NewCombatAttributes(attack, defense int, criticalRate, accuracy, evasion float64) CombatAttributes {
	return CombatAttributes{
		Attack:       attack,
		Defense:      defense,
		CriticalRate: criticalRate,
		Accuracy:     accuracy,
		Evasion:      evasion,
	}
}

// Display shows the character's combat attributes
func (c CombatAttributes) Display() {
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Critical Rate: %.2f%%\n", c.CriticalRate*100)
	fmt.Printf("Accuracy: %.2f%%\n", c.Accuracy*100)
	fmt.Printf("Evasion: %.2f%%\n", c.Evasion*100)
}
