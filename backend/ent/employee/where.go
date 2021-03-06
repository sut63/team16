// Code generated by entc, DO NOT EDIT.

package employee

import (
	"github.com/G16/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EMPLOYEEID applies equality check predicate on the "EMPLOYEEID" field. It's identical to EMPLOYEEIDEQ.
func EMPLOYEEID(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEENAME applies equality check predicate on the "EMPLOYEENAME" field. It's identical to EMPLOYEENAMEEQ.
func EMPLOYEENAME(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEEADDRESS applies equality check predicate on the "EMPLOYEEADDRESS" field. It's identical to EMPLOYEEADDRESSEQ.
func EMPLOYEEADDRESS(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// IDCARDNUMBER applies equality check predicate on the "IDCARDNUMBER" field. It's identical to IDCARDNUMBEREQ.
func IDCARDNUMBER(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIDCARDNUMBER), v))
	})
}

// EMPLOYEEIDEQ applies the EQ predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDNEQ applies the NEQ predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDIn applies the In predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEMPLOYEEID), v...))
	})
}

// EMPLOYEEIDNotIn applies the NotIn predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEMPLOYEEID), v...))
	})
}

// EMPLOYEEIDGT applies the GT predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDGTE applies the GTE predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDLT applies the LT predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDLTE applies the LTE predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDContains applies the Contains predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDHasPrefix applies the HasPrefix predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDHasSuffix applies the HasSuffix predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDEqualFold applies the EqualFold predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEEIDContainsFold applies the ContainsFold predicate on the "EMPLOYEEID" field.
func EMPLOYEEIDContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEMPLOYEEID), v))
	})
}

// EMPLOYEENAMEEQ applies the EQ predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMENEQ applies the NEQ predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMENEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEIn applies the In predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEMPLOYEENAME), v...))
	})
}

// EMPLOYEENAMENotIn applies the NotIn predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMENotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEMPLOYEENAME), v...))
	})
}

// EMPLOYEENAMEGT applies the GT predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEGTE applies the GTE predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMELT applies the LT predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMELT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMELTE applies the LTE predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMELTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEContains applies the Contains predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEHasPrefix applies the HasPrefix predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEHasSuffix applies the HasSuffix predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEEqualFold applies the EqualFold predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEENAMEContainsFold applies the ContainsFold predicate on the "EMPLOYEENAME" field.
func EMPLOYEENAMEContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEMPLOYEENAME), v))
	})
}

// EMPLOYEEADDRESSEQ applies the EQ predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSNEQ applies the NEQ predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSIn applies the In predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEMPLOYEEADDRESS), v...))
	})
}

// EMPLOYEEADDRESSNotIn applies the NotIn predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEMPLOYEEADDRESS), v...))
	})
}

// EMPLOYEEADDRESSGT applies the GT predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSGTE applies the GTE predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSLT applies the LT predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSLTE applies the LTE predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSContains applies the Contains predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSHasPrefix applies the HasPrefix predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSHasSuffix applies the HasSuffix predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSEqualFold applies the EqualFold predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// EMPLOYEEADDRESSContainsFold applies the ContainsFold predicate on the "EMPLOYEEADDRESS" field.
func EMPLOYEEADDRESSContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEMPLOYEEADDRESS), v))
	})
}

// IDCARDNUMBEREQ applies the EQ predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBEREQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERNEQ applies the NEQ predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERIn applies the In predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIDCARDNUMBER), v...))
	})
}

// IDCARDNUMBERNotIn applies the NotIn predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIDCARDNUMBER), v...))
	})
}

// IDCARDNUMBERGT applies the GT predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERGTE applies the GTE predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERLT applies the LT predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERLTE applies the LTE predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERContains applies the Contains predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERHasPrefix applies the HasPrefix predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERHasSuffix applies the HasSuffix predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBEREqualFold applies the EqualFold predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBEREqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIDCARDNUMBER), v))
	})
}

// IDCARDNUMBERContainsFold applies the ContainsFold predicate on the "IDCARDNUMBER" field.
func IDCARDNUMBERContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIDCARDNUMBER), v))
	})
}

// HasAge applies the HasEdge predicate on the "age" edge.
func HasAge() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AgeTable, AgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAgeWith applies the HasEdge predicate on the "age" edge with a given conditions (other predicates).
func HasAgeWith(preds ...predicate.Age) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AgeTable, AgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosition applies the HasEdge predicate on the "position" edge.
func HasPosition() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPositionWith applies the HasEdge predicate on the "position" edge with a given conditions (other predicates).
func HasPositionWith(preds ...predicate.Position) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PositionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PositionTable, PositionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSalary applies the HasEdge predicate on the "salary" edge.
func HasSalary() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SalaryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SalaryTable, SalaryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSalaryWith applies the HasEdge predicate on the "salary" edge with a given conditions (other predicates).
func HasSalaryWith(preds ...predicate.Salary) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SalaryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SalaryTable, SalaryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookcourse applies the HasEdge predicate on the "bookcourse" edge.
func HasBookcourse() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookcourseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookcourseTable, BookcourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookcourseWith applies the HasEdge predicate on the "bookcourse" edge with a given conditions (other predicates).
func HasBookcourseWith(preds ...predicate.Bookcourse) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func HasEquipmentrental() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EquipmentrentalTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EquipmentrentalTable, EquipmentrentalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentrentalWith applies the HasEdge predicate on the "equipmentrental" edge with a given conditions (other predicates).
func HasEquipmentrentalWith(preds ...predicate.Equipmentrental) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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

// HasPromotion applies the HasEdge predicate on the "promotion" edge.
func HasPromotion() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PromotionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PromotionTable, PromotionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPromotionWith applies the HasEdge predicate on the "promotion" edge with a given conditions (other predicates).
func HasPromotionWith(preds ...predicate.Promotion) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		p(s.Not())
	})
}
