package bt

import "fmt"

// MoveToTarget: Simulate the behavior of moving to a target
type MoveToTarget struct{}

func (m *MoveToTarget) Execute(bb *Blackboard) NodeState {
	fmt.Println("Moving to target location")
	return Success
}

// SearchForResources: Simulate the behavior of searching for resources
type SearchForResources struct{}

func (s *SearchForResources) Execute(bb *Blackboard) NodeState {
	fmt.Println("Searching for resources...")
	return Success
}

// RestIfTired: Rest if stamina is low
type RestIfTired struct{}

func (r *RestIfTired) Execute(bb *Blackboard) NodeState {
	if bb.Get("stamina").(int) < 30 {
		fmt.Println("Stamina is low, resting...")
		return Success
	}
	return Failure
}
