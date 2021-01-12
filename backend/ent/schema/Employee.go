package schema

import (
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
		field.String("EMPLOYEEID").
			NotEmpty(),
		field.String("EMPLOYEENAME").
			NotEmpty(),
		field.String("EMPLOYEEADDRESS").
			NotEmpty(),
		field.Int("IDCARDNUMBER").
			Positive(),
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
	}
}
