// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/G16/app/ent/classifier"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/equipmentrental"
	"github.com/G16/app/ent/equipmenttype"
	"github.com/G16/app/ent/zone"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// EquipmentCreate is the builder for creating a Equipment entity.
type EquipmentCreate struct {
	config
	mutation *EquipmentMutation
	hooks    []Hook
}

// SetEQUIPMENTNAME sets the EQUIPMENTNAME field.
func (ec *EquipmentCreate) SetEQUIPMENTNAME(s string) *EquipmentCreate {
	ec.mutation.SetEQUIPMENTNAME(s)
	return ec
}

// SetEQUIPMENTAMOUNT sets the EQUIPMENTAMOUNT field.
func (ec *EquipmentCreate) SetEQUIPMENTAMOUNT(i int) *EquipmentCreate {
	ec.mutation.SetEQUIPMENTAMOUNT(i)
	return ec
}

// SetEQUIPMENTDETAIL sets the EQUIPMENTDETAIL field.
func (ec *EquipmentCreate) SetEQUIPMENTDETAIL(s string) *EquipmentCreate {
	ec.mutation.SetEQUIPMENTDETAIL(s)
	return ec
}

// SetEQUIPMENTDATE sets the EQUIPMENTDATE field.
func (ec *EquipmentCreate) SetEQUIPMENTDATE(t time.Time) *EquipmentCreate {
	ec.mutation.SetEQUIPMENTDATE(t)
	return ec
}

// SetClassifierID sets the classifier edge to Classifier by id.
func (ec *EquipmentCreate) SetClassifierID(id int) *EquipmentCreate {
	ec.mutation.SetClassifierID(id)
	return ec
}

// SetNillableClassifierID sets the classifier edge to Classifier by id if the given value is not nil.
func (ec *EquipmentCreate) SetNillableClassifierID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetClassifierID(*id)
	}
	return ec
}

// SetClassifier sets the classifier edge to Classifier.
func (ec *EquipmentCreate) SetClassifier(c *Classifier) *EquipmentCreate {
	return ec.SetClassifierID(c.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (ec *EquipmentCreate) SetEmployeeID(id int) *EquipmentCreate {
	ec.mutation.SetEmployeeID(id)
	return ec
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEmployeeID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetEmployeeID(*id)
	}
	return ec
}

// SetEmployee sets the employee edge to Employee.
func (ec *EquipmentCreate) SetEmployee(e *Employee) *EquipmentCreate {
	return ec.SetEmployeeID(e.ID)
}

// SetEquipmenttypeID sets the equipmenttype edge to Equipmenttype by id.
func (ec *EquipmentCreate) SetEquipmenttypeID(id int) *EquipmentCreate {
	ec.mutation.SetEquipmenttypeID(id)
	return ec
}

// SetNillableEquipmenttypeID sets the equipmenttype edge to Equipmenttype by id if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEquipmenttypeID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetEquipmenttypeID(*id)
	}
	return ec
}

// SetEquipmenttype sets the equipmenttype edge to Equipmenttype.
func (ec *EquipmentCreate) SetEquipmenttype(e *Equipmenttype) *EquipmentCreate {
	return ec.SetEquipmenttypeID(e.ID)
}

// SetZoneID sets the zone edge to Zone by id.
func (ec *EquipmentCreate) SetZoneID(id int) *EquipmentCreate {
	ec.mutation.SetZoneID(id)
	return ec
}

// SetNillableZoneID sets the zone edge to Zone by id if the given value is not nil.
func (ec *EquipmentCreate) SetNillableZoneID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetZoneID(*id)
	}
	return ec
}

// SetZone sets the zone edge to Zone.
func (ec *EquipmentCreate) SetZone(z *Zone) *EquipmentCreate {
	return ec.SetZoneID(z.ID)
}

// AddEquipmentrentalIDs adds the equipmentrental edge to Equipmentrental by ids.
func (ec *EquipmentCreate) AddEquipmentrentalIDs(ids ...int) *EquipmentCreate {
	ec.mutation.AddEquipmentrentalIDs(ids...)
	return ec
}

// AddEquipmentrental adds the equipmentrental edges to Equipmentrental.
func (ec *EquipmentCreate) AddEquipmentrental(e ...*Equipmentrental) *EquipmentCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentrentalIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (ec *EquipmentCreate) Mutation() *EquipmentMutation {
	return ec.mutation
}

// Save creates the Equipment in the database.
func (ec *EquipmentCreate) Save(ctx context.Context) (*Equipment, error) {
	if _, ok := ec.mutation.EQUIPMENTNAME(); !ok {
		return nil, &ValidationError{Name: "EQUIPMENTNAME", err: errors.New("ent: missing required field \"EQUIPMENTNAME\"")}
	}
	if _, ok := ec.mutation.EQUIPMENTAMOUNT(); !ok {
		return nil, &ValidationError{Name: "EQUIPMENTAMOUNT", err: errors.New("ent: missing required field \"EQUIPMENTAMOUNT\"")}
	}
	if v, ok := ec.mutation.EQUIPMENTAMOUNT(); ok {
		if err := equipment.EQUIPMENTAMOUNTValidator(v); err != nil {
			return nil, &ValidationError{Name: "EQUIPMENTAMOUNT", err: fmt.Errorf("ent: validator failed for field \"EQUIPMENTAMOUNT\": %w", err)}
		}
	}
	if _, ok := ec.mutation.EQUIPMENTDETAIL(); !ok {
		return nil, &ValidationError{Name: "EQUIPMENTDETAIL", err: errors.New("ent: missing required field \"EQUIPMENTDETAIL\"")}
	}
	if _, ok := ec.mutation.EQUIPMENTDATE(); !ok {
		return nil, &ValidationError{Name: "EQUIPMENTDATE", err: errors.New("ent: missing required field \"EQUIPMENTDATE\"")}
	}
	var (
		err  error
		node *Equipment
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentMutation)
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
func (ec *EquipmentCreate) SaveX(ctx context.Context) *Equipment {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EquipmentCreate) sqlSave(ctx context.Context) (*Equipment, error) {
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

func (ec *EquipmentCreate) createSpec() (*Equipment, *sqlgraph.CreateSpec) {
	var (
		e     = &Equipment{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipment.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.EQUIPMENTNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipment.FieldEQUIPMENTNAME,
		})
		e.EQUIPMENTNAME = value
	}
	if value, ok := ec.mutation.EQUIPMENTAMOUNT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: equipment.FieldEQUIPMENTAMOUNT,
		})
		e.EQUIPMENTAMOUNT = value
	}
	if value, ok := ec.mutation.EQUIPMENTDETAIL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipment.FieldEQUIPMENTDETAIL,
		})
		e.EQUIPMENTDETAIL = value
	}
	if value, ok := ec.mutation.EQUIPMENTDATE(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: equipment.FieldEQUIPMENTDATE,
		})
		e.EQUIPMENTDATE = value
	}
	if nodes := ec.mutation.ClassifierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipment.ClassifierTable,
			Columns: []string{equipment.ClassifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: classifier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipment.EmployeeTable,
			Columns: []string{equipment.EmployeeColumn},
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
	if nodes := ec.mutation.EquipmenttypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipment.EquipmenttypeTable,
			Columns: []string{equipment.EquipmenttypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ZoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipment.ZoneTable,
			Columns: []string{equipment.ZoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: zone.FieldID,
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
			Table:   equipment.EquipmentrentalTable,
			Columns: []string{equipment.EquipmentrentalColumn},
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
	return e, _spec
}