package bt

// Selector: Selects a node, returns Success if any child node succeeds
type Selector struct {
	children []Node
}

func (s *Selector) Execute(bb *Blackboard) NodeState {
	for _, child := range s.children {
		state := child.Execute(bb)
		if state == Success {
			return Success
		}
	}
	return Failure
}

// Sequence: Sequence node, returns Success only if all child nodes succeed
type Sequence struct {
	children []Node
}

func (s *Sequence) Execute(bb *Blackboard) NodeState {
	for _, child := range s.children {
		state := child.Execute(bb)
		if state != Success {
			return state
		}
	}
	return Success
}
