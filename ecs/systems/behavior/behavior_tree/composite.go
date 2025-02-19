package bt

// Selector: Selects a node, returns Success if any child node succeeds
type Selector struct {
	Children []Node
}

func (s *Selector) Execute(bb *Blackboard) NodeState {
	for _, child := range s.Children {
		state := child.Execute(bb)
		if state == SUCCESS {
			return SUCCESS
		}
	}
	return FAILURE
}

// Sequence: Sequence node, returns Success only if all child nodes succeed
type Sequence struct {
	Children []Node
}

func (s *Sequence) Execute(bb *Blackboard) NodeState {
	for _, child := range s.Children {
		state := child.Execute(bb)
		if state != SUCCESS {
			return state
		}
	}
	return SUCCESS
}
