package schema

import (
	"errors"
	"regexp"

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
		field.String("EQUIPMENTNAME").
			Validate(func(s string) error {
				match, _ := regexp.Match("[0-9?+=-_!@#$%^&*<>:;]", []byte(s))
				if match {
					return errors.New("มีตัวเลชหรืออักษรพิเศษ")
				}
				return nil
			}).
			NotEmpty(),
		field.Int("EQUIPMENTAMOUNT").
			Positive(),
		field.String("EQUIPMENTDETAIL").
			MaxLen(100).
			NotEmpty(),
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
