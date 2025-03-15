package attributes

// AlignmentType represents the nine alignments in Dungeons & Dragons.
type AlignmentType int

const (
	LawfulGood AlignmentType = iota
	NeutralGood
	ChaoticGood
	LawfulNeutral
	TrueNeutral
	ChaoticNeutral
	LawfulEvil
	NeutralEvil
	ChaoticEvil
)

type AlignmentAttributes struct {
	Alignment     AlignmentType `json:"alignment"`
	GoodEvilValue float64       `json:"good_evil_value"` // Good-Evil value
	LawChaosValue float64       `json:"law_chaos_value"` // Law-Chaos value
}

// Update alignment type by values
func (a *AlignmentAttributes) UpdateAlignmentType() {
	if a.GoodEvilValue > 0 {
		if a.LawChaosValue > 0 {
			a.Alignment = LawfulGood
		} else if a.LawChaosValue < 0 {
			a.Alignment = ChaoticGood
		} else {
			a.Alignment = NeutralGood
		}
	} else if a.GoodEvilValue < 0 {
		if a.LawChaosValue > 0 {
			a.Alignment = LawfulEvil
		} else if a.LawChaosValue < 0 {
			a.Alignment = ChaoticEvil
		} else {
			a.Alignment = NeutralEvil
		}
	} else {
		if a.LawChaosValue > 0 {
			a.Alignment = LawfulNeutral
		} else if a.LawChaosValue < 0 {
			a.Alignment = ChaoticNeutral
		} else {
			a.Alignment = TrueNeutral
		}
	}
}

func (a AlignmentType) String() string {
	return [...]string{
		"Lawful Good",
		"Neutral Good",
		"Chaotic Good",
		"Lawful Neutral",
		"True Neutral",
		"Chaotic Neutral",
		"Lawful Evil",
		"Neutral Evil",
		"Chaotic Evil",
	}[a]
}
