// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/G16/app/ent/age"
	"github.com/G16/app/ent/bookcourse"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/equipmentrental"
	"github.com/G16/app/ent/payment"
	"github.com/G16/app/ent/position"
	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/salary"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
}

// SetEMPLOYEEID sets the EMPLOYEEID field.
func (ec *EmployeeCreate) SetEMPLOYEEID(s string) *EmployeeCreate {
	ec.mutation.SetEMPLOYEEID(s)
	return ec
}

// SetEMPLOYEENAME sets the EMPLOYEENAME field.
func (ec *EmployeeCreate) SetEMPLOYEENAME(s string) *EmployeeCreate {
	ec.mutation.SetEMPLOYEENAME(s)
	return ec
}

// SetEMPLOYEEADDRESS sets the EMPLOYEEADDRESS field.
func (ec *EmployeeCreate) SetEMPLOYEEADDRESS(s string) *EmployeeCreate {
	ec.mutation.SetEMPLOYEEADDRESS(s)
	return ec
}

// SetIDCARDNUMBER sets the IDCARDNUMBER field.
func (ec *EmployeeCreate) SetIDCARDNUMBER(s string) *EmployeeCreate {
	ec.mutation.SetIDCARDNUMBER(s)
	return ec
}

// SetAgeID sets the age edge to Age by id.
func (ec *EmployeeCreate) SetAgeID(id int) *EmployeeCreate {
	ec.mutation.SetAgeID(id)
	return ec
}

// SetNillableAgeID sets the age edge to Age by id if the given value is not nil.
func (ec *EmployeeCreate) SetNillableAgeID(id *int) *EmployeeCreate {
	if id != nil {
		ec = ec.SetAgeID(*id)
	}
	return ec
}

// SetAge sets the age edge to Age.
func (ec *EmployeeCreate) SetAge(a *Age) *EmployeeCreate {
	return ec.SetAgeID(a.ID)
}

// SetPositionID sets the position edge to Position by id.
func (ec *EmployeeCreate) SetPositionID(id int) *EmployeeCreate {
	ec.mutation.SetPositionID(id)
	return ec
}

// SetNillablePositionID sets the position edge to Position by id if the given value is not nil.
func (ec *EmployeeCreate) SetNillablePositionID(id *int) *EmployeeCreate {
	if id != nil {
		ec = ec.SetPositionID(*id)
	}
	return ec
}

// SetPosition sets the position edge to Position.
func (ec *EmployeeCreate) SetPosition(p *Position) *EmployeeCreate {
	return ec.SetPositionID(p.ID)
}

// SetSalaryID sets the salary edge to Salary by id.
func (ec *EmployeeCreate) SetSalaryID(id int) *EmployeeCreate {
	ec.mutation.SetSalaryID(id)
	return ec
}

// SetNillableSalaryID sets the salary edge to Salary by id if the given value is not nil.
func (ec *EmployeeCreate) SetNillableSalaryID(id *int) *EmployeeCreate {
	if id != nil {
		ec = ec.SetSalaryID(*id)
	}
	return ec
}

// SetSalary sets the salary edge to Salary.
func (ec *EmployeeCreate) SetSalary(s *Salary) *EmployeeCreate {
	return ec.SetSalaryID(s.ID)
}

// AddPaymentIDs adds the payment edge to Payment by ids.
func (ec *EmployeeCreate) AddPaymentIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddPaymentIDs(ids...)
	return ec
}

// AddPayment adds the payment edges to Payment.
func (ec *EmployeeCreate) AddPayment(p ...*Payment) *EmployeeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddPaymentIDs(ids...)
}

// AddEquipmentIDs adds the equipment edge to Equipment by ids.
func (ec *EmployeeCreate) AddEquipmentIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddEquipmentIDs(ids...)
	return ec
}

// AddEquipment adds the equipment edges to Equipment.
func (ec *EmployeeCreate) AddEquipment(e ...*Equipment) *EmployeeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentIDs(ids...)
}

// AddBookcourseIDs adds the bookcourse edge to Bookcourse by ids.
func (ec *EmployeeCreate) AddBookcourseIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddBookcourseIDs(ids...)
	return ec
}

// AddBookcourse adds the bookcourse edges to Bookcourse.
func (ec *EmployeeCreate) AddBookcourse(b ...*Bookcourse) *EmployeeCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ec.AddBookcourseIDs(ids...)
}

// AddEquipmentrentalIDs adds the equipmentrental edge to Equipmentrental by ids.
func (ec *EmployeeCreate) AddEquipmentrentalIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddEquipmentrentalIDs(ids...)
	return ec
}

// AddEquipmentrental adds the equipmentrental edges to Equipmentrental.
func (ec *EmployeeCreate) AddEquipmentrental(e ...*Equipmentrental) *EmployeeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentrentalIDs(ids...)
}

// AddPromotionIDs adds the promotion edge to Promotion by ids.
func (ec *EmployeeCreate) AddPromotionIDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddPromotionIDs(ids...)
	return ec
}

// AddPromotion adds the promotion edges to Promotion.
func (ec *EmployeeCreate) AddPromotion(p ...*Promotion) *EmployeeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddPromotionIDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	if _, ok := ec.mutation.EMPLOYEEID(); !ok {
		return nil, &ValidationError{Name: "EMPLOYEEID", err: errors.New("ent: missing required field \"EMPLOYEEID\"")}
	}
	if v, ok := ec.mutation.EMPLOYEEID(); ok {
		if err := employee.EMPLOYEEIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "EMPLOYEEID", err: fmt.Errorf("ent: validator failed for field \"EMPLOYEEID\": %w", err)}
		}
	}
	if _, ok := ec.mutation.EMPLOYEENAME(); !ok {
		return nil, &ValidationError{Name: "EMPLOYEENAME", err: errors.New("ent: missing required field \"EMPLOYEENAME\"")}
	}
	if v, ok := ec.mutation.EMPLOYEENAME(); ok {
		if err := employee.EMPLOYEENAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "EMPLOYEENAME", err: fmt.Errorf("ent: validator failed for field \"EMPLOYEENAME\": %w", err)}
		}
	}
	if _, ok := ec.mutation.EMPLOYEEADDRESS(); !ok {
		return nil, &ValidationError{Name: "EMPLOYEEADDRESS", err: errors.New("ent: missing required field \"EMPLOYEEADDRESS\"")}
	}
	if v, ok := ec.mutation.EMPLOYEEADDRESS(); ok {
		if err := employee.EMPLOYEEADDRESSValidator(v); err != nil {
			return nil, &ValidationError{Name: "EMPLOYEEADDRESS", err: fmt.Errorf("ent: validator failed for field \"EMPLOYEEADDRESS\": %w", err)}
		}
	}
	if _, ok := ec.mutation.IDCARDNUMBER(); !ok {
		return nil, &ValidationError{Name: "IDCARDNUMBER", err: errors.New("ent: missing required field \"IDCARDNUMBER\"")}
	}
	if v, ok := ec.mutation.IDCARDNUMBER(); ok {
		if err := employee.IDCARDNUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "IDCARDNUMBER", err: fmt.Errorf("ent: validator failed for field \"IDCARDNUMBER\": %w", err)}
		}
	}
	var (
		err  error
		node *Employee
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		e     = &Employee{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.EMPLOYEEID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEMPLOYEEID,
		})
		e.EMPLOYEEID = value
	}
	if value, ok := ec.mutation.EMPLOYEENAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEMPLOYEENAME,
		})
		e.EMPLOYEENAME = value
	}
	if value, ok := ec.mutation.EMPLOYEEADDRESS(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEMPLOYEEADDRESS,
		})
		e.EMPLOYEEADDRESS = value
	}
	if value, ok := ec.mutation.IDCARDNUMBER(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldIDCARDNUMBER,
		})
		e.IDCARDNUMBER = value
	}
	if nodes := ec.mutation.AgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.AgeTable,
			Columns: []string{employee.AgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: age.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.PositionTable,
			Columns: []string{employee.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.SalaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   employee.SalaryTable,
			Columns: []string{employee.SalaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PaymentTable,
			Columns: []string{employee.PaymentColumn},
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
	if nodes := ec.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EquipmentTable,
			Columns: []string{employee.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.BookcourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.BookcourseTable,
			Columns: []string{employee.BookcourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentrentalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.EquipmentrentalTable,
			Columns: []string{employee.EquipmentrentalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentrental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.PromotionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.PromotionTable,
			Columns: []string{employee.PromotionColumn},
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
	return e, _spec
}
