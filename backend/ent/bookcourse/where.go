// Code generated by entc, DO NOT EDIT.

package bookcourse

import (
	"time"

	"github.com/G16/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ACCESS applies equality check predicate on the "ACCESS" field. It's identical to ACCESSEQ.
func ACCESS(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldACCESS), v))
	})
}

// PHONE applies equality check predicate on the "PHONE" field. It's identical to PHONEEQ.
func PHONE(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONE), v))
	})
}

// DETAIL applies equality check predicate on the "DETAIL" field. It's identical to DETAILEQ.
func DETAIL(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDETAIL), v))
	})
}

// BOOKTIME applies equality check predicate on the "BOOKTIME" field. It's identical to BOOKTIMEEQ.
func BOOKTIME(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKTIME), v))
	})
}

// ACCESSEQ applies the EQ predicate on the "ACCESS" field.
func ACCESSEQ(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldACCESS), v))
	})
}

// ACCESSNEQ applies the NEQ predicate on the "ACCESS" field.
func ACCESSNEQ(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldACCESS), v))
	})
}

// ACCESSIn applies the In predicate on the "ACCESS" field.
func ACCESSIn(vs ...int) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldACCESS), v...))
	})
}

// ACCESSNotIn applies the NotIn predicate on the "ACCESS" field.
func ACCESSNotIn(vs ...int) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldACCESS), v...))
	})
}

// ACCESSGT applies the GT predicate on the "ACCESS" field.
func ACCESSGT(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldACCESS), v))
	})
}

// ACCESSGTE applies the GTE predicate on the "ACCESS" field.
func ACCESSGTE(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldACCESS), v))
	})
}

// ACCESSLT applies the LT predicate on the "ACCESS" field.
func ACCESSLT(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldACCESS), v))
	})
}

// ACCESSLTE applies the LTE predicate on the "ACCESS" field.
func ACCESSLTE(v int) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldACCESS), v))
	})
}

// PHONEEQ applies the EQ predicate on the "PHONE" field.
func PHONEEQ(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONE), v))
	})
}

// PHONENEQ applies the NEQ predicate on the "PHONE" field.
func PHONENEQ(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHONE), v))
	})
}

// PHONEIn applies the In predicate on the "PHONE" field.
func PHONEIn(vs ...string) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHONE), v...))
	})
}

// PHONENotIn applies the NotIn predicate on the "PHONE" field.
func PHONENotIn(vs ...string) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHONE), v...))
	})
}

// PHONEGT applies the GT predicate on the "PHONE" field.
func PHONEGT(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHONE), v))
	})
}

// PHONEGTE applies the GTE predicate on the "PHONE" field.
func PHONEGTE(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHONE), v))
	})
}

// PHONELT applies the LT predicate on the "PHONE" field.
func PHONELT(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHONE), v))
	})
}

// PHONELTE applies the LTE predicate on the "PHONE" field.
func PHONELTE(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHONE), v))
	})
}

// PHONEContains applies the Contains predicate on the "PHONE" field.
func PHONEContains(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHONE), v))
	})
}

// PHONEHasPrefix applies the HasPrefix predicate on the "PHONE" field.
func PHONEHasPrefix(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHONE), v))
	})
}

// PHONEHasSuffix applies the HasSuffix predicate on the "PHONE" field.
func PHONEHasSuffix(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHONE), v))
	})
}

// PHONEEqualFold applies the EqualFold predicate on the "PHONE" field.
func PHONEEqualFold(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHONE), v))
	})
}

// PHONEContainsFold applies the ContainsFold predicate on the "PHONE" field.
func PHONEContainsFold(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHONE), v))
	})
}

// DETAILEQ applies the EQ predicate on the "DETAIL" field.
func DETAILEQ(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDETAIL), v))
	})
}

// DETAILNEQ applies the NEQ predicate on the "DETAIL" field.
func DETAILNEQ(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDETAIL), v))
	})
}

// DETAILIn applies the In predicate on the "DETAIL" field.
func DETAILIn(vs ...string) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDETAIL), v...))
	})
}

// DETAILNotIn applies the NotIn predicate on the "DETAIL" field.
func DETAILNotIn(vs ...string) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDETAIL), v...))
	})
}

// DETAILGT applies the GT predicate on the "DETAIL" field.
func DETAILGT(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDETAIL), v))
	})
}

// DETAILGTE applies the GTE predicate on the "DETAIL" field.
func DETAILGTE(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDETAIL), v))
	})
}

// DETAILLT applies the LT predicate on the "DETAIL" field.
func DETAILLT(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDETAIL), v))
	})
}

// DETAILLTE applies the LTE predicate on the "DETAIL" field.
func DETAILLTE(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDETAIL), v))
	})
}

// DETAILContains applies the Contains predicate on the "DETAIL" field.
func DETAILContains(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDETAIL), v))
	})
}

// DETAILHasPrefix applies the HasPrefix predicate on the "DETAIL" field.
func DETAILHasPrefix(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDETAIL), v))
	})
}

// DETAILHasSuffix applies the HasSuffix predicate on the "DETAIL" field.
func DETAILHasSuffix(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDETAIL), v))
	})
}

// DETAILEqualFold applies the EqualFold predicate on the "DETAIL" field.
func DETAILEqualFold(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDETAIL), v))
	})
}

// DETAILContainsFold applies the ContainsFold predicate on the "DETAIL" field.
func DETAILContainsFold(v string) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDETAIL), v))
	})
}

// BOOKTIMEEQ applies the EQ predicate on the "BOOKTIME" field.
func BOOKTIMEEQ(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBOOKTIME), v))
	})
}

// BOOKTIMENEQ applies the NEQ predicate on the "BOOKTIME" field.
func BOOKTIMENEQ(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBOOKTIME), v))
	})
}

// BOOKTIMEIn applies the In predicate on the "BOOKTIME" field.
func BOOKTIMEIn(vs ...time.Time) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBOOKTIME), v...))
	})
}

// BOOKTIMENotIn applies the NotIn predicate on the "BOOKTIME" field.
func BOOKTIMENotIn(vs ...time.Time) predicate.Bookcourse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bookcourse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBOOKTIME), v...))
	})
}

// BOOKTIMEGT applies the GT predicate on the "BOOKTIME" field.
func BOOKTIMEGT(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBOOKTIME), v))
	})
}

// BOOKTIMEGTE applies the GTE predicate on the "BOOKTIME" field.
func BOOKTIMEGTE(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBOOKTIME), v))
	})
}

// BOOKTIMELT applies the LT predicate on the "BOOKTIME" field.
func BOOKTIMELT(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBOOKTIME), v))
	})
}

// BOOKTIMELTE applies the LTE predicate on the "BOOKTIME" field.
func BOOKTIMELTE(v time.Time) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBOOKTIME), v))
	})
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
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

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
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

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Bookcourse) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Bookcourse) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
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
func Not(p predicate.Bookcourse) predicate.Bookcourse {
	return predicate.Bookcourse(func(s *sql.Selector) {
		p(s.Not())
	})
}
