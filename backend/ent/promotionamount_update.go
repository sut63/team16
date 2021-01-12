// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/G16/app/ent/predicate"
	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/promotionamount"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PromotionamountUpdate is the builder for updating Promotionamount entities.
type PromotionamountUpdate struct {
	config
	hooks      []Hook
	mutation   *PromotionamountMutation
	predicates []predicate.Promotionamount
}

// Where adds a new predicate for the builder.
func (pu *PromotionamountUpdate) Where(ps ...predicate.Promotionamount) *PromotionamountUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAMOUNT sets the AMOUNT field.
func (pu *PromotionamountUpdate) SetAMOUNT(i int) *PromotionamountUpdate {
	pu.mutation.ResetAMOUNT()
	pu.mutation.SetAMOUNT(i)
	return pu
}

// AddAMOUNT adds i to AMOUNT.
func (pu *PromotionamountUpdate) AddAMOUNT(i int) *PromotionamountUpdate {
	pu.mutation.AddAMOUNT(i)
	return pu
}

// AddPromotionIDs adds the promotion edge to Promotion by ids.
func (pu *PromotionamountUpdate) AddPromotionIDs(ids ...int) *PromotionamountUpdate {
	pu.mutation.AddPromotionIDs(ids...)
	return pu
}

// AddPromotion adds the promotion edges to Promotion.
func (pu *PromotionamountUpdate) AddPromotion(p ...*Promotion) *PromotionamountUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPromotionIDs(ids...)
}

// Mutation returns the PromotionamountMutation object of the builder.
func (pu *PromotionamountUpdate) Mutation() *PromotionamountMutation {
	return pu.mutation
}

// RemovePromotionIDs removes the promotion edge to Promotion by ids.
func (pu *PromotionamountUpdate) RemovePromotionIDs(ids ...int) *PromotionamountUpdate {
	pu.mutation.RemovePromotionIDs(ids...)
	return pu
}

// RemovePromotion removes promotion edges to Promotion.
func (pu *PromotionamountUpdate) RemovePromotion(p ...*Promotion) *PromotionamountUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePromotionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PromotionamountUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.AMOUNT(); ok {
		if err := promotionamount.AMOUNTValidator(v); err != nil {
			return 0, &ValidationError{Name: "AMOUNT", err: fmt.Errorf("ent: validator failed for field \"AMOUNT\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionamountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PromotionamountUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PromotionamountUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PromotionamountUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PromotionamountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotionamount.Table,
			Columns: promotionamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotionamount.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.AMOUNT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: promotionamount.FieldAMOUNT,
		})
	}
	if value, ok := pu.mutation.AddedAMOUNT(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: promotionamount.FieldAMOUNT,
		})
	}
	if nodes := pu.mutation.RemovedPromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotionamount.PromotionTable,
			Columns: []string{promotionamount.PromotionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotionamount.PromotionTable,
			Columns: []string{promotionamount.PromotionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotionamount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PromotionamountUpdateOne is the builder for updating a single Promotionamount entity.
type PromotionamountUpdateOne struct {
	config
	hooks    []Hook
	mutation *PromotionamountMutation
}

// SetAMOUNT sets the AMOUNT field.
func (puo *PromotionamountUpdateOne) SetAMOUNT(i int) *PromotionamountUpdateOne {
	puo.mutation.ResetAMOUNT()
	puo.mutation.SetAMOUNT(i)
	return puo
}

// AddAMOUNT adds i to AMOUNT.
func (puo *PromotionamountUpdateOne) AddAMOUNT(i int) *PromotionamountUpdateOne {
	puo.mutation.AddAMOUNT(i)
	return puo
}

// AddPromotionIDs adds the promotion edge to Promotion by ids.
func (puo *PromotionamountUpdateOne) AddPromotionIDs(ids ...int) *PromotionamountUpdateOne {
	puo.mutation.AddPromotionIDs(ids...)
	return puo
}

// AddPromotion adds the promotion edges to Promotion.
func (puo *PromotionamountUpdateOne) AddPromotion(p ...*Promotion) *PromotionamountUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPromotionIDs(ids...)
}

// Mutation returns the PromotionamountMutation object of the builder.
func (puo *PromotionamountUpdateOne) Mutation() *PromotionamountMutation {
	return puo.mutation
}

// RemovePromotionIDs removes the promotion edge to Promotion by ids.
func (puo *PromotionamountUpdateOne) RemovePromotionIDs(ids ...int) *PromotionamountUpdateOne {
	puo.mutation.RemovePromotionIDs(ids...)
	return puo
}

// RemovePromotion removes promotion edges to Promotion.
func (puo *PromotionamountUpdateOne) RemovePromotion(p ...*Promotion) *PromotionamountUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePromotionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PromotionamountUpdateOne) Save(ctx context.Context) (*Promotionamount, error) {
	if v, ok := puo.mutation.AMOUNT(); ok {
		if err := promotionamount.AMOUNTValidator(v); err != nil {
			return nil, &ValidationError{Name: "AMOUNT", err: fmt.Errorf("ent: validator failed for field \"AMOUNT\": %w", err)}
		}
	}

	var (
		err  error
		node *Promotionamount
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionamountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PromotionamountUpdateOne) SaveX(ctx context.Context) *Promotionamount {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PromotionamountUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PromotionamountUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PromotionamountUpdateOne) sqlSave(ctx context.Context) (pr *Promotionamount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotionamount.Table,
			Columns: promotionamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotionamount.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Promotionamount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.AMOUNT(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: promotionamount.FieldAMOUNT,
		})
	}
	if value, ok := puo.mutation.AddedAMOUNT(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: promotionamount.FieldAMOUNT,
		})
	}
	if nodes := puo.mutation.RemovedPromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotionamount.PromotionTable,
			Columns: []string{promotionamount.PromotionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotionamount.PromotionTable,
			Columns: []string{promotionamount.PromotionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Promotionamount{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotionamount.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}