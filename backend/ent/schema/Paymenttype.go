package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Paymenttype schema.
type Paymenttype struct {
	ent.Schema
}

// Fields of the Paymenttype.
func (Paymenttype) Fields() []ent.Field {
	return []ent.Field{
		field.String("TYPE"),
	}
}

// Edges of the Paymenttype.
func (Paymenttype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payment", Payment.Type),
	}
}
