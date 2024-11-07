package behavior

import (
	"fmt"
	"math/rand"

	"github.com/logos/ecs/entities/unit"
)

// Abstract interface for behavior
type Behavior interface {
	Execute(*unit.Unit) bool
}

// High-level goal behavior
type GoalBehavior struct {
	SubBehaviors []Behavior
}

func (g *GoalBehavior) Execute(u *unit.Unit) bool {
	fmt.Println("Executing high-level goal behavior...")
	for _, subBehavior := range g.SubBehaviors {
		if subBehavior.Execute(u) {
			return true
		}
	}
	return false
}

// Environment sensing behavior
type PerceptionBehavior struct {
	Range int
}

func (p *PerceptionBehavior) Execute(u *unit.Unit) bool {
	fmt.Println("Perceiving environment...")
	if rand.Intn(100) < p.Range {
		fmt.Println("Detected high-energy area, considering exploration...")
		return true
	}
	return false
}

// Behavior influenced by personality
type PersonalityBehavior struct {
	Type string
}

func (p *PersonalityBehavior) Execute(u *unit.Unit) bool {
	fmt.Printf("Personality-based behavior: %s\n", p.Type)
	if p.Type == "explorer" {
		fmt.Println("Explorer Unit prefers unknown areas...")
		return true
	} else if p.Type == "self-centered" {
		fmt.Println("Self-centered Unit chooses low-risk paths...")
		return true
	}
	return false
}

// Dynamic learning behavior
// TODO: Add experience-based decision-making
type LearningBehavior struct {
	Experience int
}

func (l *LearningBehavior) Execute(u *unit.Unit) bool {
	fmt.Println("Evaluating past experience to refine decision-making...")
	if l.Experience > 50 {
		fmt.Println("Experience is high, optimizing actions based on previous outcomes...")
		return true
	}
	return false
}

// Resource and risk management behavior
type ResourceManagementBehavior struct {
	ResourceCost int
	RiskFactor   int
}

func (r *ResourceManagementBehavior) Execute(u *unit.Unit) bool {
	fmt.Println("Evaluating resource cost and risk...")
	if u.Character.Vitality > r.ResourceCost && rand.Intn(100) > r.RiskFactor {
		fmt.Printf("Sufficient resources. Proceeding with action (Cost: %d, Risk: %d)\n", r.ResourceCost, r.RiskFactor)
		return true
	}
	fmt.Println("Insufficient resources or high risk, choosing a safer path...")
	return false
}

// Decision tree execution function
func executeDecisionTree(unit *unit.Unit) {
	// Initialize behaviors at each level as needed
	goalBehavior := &GoalBehavior{
		SubBehaviors: []Behavior{
			&PerceptionBehavior{Range: 75},
			&PersonalityBehavior{Type: "explorer"},
			&LearningBehavior{Experience: 60},
			&ResourceManagementBehavior{ResourceCost: 20, RiskFactor: 30},
		},
	}

	// Execute the high-level decision tree behavior
	if goalBehavior.Execute(unit) {
		fmt.Println("Action successfully executed.")
	} else {
		fmt.Println("Failed to find a suitable action.")
	}
}
