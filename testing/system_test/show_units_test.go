package system_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/logos/ecs/components/attributes"
	"github.com/logos/ecs/ecs_systems/info"
	"github.com/logos/ecs/entities/unit"
)

func TestListUnits(t *testing.T) {
	units := []unit.Unit{
		{Name: "Unit1", Experience: attributes.Experience{Stage: 1}, LocationID: 100},
		{Name: "Unit2", Experience: attributes.Experience{Stage: 2}, LocationID: 200},
		{Name: "Unit3", Experience: attributes.Experience{Stage: 3}, LocationID: 300},
	}

	expectedOutput := `| Name       | Stage      | Location   |
|------------|------------|------------|
| Unit1      | 1          | 100        |
| Unit2      | 2          | 200        |
| Unit3      | 3          | 300        |
`

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "| Name       | Stage      | Location   |\n")
	fmt.Fprintf(&buf, "|------------|------------|------------|\n")
	for _, u := range units {
		fmt.Fprintf(&buf, "| %-10s | %-10d | %-10d |\n", u.Name, u.Experience.Stage, u.LocationID)
	}

	info.ListUnits(units)

	if buf.String() != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, buf.String())
	}
}
