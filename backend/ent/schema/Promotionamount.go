package schema

import(
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Promotionamount schema.
type Promotionamount struct {
	ent.Schema
}

// Fields of the Promotionamount.
func (Promotionamount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("AMOUNT").
			Positive(),
	}
}

// Edges of the Promotionamount.
func (Promotionamount) Edges() []ent.Edge {
	return  []ent.Edge{
		edge.To("promotion", Promotion.Type),
	}
}
