package unit

import (
	"fmt"

	"github.com/logos/ecs/components"
	"github.com/logos/ecs/components/attributes"
)

type Inanimate struct {
	ID          int64                          `json:"id"`
	Name        string                         `json:"name"`
	UnitTypeID  components.UnitType            `json:"unit_type_id"`
	Character   attributes.CharacterAttributes `json:"character"` // For bunos of unit
	Practice    attributes.PracticeAttributes  `json:"practice"`
	Description string                         `json:"description"`
	LocationID  int64                          `json:"location_id"`
	WuXing      attributes.WuXingAttributes    `json:"wuxing"`
	Diagram     attributes.DiagramsAttributes  `json:"diagram"`
}

// NewInanimate creates and initializes a new inanimate object
func NewInanimate(id int64, name string, unitTypeID components.UnitType, character attributes.CharacterAttributes, practice attributes.PracticeAttributes, description string, locationID int64, wuXing attributes.WuXingAttributes, diagram attributes.DiagramsAttributes) Inanimate {
	return Inanimate{
		ID:          id,
		Name:        name,
		UnitTypeID:  unitTypeID,
		Character:   character,
		Practice:    practice,
		Description: description,
		LocationID:  locationID,
		WuXing:      wuXing,
		Diagram:     diagram,
	}
}

// SetLocationID sets the inanimate object's location ID
func (i *Inanimate) SetLocationID(locationID int64) {
	i.LocationID = locationID
}

// Display shows the inanimate object's full attributes
func (i Inanimate) Display() {
	fmt.Printf("Inanimate Object Name: %s\n", i.Name)
	fmt.Printf("Location ID: %d\n", i.LocationID)
	fmt.Printf("Description: %s\n", i.Description)
}
