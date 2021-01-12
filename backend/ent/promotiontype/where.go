// Code generated by entc, DO NOT EDIT.

package promotiontype

import (
	"github.com/G16/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TYPE applies equality check predicate on the "TYPE" field. It's identical to TYPEEQ.
func TYPE(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTYPE), v))
	})
}

// TYPEEQ applies the EQ predicate on the "TYPE" field.
func TYPEEQ(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTYPE), v))
	})
}

// TYPENEQ applies the NEQ predicate on the "TYPE" field.
func TYPENEQ(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTYPE), v))
	})
}

// TYPEIn applies the In predicate on the "TYPE" field.
func TYPEIn(vs ...string) predicate.Promotiontype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Promotiontype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTYPE), v...))
	})
}

// TYPENotIn applies the NotIn predicate on the "TYPE" field.
func TYPENotIn(vs ...string) predicate.Promotiontype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Promotiontype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTYPE), v...))
	})
}

// TYPEGT applies the GT predicate on the "TYPE" field.
func TYPEGT(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTYPE), v))
	})
}

// TYPEGTE applies the GTE predicate on the "TYPE" field.
func TYPEGTE(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTYPE), v))
	})
}

// TYPELT applies the LT predicate on the "TYPE" field.
func TYPELT(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTYPE), v))
	})
}

// TYPELTE applies the LTE predicate on the "TYPE" field.
func TYPELTE(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTYPE), v))
	})
}

// TYPEContains applies the Contains predicate on the "TYPE" field.
func TYPEContains(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTYPE), v))
	})
}

// TYPEHasPrefix applies the HasPrefix predicate on the "TYPE" field.
func TYPEHasPrefix(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTYPE), v))
	})
}

// TYPEHasSuffix applies the HasSuffix predicate on the "TYPE" field.
func TYPEHasSuffix(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTYPE), v))
	})
}

// TYPEEqualFold applies the EqualFold predicate on the "TYPE" field.
func TYPEEqualFold(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTYPE), v))
	})
}

// TYPEContainsFold applies the ContainsFold predicate on the "TYPE" field.
func TYPEContainsFold(v string) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTYPE), v))
	})
}

// HasPromotion applies the HasEdge predicate on the "promotion" edge.
func HasPromotion() predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPromotionWith applies the HasEdge predicate on the "promotion" edge with a given conditions (other predicates).
func HasPromotionWith(preds ...predicate.Promotion) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Promotiontype) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Promotiontype) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
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
func Not(p predicate.Promotiontype) predicate.Promotiontype {
	return predicate.Promotiontype(func(s *sql.Selector) {
		p(s.Not())
	})
}