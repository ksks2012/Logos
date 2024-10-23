package attributes

import "fmt"

type Experience struct {
	Level int
	Stage int
}

func NewExperience(level int, stage int) Experience {
	return Experience{
		Level: level,
		Stage: stage,
	}
}

func (e Experience) Display() {
	fmt.Printf("Level: %d\n", e.Level)
	fmt.Printf("Stage: %d\n", e.Stage)
}