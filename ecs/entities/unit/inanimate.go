package unit

import (
	"fmt"

	"github.com/logos/ecs/components/attributes"
)

type Inanimate struct {
	ID          int64                         `json:"id"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	LocationID  int64                         `json:"location_id"`
	WuXing      attributes.WuXingAttributes   `json:"wuxing"`
	Diagram     attributes.DiagramsAttributes `json:"diagram"`
}

// NewInanimate creates and initializes a new inanimate object
func NewInanimate(id int64, name, description string, locationID int64) Inanimate {
	return Inanimate{
		ID:          id,
		Name:        name,
		Description: description,
		LocationID:  locationID,
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
