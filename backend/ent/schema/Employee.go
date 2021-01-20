package schema

import (
	"errors"
	"regexp"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Employee schema.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("EMPLOYEEID").NotEmpty().Unique().Validate(func(s string) error {
			match, _ := regexp.MatchString("[E][M]\\d{2}", s)
			if !match {
				return errors.New("รูปแบบรหัสของพนักงานไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("EMPLOYEENAME").NotEmpty(),
		field.String("EMPLOYEEADDRESS").NotEmpty(),
		field.String("IDCARDNUMBER").NotEmpty().Unique().MaxLen(13).MinLen(13),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("age", Age.Type).
			Ref("employee").
			Unique(),
		edge.From("position", Position.Type).
			Ref("employee").
			Unique(),
		edge.From("salary", Salary.Type).
			Ref("employee").
			Unique(),
		edge.To("payment", Payment.Type),
		edge.To("equipment", Equipment.Type),
		edge.To("bookcourse", Bookcourse.Type),
		edge.To("equipmentrental", Equipmentrental.Type),
		edge.To("promotion", Promotion.Type),
	}
}
