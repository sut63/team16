// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/G16/app/ent/payment"
	"github.com/G16/app/ent/predicate"
	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/promotionamount"
	"github.com/G16/app/ent/promotiontype"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PromotionUpdate is the builder for updating Promotion entities.
type PromotionUpdate struct {
	config
	hooks      []Hook
	mutation   *PromotionMutation
	predicates []predicate.Promotion
}

// Where adds a new predicate for the builder.
func (pu *PromotionUpdate) Where(ps ...predicate.Promotion) *PromotionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetNAME sets the NAME field.
func (pu *PromotionUpdate) SetNAME(s string) *PromotionUpdate {
	pu.mutation.SetNAME(s)
	return pu
}

// SetDESC sets the DESC field.
func (pu *PromotionUpdate) SetDESC(s string) *PromotionUpdate {
	pu.mutation.SetDESC(s)
	return pu
}

// SetCODE sets the CODE field.
func (pu *PromotionUpdate) SetCODE(s string) *PromotionUpdate {
	pu.mutation.SetCODE(s)
	return pu
}

// SetDATE sets the DATE field.
func (pu *PromotionUpdate) SetDATE(t time.Time) *PromotionUpdate {
	pu.mutation.SetDATE(t)
	return pu
}

// SetPromotiontypeID sets the promotiontype edge to Promotiontype by id.
func (pu *PromotionUpdate) SetPromotiontypeID(id int) *PromotionUpdate {
	pu.mutation.SetPromotiontypeID(id)
	return pu
}

// SetNillablePromotiontypeID sets the promotiontype edge to Promotiontype by id if the given value is not nil.
func (pu *PromotionUpdate) SetNillablePromotiontypeID(id *int) *PromotionUpdate {
	if id != nil {
		pu = pu.SetPromotiontypeID(*id)
	}
	return pu
}

// SetPromotiontype sets the promotiontype edge to Promotiontype.
func (pu *PromotionUpdate) SetPromotiontype(p *Promotiontype) *PromotionUpdate {
	return pu.SetPromotiontypeID(p.ID)
}

// SetPromotionamountID sets the promotionamount edge to Promotionamount by id.
func (pu *PromotionUpdate) SetPromotionamountID(id int) *PromotionUpdate {
	pu.mutation.SetPromotionamountID(id)
	return pu
}

// SetNillablePromotionamountID sets the promotionamount edge to Promotionamount by id if the given value is not nil.
func (pu *PromotionUpdate) SetNillablePromotionamountID(id *int) *PromotionUpdate {
	if id != nil {
		pu = pu.SetPromotionamountID(*id)
	}
	return pu
}

// SetPromotionamount sets the promotionamount edge to Promotionamount.
func (pu *PromotionUpdate) SetPromotionamount(p *Promotionamount) *PromotionUpdate {
	return pu.SetPromotionamountID(p.ID)
}

// AddPaymentIDs adds the payment edge to Payment by ids.
func (pu *PromotionUpdate) AddPaymentIDs(ids ...int) *PromotionUpdate {
	pu.mutation.AddPaymentIDs(ids...)
	return pu
}

// AddPayment adds the payment edges to Payment.
func (pu *PromotionUpdate) AddPayment(p ...*Payment) *PromotionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPaymentIDs(ids...)
}

// Mutation returns the PromotionMutation object of the builder.
func (pu *PromotionUpdate) Mutation() *PromotionMutation {
	return pu.mutation
}

// ClearPromotiontype clears the promotiontype edge to Promotiontype.
func (pu *PromotionUpdate) ClearPromotiontype() *PromotionUpdate {
	pu.mutation.ClearPromotiontype()
	return pu
}

// ClearPromotionamount clears the promotionamount edge to Promotionamount.
func (pu *PromotionUpdate) ClearPromotionamount() *PromotionUpdate {
	pu.mutation.ClearPromotionamount()
	return pu
}

// RemovePaymentIDs removes the payment edge to Payment by ids.
func (pu *PromotionUpdate) RemovePaymentIDs(ids ...int) *PromotionUpdate {
	pu.mutation.RemovePaymentIDs(ids...)
	return pu
}

// RemovePayment removes payment edges to Payment.
func (pu *PromotionUpdate) RemovePayment(p ...*Payment) *PromotionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePaymentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PromotionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.NAME(); ok {
		if err := promotion.NAMEValidator(v); err != nil {
			return 0, &ValidationError{Name: "NAME", err: fmt.Errorf("ent: validator failed for field \"NAME\": %w", err)}
		}
	}
	if v, ok := pu.mutation.DESC(); ok {
		if err := promotion.DESCValidator(v); err != nil {
			return 0, &ValidationError{Name: "DESC", err: fmt.Errorf("ent: validator failed for field \"DESC\": %w", err)}
		}
	}
	if v, ok := pu.mutation.CODE(); ok {
		if err := promotion.CODEValidator(v); err != nil {
			return 0, &ValidationError{Name: "CODE", err: fmt.Errorf("ent: validator failed for field \"CODE\": %w", err)}
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
			mutation, ok := m.(*PromotionMutation)
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
func (pu *PromotionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PromotionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PromotionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PromotionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotion.Table,
			Columns: promotion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotion.FieldID,
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
	if value, ok := pu.mutation.NAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldNAME,
		})
	}
	if value, ok := pu.mutation.DESC(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldDESC,
		})
	}
	if value, ok := pu.mutation.CODE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldCODE,
		})
	}
	if value, ok := pu.mutation.DATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldDATE,
		})
	}
	if pu.mutation.PromotiontypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotiontypeTable,
			Columns: []string{promotion.PromotiontypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotiontype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PromotiontypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotiontypeTable,
			Columns: []string{promotion.PromotiontypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PromotionamountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotionamountTable,
			Columns: []string{promotion.PromotionamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotionamount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PromotionamountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotionamountTable,
			Columns: []string{promotion.PromotionamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotionamount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotion.PaymentTable,
			Columns: []string{promotion.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotion.PaymentTable,
			Columns: []string{promotion.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
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
			err = &NotFoundError{promotion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PromotionUpdateOne is the builder for updating a single Promotion entity.
type PromotionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PromotionMutation
}

// SetNAME sets the NAME field.
func (puo *PromotionUpdateOne) SetNAME(s string) *PromotionUpdateOne {
	puo.mutation.SetNAME(s)
	return puo
}

// SetDESC sets the DESC field.
func (puo *PromotionUpdateOne) SetDESC(s string) *PromotionUpdateOne {
	puo.mutation.SetDESC(s)
	return puo
}

// SetCODE sets the CODE field.
func (puo *PromotionUpdateOne) SetCODE(s string) *PromotionUpdateOne {
	puo.mutation.SetCODE(s)
	return puo
}

// SetDATE sets the DATE field.
func (puo *PromotionUpdateOne) SetDATE(t time.Time) *PromotionUpdateOne {
	puo.mutation.SetDATE(t)
	return puo
}

// SetPromotiontypeID sets the promotiontype edge to Promotiontype by id.
func (puo *PromotionUpdateOne) SetPromotiontypeID(id int) *PromotionUpdateOne {
	puo.mutation.SetPromotiontypeID(id)
	return puo
}

// SetNillablePromotiontypeID sets the promotiontype edge to Promotiontype by id if the given value is not nil.
func (puo *PromotionUpdateOne) SetNillablePromotiontypeID(id *int) *PromotionUpdateOne {
	if id != nil {
		puo = puo.SetPromotiontypeID(*id)
	}
	return puo
}

// SetPromotiontype sets the promotiontype edge to Promotiontype.
func (puo *PromotionUpdateOne) SetPromotiontype(p *Promotiontype) *PromotionUpdateOne {
	return puo.SetPromotiontypeID(p.ID)
}

// SetPromotionamountID sets the promotionamount edge to Promotionamount by id.
func (puo *PromotionUpdateOne) SetPromotionamountID(id int) *PromotionUpdateOne {
	puo.mutation.SetPromotionamountID(id)
	return puo
}

// SetNillablePromotionamountID sets the promotionamount edge to Promotionamount by id if the given value is not nil.
func (puo *PromotionUpdateOne) SetNillablePromotionamountID(id *int) *PromotionUpdateOne {
	if id != nil {
		puo = puo.SetPromotionamountID(*id)
	}
	return puo
}

// SetPromotionamount sets the promotionamount edge to Promotionamount.
func (puo *PromotionUpdateOne) SetPromotionamount(p *Promotionamount) *PromotionUpdateOne {
	return puo.SetPromotionamountID(p.ID)
}

// AddPaymentIDs adds the payment edge to Payment by ids.
func (puo *PromotionUpdateOne) AddPaymentIDs(ids ...int) *PromotionUpdateOne {
	puo.mutation.AddPaymentIDs(ids...)
	return puo
}

// AddPayment adds the payment edges to Payment.
func (puo *PromotionUpdateOne) AddPayment(p ...*Payment) *PromotionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPaymentIDs(ids...)
}

// Mutation returns the PromotionMutation object of the builder.
func (puo *PromotionUpdateOne) Mutation() *PromotionMutation {
	return puo.mutation
}

// ClearPromotiontype clears the promotiontype edge to Promotiontype.
func (puo *PromotionUpdateOne) ClearPromotiontype() *PromotionUpdateOne {
	puo.mutation.ClearPromotiontype()
	return puo
}

// ClearPromotionamount clears the promotionamount edge to Promotionamount.
func (puo *PromotionUpdateOne) ClearPromotionamount() *PromotionUpdateOne {
	puo.mutation.ClearPromotionamount()
	return puo
}

// RemovePaymentIDs removes the payment edge to Payment by ids.
func (puo *PromotionUpdateOne) RemovePaymentIDs(ids ...int) *PromotionUpdateOne {
	puo.mutation.RemovePaymentIDs(ids...)
	return puo
}

// RemovePayment removes payment edges to Payment.
func (puo *PromotionUpdateOne) RemovePayment(p ...*Payment) *PromotionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePaymentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PromotionUpdateOne) Save(ctx context.Context) (*Promotion, error) {
	if v, ok := puo.mutation.NAME(); ok {
		if err := promotion.NAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "NAME", err: fmt.Errorf("ent: validator failed for field \"NAME\": %w", err)}
		}
	}
	if v, ok := puo.mutation.DESC(); ok {
		if err := promotion.DESCValidator(v); err != nil {
			return nil, &ValidationError{Name: "DESC", err: fmt.Errorf("ent: validator failed for field \"DESC\": %w", err)}
		}
	}
	if v, ok := puo.mutation.CODE(); ok {
		if err := promotion.CODEValidator(v); err != nil {
			return nil, &ValidationError{Name: "CODE", err: fmt.Errorf("ent: validator failed for field \"CODE\": %w", err)}
		}
	}

	var (
		err  error
		node *Promotion
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionMutation)
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
func (puo *PromotionUpdateOne) SaveX(ctx context.Context) *Promotion {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PromotionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PromotionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PromotionUpdateOne) sqlSave(ctx context.Context) (pr *Promotion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotion.Table,
			Columns: promotion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotion.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Promotion.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.NAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldNAME,
		})
	}
	if value, ok := puo.mutation.DESC(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldDESC,
		})
	}
	if value, ok := puo.mutation.CODE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldCODE,
		})
	}
	if value, ok := puo.mutation.DATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldDATE,
		})
	}
	if puo.mutation.PromotiontypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotiontypeTable,
			Columns: []string{promotion.PromotiontypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotiontype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PromotiontypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotiontypeTable,
			Columns: []string{promotion.PromotiontypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PromotionamountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotionamountTable,
			Columns: []string{promotion.PromotionamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotionamount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PromotionamountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.PromotionamountTable,
			Columns: []string{promotion.PromotionamountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotionamount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotion.PaymentTable,
			Columns: []string{promotion.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   promotion.PaymentTable,
			Columns: []string{promotion.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Promotion{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
