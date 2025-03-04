package bt

// Blackboard: stores shared data for NPC states
type Blackboard struct {
	data map[string]interface{}
}

func NewBlackboard() *Blackboard {
	return &Blackboard{data: make(map[string]interface{})}
}

func (b *Blackboard) Set(key string, value interface{}) {
	b.data[key] = value
}

func (b *Blackboard) Get(key string) interface{} {
	return b.data[key]
}
