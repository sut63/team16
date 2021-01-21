package schema

import (
	"regexp"
	"errors"

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
			Validate(func(s string) error {
				match, _ := regexp.Match("[-_=@#$%?]", []byte(s))
					if match {
						return errors.New("มีตัวเลชหรืออักษรพิเศษ")	
					}
				return nil
			}).
			NotEmpty(),
		field.String("DESC").
			MaxLen(30).
			NotEmpty(),
		field.String("CODE").
			Match(regexp.MustCompile("[A-Z]+[A-Z]+[0-9]+[0-9]+[0-9]")).
			MaxLen(5).
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
		edge.From("employee", Employee.Type).
			Ref("promotion").
			Unique(),
		edge.To("payment", Payment.Type),
	}
}
