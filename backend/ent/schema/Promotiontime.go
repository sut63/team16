package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Promotiontime schema.
type Promotiontime struct {
	ent.Schema
}

// Fields of the Promotiontime.
func (Promotiontime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("DATE"),
		field.Int("HOUR").
			Positive(),
		field.Int("MINUTE").
			Positive(),
	}
}

// Edges of the Promotiontime.
func (Promotiontime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("promotion", Promotion.Type),
	}
}
