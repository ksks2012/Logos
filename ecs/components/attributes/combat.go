package attributes

import "fmt"

// CombatAttributes defines the combat attributes of a character
// NOTE: By calculting
type CombatAttributes struct {
	Attack       int     `json:"attack"`       // Attack
	Defense      int     `json:"defense"`      // Defense
	CriticalRate float64 `json:"criticalRate"` // Critical Rate
	Accuracy     float64 `json:"accuracy"`     // Accuracy
	Evasion      float64 `json:"evasion"`      // Evasion
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

func CalculateCombatAttributes(ca CharacterAttributes, pa PracticeAttributes) CombatAttributes {
	// TODO: Value Adjustment
	attack := ca.Strength*2 + ca.QiEnergy*1 + pa.CultivationLevel*2
	defense := ca.Vitality*2 + pa.BodyConstitution*1 + pa.CultivationLevel*1
	criticalRate := (float64(ca.Agility)*5 + float64(pa.LuckFortune)*3) * 1.5
	accuracy := (float64(ca.Intelligence)*10 + float64(ca.Agility)*5) * 1.2
	evasion := (float64(ca.Agility)*5 + float64(pa.LuckFortune)*2) * 1.3

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
