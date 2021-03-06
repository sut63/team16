// Code generated by entc, DO NOT EDIT.

package payment

import (
	"time"

	"github.com/G16/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PAYMENTAMOUNT applies equality check predicate on the "PAYMENTAMOUNT" field. It's identical to PAYMENTAMOUNTEQ.
func PAYMENTAMOUNT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PHONENUMBER applies equality check predicate on the "PHONENUMBER" field. It's identical to PHONENUMBEREQ.
func PHONENUMBER(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// EMAIL applies equality check predicate on the "EMAIL" field. It's identical to EMAILEQ.
func EMAIL(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMAIL), v))
	})
}

// PAYMENTDATE applies equality check predicate on the "PAYMENTDATE" field. It's identical to PAYMENTDATEEQ.
func PAYMENTDATE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTAMOUNTEQ applies the EQ predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTNEQ applies the NEQ predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTIn applies the In predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPAYMENTAMOUNT), v...))
	})
}

// PAYMENTAMOUNTNotIn applies the NotIn predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPAYMENTAMOUNT), v...))
	})
}

// PAYMENTAMOUNTGT applies the GT predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTGTE applies the GTE predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTLT applies the LT predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTLTE applies the LTE predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTContains applies the Contains predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTHasPrefix applies the HasPrefix predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTHasSuffix applies the HasSuffix predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTEqualFold applies the EqualFold predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTEqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PAYMENTAMOUNTContainsFold applies the ContainsFold predicate on the "PAYMENTAMOUNT" field.
func PAYMENTAMOUNTContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPAYMENTAMOUNT), v))
	})
}

// PHONENUMBEREQ applies the EQ predicate on the "PHONENUMBER" field.
func PHONENUMBEREQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERNEQ applies the NEQ predicate on the "PHONENUMBER" field.
func PHONENUMBERNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERIn applies the In predicate on the "PHONENUMBER" field.
func PHONENUMBERIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERNotIn applies the NotIn predicate on the "PHONENUMBER" field.
func PHONENUMBERNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERGT applies the GT predicate on the "PHONENUMBER" field.
func PHONENUMBERGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERGTE applies the GTE predicate on the "PHONENUMBER" field.
func PHONENUMBERGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLT applies the LT predicate on the "PHONENUMBER" field.
func PHONENUMBERLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLTE applies the LTE predicate on the "PHONENUMBER" field.
func PHONENUMBERLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContains applies the Contains predicate on the "PHONENUMBER" field.
func PHONENUMBERContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasPrefix applies the HasPrefix predicate on the "PHONENUMBER" field.
func PHONENUMBERHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasSuffix applies the HasSuffix predicate on the "PHONENUMBER" field.
func PHONENUMBERHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBEREqualFold applies the EqualFold predicate on the "PHONENUMBER" field.
func PHONENUMBEREqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContainsFold applies the ContainsFold predicate on the "PHONENUMBER" field.
func PHONENUMBERContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHONENUMBER), v))
	})
}

// EMAILEQ applies the EQ predicate on the "EMAIL" field.
func EMAILEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMAIL), v))
	})
}

// EMAILNEQ applies the NEQ predicate on the "EMAIL" field.
func EMAILNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEMAIL), v))
	})
}

// EMAILIn applies the In predicate on the "EMAIL" field.
func EMAILIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEMAIL), v...))
	})
}

// EMAILNotIn applies the NotIn predicate on the "EMAIL" field.
func EMAILNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEMAIL), v...))
	})
}

// EMAILGT applies the GT predicate on the "EMAIL" field.
func EMAILGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEMAIL), v))
	})
}

// EMAILGTE applies the GTE predicate on the "EMAIL" field.
func EMAILGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEMAIL), v))
	})
}

// EMAILLT applies the LT predicate on the "EMAIL" field.
func EMAILLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEMAIL), v))
	})
}

// EMAILLTE applies the LTE predicate on the "EMAIL" field.
func EMAILLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEMAIL), v))
	})
}

// EMAILContains applies the Contains predicate on the "EMAIL" field.
func EMAILContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEMAIL), v))
	})
}

// EMAILHasPrefix applies the HasPrefix predicate on the "EMAIL" field.
func EMAILHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEMAIL), v))
	})
}

// EMAILHasSuffix applies the HasSuffix predicate on the "EMAIL" field.
func EMAILHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEMAIL), v))
	})
}

// EMAILEqualFold applies the EqualFold predicate on the "EMAIL" field.
func EMAILEqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEMAIL), v))
	})
}

// EMAILContainsFold applies the ContainsFold predicate on the "EMAIL" field.
func EMAILContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEMAIL), v))
	})
}

// PAYMENTDATEEQ applies the EQ predicate on the "PAYMENTDATE" field.
func PAYMENTDATEEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTDATENEQ applies the NEQ predicate on the "PAYMENTDATE" field.
func PAYMENTDATENEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTDATEIn applies the In predicate on the "PAYMENTDATE" field.
func PAYMENTDATEIn(vs ...time.Time) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPAYMENTDATE), v...))
	})
}

// PAYMENTDATENotIn applies the NotIn predicate on the "PAYMENTDATE" field.
func PAYMENTDATENotIn(vs ...time.Time) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPAYMENTDATE), v...))
	})
}

// PAYMENTDATEGT applies the GT predicate on the "PAYMENTDATE" field.
func PAYMENTDATEGT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTDATEGTE applies the GTE predicate on the "PAYMENTDATE" field.
func PAYMENTDATEGTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTDATELT applies the LT predicate on the "PAYMENTDATE" field.
func PAYMENTDATELT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPAYMENTDATE), v))
	})
}

// PAYMENTDATELTE applies the LTE predicate on the "PAYMENTDATE" field.
func PAYMENTDATELTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPAYMENTDATE), v))
	})
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPaymenttype applies the HasEdge predicate on the "paymenttype" edge.
func HasPaymenttype() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymenttypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymenttypeTable, PaymenttypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymenttypeWith applies the HasEdge predicate on the "paymenttype" edge with a given conditions (other predicates).
func HasPaymenttypeWith(preds ...predicate.Paymenttype) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymenttypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymenttypeTable, PaymenttypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPromotion applies the HasEdge predicate on the "promotion" edge.
func HasPromotion() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPromotionWith applies the HasEdge predicate on the "promotion" edge with a given conditions (other predicates).
func HasPromotionWith(preds ...predicate.Promotion) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
