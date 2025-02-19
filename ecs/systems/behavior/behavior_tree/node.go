package bt

// NodeState defines the return state of a node
type NodeState int

const (
	Success NodeState = iota
	Failure
	Running
)

// Node is the basic interface for behavior tree nodes
type Node interface {
	Execute(*Blackboard) NodeState
}
