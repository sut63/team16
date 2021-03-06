// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/promotionamount"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PromotionamountCreate is the builder for creating a Promotionamount entity.
type PromotionamountCreate struct {
	config
	mutation *PromotionamountMutation
	hooks    []Hook
}

// SetAMOUNT sets the AMOUNT field.
func (pc *PromotionamountCreate) SetAMOUNT(i int) *PromotionamountCreate {
	pc.mutation.SetAMOUNT(i)
	return pc
}

// AddPromotionIDs adds the promotion edge to Promotion by ids.
func (pc *PromotionamountCreate) AddPromotionIDs(ids ...int) *PromotionamountCreate {
	pc.mutation.AddPromotionIDs(ids...)
	return pc
}

// AddPromotion adds the promotion edges to Promotion.
func (pc *PromotionamountCreate) AddPromotion(p ...*Promotion) *PromotionamountCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPromotionIDs(ids...)
}

// Mutation returns the PromotionamountMutation object of the builder.
func (pc *PromotionamountCreate) Mutation() *PromotionamountMutation {
	return pc.mutation
}

// Save creates the Promotionamount in the database.
func (pc *PromotionamountCreate) Save(ctx context.Context) (*Promotionamount, error) {
	if _, ok := pc.mutation.AMOUNT(); !ok {
		return nil, &ValidationError{Name: "AMOUNT", err: errors.New("ent: missing required field \"AMOUNT\"")}
	}
	if v, ok := pc.mutation.AMOUNT(); ok {
		if err := promotionamount.AMOUNTValidator(v); err != nil {
			return nil, &ValidationError{Name: "AMOUNT", err: fmt.Errorf("ent: validator failed for field \"AMOUNT\": %w", err)}
		}
	}
	var (
		err  error
		node *Promotionamount
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionamountMutation)
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
func (pc *PromotionamountCreate) SaveX(ctx context.Context) *Promotionamount {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PromotionamountCreate) sqlSave(ctx context.Context) (*Promotionamount, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PromotionamountCreate) createSpec() (*Promotionamount, *sqlgraph.CreateSpec) {
	var (
		pr    = &Promotionamount{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: promotionamount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotionamount.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.AMOUNT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: promotionamount.FieldAMOUNT,
		})
		pr.AMOUNT = value
	}
	if nodes := pc.mutation.PromotionIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}
