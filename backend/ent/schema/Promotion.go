package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Promotion schema.
type Promotion struct {
	ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME").
			NotEmpty(),
		field.String("DESC").
			NotEmpty(),
		field.String("CODE").
			NotEmpty(),
		field.Time("DATE"),
	}
}

// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("promotiontype", Promotiontype.Type).
			Ref("promotion").
			Unique(),
		edge.From("promotionamount", Promotionamount.Type).
			Ref("promotion").
			Unique(),
		edge.To("payment", Payment.Type),
	}
}
