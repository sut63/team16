package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Promotiontype schema.
type Promotiontype struct {
	ent.Schema
}

// Fields of the Promotiontype.
func (Promotiontype) Fields() []ent.Field {
	return []ent.Field{
		field.String("TYPE").
			NotEmpty(),
	}
}

// Edges of the Promotiontype.
func (Promotiontype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("promotion", Promotion.Type),
	}
}
