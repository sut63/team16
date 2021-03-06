// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/predicate"
	"github.com/G16/app/ent/salary"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SalaryUpdate is the builder for updating Salary entities.
type SalaryUpdate struct {
	config
	hooks      []Hook
	mutation   *SalaryMutation
	predicates []predicate.Salary
}

// Where adds a new predicate for the builder.
func (su *SalaryUpdate) Where(ps ...predicate.Salary) *SalaryUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetSALARY sets the SALARY field.
func (su *SalaryUpdate) SetSALARY(i int) *SalaryUpdate {
	su.mutation.ResetSALARY()
	su.mutation.SetSALARY(i)
	return su
}

// AddSALARY adds i to SALARY.
func (su *SalaryUpdate) AddSALARY(i int) *SalaryUpdate {
	su.mutation.AddSALARY(i)
	return su
}

// AddEmployeeIDs adds the employee edge to Employee by ids.
func (su *SalaryUpdate) AddEmployeeIDs(ids ...int) *SalaryUpdate {
	su.mutation.AddEmployeeIDs(ids...)
	return su
}

// AddEmployee adds the employee edges to Employee.
func (su *SalaryUpdate) AddEmployee(e ...*Employee) *SalaryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEmployeeIDs(ids...)
}

// Mutation returns the SalaryMutation object of the builder.
func (su *SalaryUpdate) Mutation() *SalaryMutation {
	return su.mutation
}

// RemoveEmployeeIDs removes the employee edge to Employee by ids.
func (su *SalaryUpdate) RemoveEmployeeIDs(ids ...int) *SalaryUpdate {
	su.mutation.RemoveEmployeeIDs(ids...)
	return su
}

// RemoveEmployee removes employee edges to Employee.
func (su *SalaryUpdate) RemoveEmployee(e ...*Employee) *SalaryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SalaryUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SalaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SalaryUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SalaryUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SalaryUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SalaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   salary.Table,
			Columns: salary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: salary.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SALARY(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: salary.FieldSALARY,
		})
	}
	if value, ok := su.mutation.AddedSALARY(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: salary.FieldSALARY,
		})
	}
	if nodes := su.mutation.RemovedEmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{salary.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SalaryUpdateOne is the builder for updating a single Salary entity.
type SalaryUpdateOne struct {
	config
	hooks    []Hook
	mutation *SalaryMutation
}

// SetSALARY sets the SALARY field.
func (suo *SalaryUpdateOne) SetSALARY(i int) *SalaryUpdateOne {
	suo.mutation.ResetSALARY()
	suo.mutation.SetSALARY(i)
	return suo
}

// AddSALARY adds i to SALARY.
func (suo *SalaryUpdateOne) AddSALARY(i int) *SalaryUpdateOne {
	suo.mutation.AddSALARY(i)
	return suo
}

// AddEmployeeIDs adds the employee edge to Employee by ids.
func (suo *SalaryUpdateOne) AddEmployeeIDs(ids ...int) *SalaryUpdateOne {
	suo.mutation.AddEmployeeIDs(ids...)
	return suo
}

// AddEmployee adds the employee edges to Employee.
func (suo *SalaryUpdateOne) AddEmployee(e ...*Employee) *SalaryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEmployeeIDs(ids...)
}

// Mutation returns the SalaryMutation object of the builder.
func (suo *SalaryUpdateOne) Mutation() *SalaryMutation {
	return suo.mutation
}

// RemoveEmployeeIDs removes the employee edge to Employee by ids.
func (suo *SalaryUpdateOne) RemoveEmployeeIDs(ids ...int) *SalaryUpdateOne {
	suo.mutation.RemoveEmployeeIDs(ids...)
	return suo
}

// RemoveEmployee removes employee edges to Employee.
func (suo *SalaryUpdateOne) RemoveEmployee(e ...*Employee) *SalaryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SalaryUpdateOne) Save(ctx context.Context) (*Salary, error) {

	var (
		err  error
		node *Salary
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SalaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SalaryUpdateOne) SaveX(ctx context.Context) *Salary {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SalaryUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SalaryUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SalaryUpdateOne) sqlSave(ctx context.Context) (s *Salary, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   salary.Table,
			Columns: salary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: salary.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Salary.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.SALARY(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: salary.FieldSALARY,
		})
	}
	if value, ok := suo.mutation.AddedSALARY(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: salary.FieldSALARY,
		})
	}
	if nodes := suo.mutation.RemovedEmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Salary{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{salary.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
