package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Equipment schema.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("EQUIPMENTNAME"),
		field.String("EQUIPMENTAMOUNT"),
		field.String("EQUIPMENTDETAIL"),
		field.Time("EQUIPMENTDATE"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classifier", Classifier.Type).
			Ref("equipment").
			Unique(),
		edge.From("employee", Employee.Type).
			Ref("equipment").
			Unique(),
		edge.From("equipmenttype", Equipmenttype.Type).
			Ref("equipment").
			Unique(),
		edge.From("zone", Zone.Type).
			Ref("equipment").
			Unique(),
		edge.To("equipmentrental", Equipmentrental.Type),	
	}
}
