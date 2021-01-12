// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/G16/app/ent/payment"
	"github.com/G16/app/ent/paymenttype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PaymenttypeCreate is the builder for creating a Paymenttype entity.
type PaymenttypeCreate struct {
	config
	mutation *PaymenttypeMutation
	hooks    []Hook
}

// SetTYPE sets the TYPE field.
func (pc *PaymenttypeCreate) SetTYPE(s string) *PaymenttypeCreate {
	pc.mutation.SetTYPE(s)
	return pc
}

// AddPaymentIDs adds the payment edge to Payment by ids.
func (pc *PaymenttypeCreate) AddPaymentIDs(ids ...int) *PaymenttypeCreate {
	pc.mutation.AddPaymentIDs(ids...)
	return pc
}

// AddPayment adds the payment edges to Payment.
func (pc *PaymenttypeCreate) AddPayment(p ...*Payment) *PaymenttypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPaymentIDs(ids...)
}

// Mutation returns the PaymenttypeMutation object of the builder.
func (pc *PaymenttypeCreate) Mutation() *PaymenttypeMutation {
	return pc.mutation
}

// Save creates the Paymenttype in the database.
func (pc *PaymenttypeCreate) Save(ctx context.Context) (*Paymenttype, error) {
	if _, ok := pc.mutation.TYPE(); !ok {
		return nil, &ValidationError{Name: "TYPE", err: errors.New("ent: missing required field \"TYPE\"")}
	}
	var (
		err  error
		node *Paymenttype
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymenttypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymenttypeCreate) SaveX(ctx context.Context) *Paymenttype {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PaymenttypeCreate) sqlSave(ctx context.Context) (*Paymenttype, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PaymenttypeCreate) createSpec() (*Paymenttype, *sqlgraph.CreateSpec) {
	var (
		pa    = &Paymenttype{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paymenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymenttype.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.TYPE(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paymenttype.FieldTYPE,
		})
		pa.TYPE = value
	}
	if nodes := pc.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymenttype.PaymentTable,
			Columns: []string{paymenttype.PaymentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}