package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Equipmentrental schema.
type Equipmentrental struct {
	ent.Schema
}

// Fields of the Equipmentrental.
func (Equipmentrental) Fields() []ent.Field {
	return []ent.Field{
		field.Int("RENTALAMOUNT"),
		field.Time("RENTALDATE"),
		field.Time("RETURNDATE"),
	}
}

// Edges of the Equipmentrental.
func (Equipmentrental) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("equipmentrental").
			Unique(),
		edge.From("employee", Employee.Type).
			Ref("equipmentrental").
			Unique(),
		edge.From("member", Member.Type).
			Ref("equipmentrental").
			Unique(),
		edge.From("equipmenttype", Equipmenttype.Type).
			Ref("equipmentrental").
			Unique(),
	}
}
