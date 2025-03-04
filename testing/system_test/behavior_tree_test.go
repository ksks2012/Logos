package system_test

import (
	"testing"

	bt "github.com/logos/ecs/ecs_systems/behavior/behavior_tree"
)

func TestBehaviorTree_SearchForResources(t *testing.T) {
	bb := bt.NewBlackboard()
	bb.Set("stamina", 30) // Set NPC stamina above threshold

	behaviorTree := &bt.Selector{
		Children: []bt.Node{
			&bt.RestIfTired{},
			&bt.Sequence{
				Children: []bt.Node{
					&bt.SearchForResources{},
					&bt.MoveToTarget{},
				},
			},
		},
	}

	result := behaviorTree.Execute(bb)
	if result != bt.SUCCESS {
		t.Errorf("Expected SUCCESS, got %v", result)
	}
}

func TestBehaviorTree_MoveToTarget(t *testing.T) {
	bb := bt.NewBlackboard()
	bb.Set("stamina", 30)       // Set NPC stamina above threshold
	bb.Set("targetFound", true) // Set target found

	behaviorTree := &bt.Selector{
		Children: []bt.Node{
			&bt.RestIfTired{},
			&bt.Sequence{
				Children: []bt.Node{
					&bt.SearchForResources{},
					&bt.MoveToTarget{},
				},
			},
		},
	}

	result := behaviorTree.Execute(bb)
	if result != bt.SUCCESS {
		t.Errorf("Expected SUCCESS, got %v", result)
	}
}
