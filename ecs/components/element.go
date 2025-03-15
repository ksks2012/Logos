package components

import (
	"fmt"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/internal/utils"
)

type ElementAttributes struct {
	Value       [5]attributes.DiagramsAttributes `json:"value"`
	EncodeValue uint64                           `json:"encode_value"`
}

func DefaultElementAttributes() ElementAttributes {
	return ElementAttributes{
		Value: [5]attributes.DiagramsAttributes{},
	}
}

func NewElementAttributes(value [5]attributes.DiagramsAttributes) ElementAttributes {
	return ElementAttributes{
		Value: value,
	}
}

func (e *ElementAttributes) SetElementValue(value [5]attributes.DiagramsAttributes) {
	e.Value = value
}

func (e *ElementAttributes) GetElementValue() [5]attributes.DiagramsAttributes {
	return e.Value
}

func (e *ElementAttributes) EncodeByDiagrams() {
	e.EncodeValue = 0
	for i, v := range e.Value {
		var tmp = v.Encode()
		e.EncodeValue |= uint64(tmp) << uint(i*8)
	}
}

func (e ElementAttributes) Display() {
	for i, v := range e.Value {
		switch i {
		case 0:
			fmt.Printf("Wood: %v\n", v)
		case 1:
			fmt.Printf("Fire: %v\n", v)
		case 2:
			fmt.Printf("Earth: %v\n", v)
		case 3:
			fmt.Printf("Metal: %v\n", v)
		case 4:
			fmt.Printf("Water: %v\n", v)
		}
	}
	fmt.Printf("EncodeValue: %v\n", utils.FormatBinary(e.EncodeValue, 40))
}
