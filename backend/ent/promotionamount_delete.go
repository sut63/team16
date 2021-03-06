// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/G16/app/ent/predicate"
	"github.com/G16/app/ent/promotionamount"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PromotionamountDelete is the builder for deleting a Promotionamount entity.
type PromotionamountDelete struct {
	config
	hooks      []Hook
	mutation   *PromotionamountMutation
	predicates []predicate.Promotionamount
}

// Where adds a new predicate to the delete builder.
func (pd *PromotionamountDelete) Where(ps ...predicate.Promotionamount) *PromotionamountDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PromotionamountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionamountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PromotionamountDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PromotionamountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: promotionamount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotionamount.FieldID,
			},
		},
	}
	if ps := pd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PromotionamountDeleteOne is the builder for deleting a single Promotionamount entity.
type PromotionamountDeleteOne struct {
	pd *PromotionamountDelete
}

// Exec executes the deletion query.
func (pdo *PromotionamountDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{promotionamount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PromotionamountDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
