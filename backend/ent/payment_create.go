// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/member"
	"github.com/G16/app/ent/payment"
	"github.com/G16/app/ent/paymenttype"
	"github.com/G16/app/ent/promotion"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetPAYMENTAMOUNT sets the PAYMENTAMOUNT field.
func (pc *PaymentCreate) SetPAYMENTAMOUNT(s string) *PaymentCreate {
	pc.mutation.SetPAYMENTAMOUNT(s)
	return pc
}

// SetPHONENUMBER sets the PHONENUMBER field.
func (pc *PaymentCreate) SetPHONENUMBER(s string) *PaymentCreate {
	pc.mutation.SetPHONENUMBER(s)
	return pc
}

// SetEMAIL sets the EMAIL field.
func (pc *PaymentCreate) SetEMAIL(s string) *PaymentCreate {
	pc.mutation.SetEMAIL(s)
	return pc
}

// SetPAYMENTDATE sets the PAYMENTDATE field.
func (pc *PaymentCreate) SetPAYMENTDATE(t time.Time) *PaymentCreate {
	pc.mutation.SetPAYMENTDATE(t)
	return pc
}

// SetMemberID sets the member edge to Member by id.
func (pc *PaymentCreate) SetMemberID(id int) *PaymentCreate {
	pc.mutation.SetMemberID(id)
	return pc
}

// SetNillableMemberID sets the member edge to Member by id if the given value is not nil.
func (pc *PaymentCreate) SetNillableMemberID(id *int) *PaymentCreate {
	if id != nil {
		pc = pc.SetMemberID(*id)
	}
	return pc
}

// SetMember sets the member edge to Member.
func (pc *PaymentCreate) SetMember(m *Member) *PaymentCreate {
	return pc.SetMemberID(m.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (pc *PaymentCreate) SetEmployeeID(id int) *PaymentCreate {
	pc.mutation.SetEmployeeID(id)
	return pc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (pc *PaymentCreate) SetNillableEmployeeID(id *int) *PaymentCreate {
	if id != nil {
		pc = pc.SetEmployeeID(*id)
	}
	return pc
}

// SetEmployee sets the employee edge to Employee.
func (pc *PaymentCreate) SetEmployee(e *Employee) *PaymentCreate {
	return pc.SetEmployeeID(e.ID)
}

// SetPaymenttypeID sets the paymenttype edge to Paymenttype by id.
func (pc *PaymentCreate) SetPaymenttypeID(id int) *PaymentCreate {
	pc.mutation.SetPaymenttypeID(id)
	return pc
}

// SetNillablePaymenttypeID sets the paymenttype edge to Paymenttype by id if the given value is not nil.
func (pc *PaymentCreate) SetNillablePaymenttypeID(id *int) *PaymentCreate {
	if id != nil {
		pc = pc.SetPaymenttypeID(*id)
	}
	return pc
}

// SetPaymenttype sets the paymenttype edge to Paymenttype.
func (pc *PaymentCreate) SetPaymenttype(p *Paymenttype) *PaymentCreate {
	return pc.SetPaymenttypeID(p.ID)
}

// SetPromotionID sets the promotion edge to Promotion by id.
func (pc *PaymentCreate) SetPromotionID(id int) *PaymentCreate {
	pc.mutation.SetPromotionID(id)
	return pc
}

// SetNillablePromotionID sets the promotion edge to Promotion by id if the given value is not nil.
func (pc *PaymentCreate) SetNillablePromotionID(id *int) *PaymentCreate {
	if id != nil {
		pc = pc.SetPromotionID(*id)
	}
	return pc
}

// SetPromotion sets the promotion edge to Promotion.
func (pc *PaymentCreate) SetPromotion(p *Promotion) *PaymentCreate {
	return pc.SetPromotionID(p.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	if _, ok := pc.mutation.PAYMENTAMOUNT(); !ok {
		return nil, &ValidationError{Name: "PAYMENTAMOUNT", err: errors.New("ent: missing required field \"PAYMENTAMOUNT\"")}
	}
	if v, ok := pc.mutation.PAYMENTAMOUNT(); ok {
		if err := payment.PAYMENTAMOUNTValidator(v); err != nil {
			return nil, &ValidationError{Name: "PAYMENTAMOUNT", err: fmt.Errorf("ent: validator failed for field \"PAYMENTAMOUNT\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PHONENUMBER(); !ok {
		return nil, &ValidationError{Name: "PHONENUMBER", err: errors.New("ent: missing required field \"PHONENUMBER\"")}
	}
	if v, ok := pc.mutation.PHONENUMBER(); ok {
		if err := payment.PHONENUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "PHONENUMBER", err: fmt.Errorf("ent: validator failed for field \"PHONENUMBER\": %w", err)}
		}
	}
	if _, ok := pc.mutation.EMAIL(); !ok {
		return nil, &ValidationError{Name: "EMAIL", err: errors.New("ent: missing required field \"EMAIL\"")}
	}
	if v, ok := pc.mutation.EMAIL(); ok {
		if err := payment.EMAILValidator(v); err != nil {
			return nil, &ValidationError{Name: "EMAIL", err: fmt.Errorf("ent: validator failed for field \"EMAIL\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PAYMENTDATE(); !ok {
		return nil, &ValidationError{Name: "PAYMENTDATE", err: errors.New("ent: missing required field \"PAYMENTDATE\"")}
	}
	var (
		err  error
		node *Payment
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
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
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
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

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		pa    = &Payment{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: payment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payment.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PAYMENTAMOUNT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldPAYMENTAMOUNT,
		})
		pa.PAYMENTAMOUNT = value
	}
	if value, ok := pc.mutation.PHONENUMBER(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldPHONENUMBER,
		})
		pa.PHONENUMBER = value
	}
	if value, ok := pc.mutation.EMAIL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payment.FieldEMAIL,
		})
		pa.EMAIL = value
	}
	if value, ok := pc.mutation.PAYMENTDATE(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payment.FieldPAYMENTDATE,
		})
		pa.PAYMENTDATE = value
	}
	if nodes := pc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.MemberTable,
			Columns: []string{payment.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.EmployeeTable,
			Columns: []string{payment.EmployeeColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PaymenttypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.PaymenttypeTable,
			Columns: []string{payment.PaymenttypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.PromotionTable,
			Columns: []string{payment.PromotionColumn},
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
	return pa, _spec
}
