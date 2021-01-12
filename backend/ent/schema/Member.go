package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Member schema.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("MEMBERID").
			NotEmpty(),
		field.String("NAME").
			NotEmpty(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payment", Payment.Type),
		edge.To("bookcourse", Bookcourse.Type),	
		edge.To("equipmentrental", Equipmentrental.Type),	
	}
}
