package schema

import (
	"regexp"
	"errors"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payment schema.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.String("PAYMENTAMOUNT").
			Validate(func(s string) error {
				match, _ := regexp.Match("[0-9]+[B$]", []byte(s))
					if !match {
						return errors.New("รูปแบบราคาไม่ถูกต้อง")
					}
				return nil
			}).
			NotEmpty(),
		field.String("PHONENUMBER").
			MaxLen(10).
			MinLen(10).
			NotEmpty(),
		field.String("EMAIL").
			Validate(func(s string) error {
				match, _ := regexp.Match("[A-Za-z0-9-_]+[@]+[a-z]+[.]+[a-z]", []byte(s))
					if !match {
						return errors.New("รูปแบบ EMAIL ไม่ถูกต้อง")
					}
				return nil
			}).
			NotEmpty(),
		field.Time("PAYMENTDATE"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).
			Ref("payment").
			Unique(),
		edge.From("employee", Employee.Type).
			Ref("payment").
			Unique(),
		edge.From("paymenttype", Paymenttype.Type).
			Ref("payment").
			Unique(),
		edge.From("promotion", Promotion.Type).
			Ref("payment").
			Unique(),
	}
}
