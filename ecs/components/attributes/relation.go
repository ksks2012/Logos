package attributes

type RelationType int

const (
	RelationTypeNone RelationType = iota
	RelationTypeFriendship
	RelationTypeHatred
	RelationTypeMasterDisciple
	RelationTypeFaction
	RelationTypeFamilyTies
	RelationTypeMarriage
)

type RelationLink struct {
	Value float64 `json:"value"` // Relation value
	Type  int     `json:"type"`  // Relation type
}

type RelationAttributes struct {
	Relation map[int]map[int]RelationLink `json:"relation_link"` // Relation map of unit ID to relation value
}

// NewRelationAttributes creates and initializes a character's relation attributes
func NewRelationAttributes() RelationAttributes {
	return RelationAttributes{
		Relation: make(map[int]map[int]RelationLink),
	}
}
