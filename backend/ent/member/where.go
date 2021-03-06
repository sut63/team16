// Code generated by entc, DO NOT EDIT.

package member

import (
	"github.com/G16/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MEMBERID applies equality check predicate on the "MEMBERID" field. It's identical to MEMBERIDEQ.
func MEMBERID(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMEMBERID), v))
	})
}

// MEMBERNAME applies equality check predicate on the "MEMBERNAME" field. It's identical to MEMBERNAMEEQ.
func MEMBERNAME(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERIDEQ applies the EQ predicate on the "MEMBERID" field.
func MEMBERIDEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDNEQ applies the NEQ predicate on the "MEMBERID" field.
func MEMBERIDNEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDIn applies the In predicate on the "MEMBERID" field.
func MEMBERIDIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMEMBERID), v...))
	})
}

// MEMBERIDNotIn applies the NotIn predicate on the "MEMBERID" field.
func MEMBERIDNotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMEMBERID), v...))
	})
}

// MEMBERIDGT applies the GT predicate on the "MEMBERID" field.
func MEMBERIDGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDGTE applies the GTE predicate on the "MEMBERID" field.
func MEMBERIDGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDLT applies the LT predicate on the "MEMBERID" field.
func MEMBERIDLT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDLTE applies the LTE predicate on the "MEMBERID" field.
func MEMBERIDLTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDContains applies the Contains predicate on the "MEMBERID" field.
func MEMBERIDContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDHasPrefix applies the HasPrefix predicate on the "MEMBERID" field.
func MEMBERIDHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDHasSuffix applies the HasSuffix predicate on the "MEMBERID" field.
func MEMBERIDHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDEqualFold applies the EqualFold predicate on the "MEMBERID" field.
func MEMBERIDEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMEMBERID), v))
	})
}

// MEMBERIDContainsFold applies the ContainsFold predicate on the "MEMBERID" field.
func MEMBERIDContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMEMBERID), v))
	})
}

// MEMBERNAMEEQ applies the EQ predicate on the "MEMBERNAME" field.
func MEMBERNAMEEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMENEQ applies the NEQ predicate on the "MEMBERNAME" field.
func MEMBERNAMENEQ(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEIn applies the In predicate on the "MEMBERNAME" field.
func MEMBERNAMEIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMEMBERNAME), v...))
	})
}

// MEMBERNAMENotIn applies the NotIn predicate on the "MEMBERNAME" field.
func MEMBERNAMENotIn(vs ...string) predicate.Member {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Member(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMEMBERNAME), v...))
	})
}

// MEMBERNAMEGT applies the GT predicate on the "MEMBERNAME" field.
func MEMBERNAMEGT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEGTE applies the GTE predicate on the "MEMBERNAME" field.
func MEMBERNAMEGTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMELT applies the LT predicate on the "MEMBERNAME" field.
func MEMBERNAMELT(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMELTE applies the LTE predicate on the "MEMBERNAME" field.
func MEMBERNAMELTE(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEContains applies the Contains predicate on the "MEMBERNAME" field.
func MEMBERNAMEContains(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEHasPrefix applies the HasPrefix predicate on the "MEMBERNAME" field.
func MEMBERNAMEHasPrefix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEHasSuffix applies the HasSuffix predicate on the "MEMBERNAME" field.
func MEMBERNAMEHasSuffix(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEEqualFold applies the EqualFold predicate on the "MEMBERNAME" field.
func MEMBERNAMEEqualFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMEMBERNAME), v))
	})
}

// MEMBERNAMEContainsFold applies the ContainsFold predicate on the "MEMBERNAME" field.
func MEMBERNAMEContainsFold(v string) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMEMBERNAME), v))
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookcourse applies the HasEdge predicate on the "bookcourse" edge.
func HasBookcourse() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookcourseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookcourseTable, BookcourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookcourseWith applies the HasEdge predicate on the "bookcourse" edge with a given conditions (other predicates).
func HasBookcourseWith(preds ...predicate.Bookcourse) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookcourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookcourseTable, BookcourseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipmentrental applies the HasEdge predicate on the "equipmentrental" edge.
func HasEquipmentrental() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentrentalTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EquipmentrentalTable, EquipmentrentalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentrentalWith applies the HasEdge predicate on the "equipmentrental" edge with a given conditions (other predicates).
func HasEquipmentrentalWith(preds ...predicate.Equipmentrental) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentrentalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EquipmentrentalTable, EquipmentrentalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
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
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		p(s.Not())
	})
}
