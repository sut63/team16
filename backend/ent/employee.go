// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/G16/app/ent/age"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/position"
	"github.com/G16/app/ent/salary"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EMPLOYEEID holds the value of the "EMPLOYEEID" field.
	EMPLOYEEID string `json:"EMPLOYEEID,omitempty"`
	// EMPLOYEENAME holds the value of the "EMPLOYEENAME" field.
	EMPLOYEENAME string `json:"EMPLOYEENAME,omitempty"`
	// EMPLOYEEADDRESS holds the value of the "EMPLOYEEADDRESS" field.
	EMPLOYEEADDRESS string `json:"EMPLOYEEADDRESS,omitempty"`
	// IDCARDNUMBER holds the value of the "IDCARDNUMBER" field.
	IDCARDNUMBER string `json:"IDCARDNUMBER,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges             EmployeeEdges `json:"edges"`
	age_employee      *int
	position_employee *int
	salary_employee   *int
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// Age holds the value of the age edge.
	Age *Age
	// Position holds the value of the position edge.
	Position *Position
	// Salary holds the value of the salary edge.
	Salary *Salary
	// Payment holds the value of the payment edge.
	Payment []*Payment
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment
	// Bookcourse holds the value of the bookcourse edge.
	Bookcourse []*Bookcourse
	// Equipmentrental holds the value of the equipmentrental edge.
	Equipmentrental []*Equipmentrental
	// Promotion holds the value of the promotion edge.
	Promotion []*Promotion
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// AgeOrErr returns the Age value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeEdges) AgeOrErr() (*Age, error) {
	if e.loadedTypes[0] {
		if e.Age == nil {
			// The edge age was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: age.Label}
		}
		return e.Age, nil
	}
	return nil, &NotLoadedError{edge: "age"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[1] {
		if e.Position == nil {
			// The edge position was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// SalaryOrErr returns the Salary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EmployeeEdges) SalaryOrErr() (*Salary, error) {
	if e.loadedTypes[2] {
		if e.Salary == nil {
			// The edge salary was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: salary.Label}
		}
		return e.Salary, nil
	}
	return nil, &NotLoadedError{edge: "salary"}
}

// PaymentOrErr returns the Payment value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) PaymentOrErr() ([]*Payment, error) {
	if e.loadedTypes[3] {
		return e.Payment, nil
	}
	return nil, &NotLoadedError{edge: "payment"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[4] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// BookcourseOrErr returns the Bookcourse value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) BookcourseOrErr() ([]*Bookcourse, error) {
	if e.loadedTypes[5] {
		return e.Bookcourse, nil
	}
	return nil, &NotLoadedError{edge: "bookcourse"}
}

// EquipmentrentalOrErr returns the Equipmentrental value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) EquipmentrentalOrErr() ([]*Equipmentrental, error) {
	if e.loadedTypes[6] {
		return e.Equipmentrental, nil
	}
	return nil, &NotLoadedError{edge: "equipmentrental"}
}

// PromotionOrErr returns the Promotion value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) PromotionOrErr() ([]*Promotion, error) {
	if e.loadedTypes[7] {
		return e.Promotion, nil
	}
	return nil, &NotLoadedError{edge: "promotion"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // EMPLOYEEID
		&sql.NullString{}, // EMPLOYEENAME
		&sql.NullString{}, // EMPLOYEEADDRESS
		&sql.NullString{}, // IDCARDNUMBER
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Employee) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // age_employee
		&sql.NullInt64{}, // position_employee
		&sql.NullInt64{}, // salary_employee
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(values ...interface{}) error {
	if m, n := len(values), len(employee.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMPLOYEEID", values[0])
	} else if value.Valid {
		e.EMPLOYEEID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMPLOYEENAME", values[1])
	} else if value.Valid {
		e.EMPLOYEENAME = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMPLOYEEADDRESS", values[2])
	} else if value.Valid {
		e.EMPLOYEEADDRESS = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field IDCARDNUMBER", values[3])
	} else if value.Valid {
		e.IDCARDNUMBER = value.String
	}
	values = values[4:]
	if len(values) == len(employee.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field age_employee", value)
		} else if value.Valid {
			e.age_employee = new(int)
			*e.age_employee = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field position_employee", value)
		} else if value.Valid {
			e.position_employee = new(int)
			*e.position_employee = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field salary_employee", value)
		} else if value.Valid {
			e.salary_employee = new(int)
			*e.salary_employee = int(value.Int64)
		}
	}
	return nil
}

// QueryAge queries the age edge of the Employee.
func (e *Employee) QueryAge() *AgeQuery {
	return (&EmployeeClient{config: e.config}).QueryAge(e)
}

// QueryPosition queries the position edge of the Employee.
func (e *Employee) QueryPosition() *PositionQuery {
	return (&EmployeeClient{config: e.config}).QueryPosition(e)
}

// QuerySalary queries the salary edge of the Employee.
func (e *Employee) QuerySalary() *SalaryQuery {
	return (&EmployeeClient{config: e.config}).QuerySalary(e)
}

// QueryPayment queries the payment edge of the Employee.
func (e *Employee) QueryPayment() *PaymentQuery {
	return (&EmployeeClient{config: e.config}).QueryPayment(e)
}

// QueryEquipment queries the equipment edge of the Employee.
func (e *Employee) QueryEquipment() *EquipmentQuery {
	return (&EmployeeClient{config: e.config}).QueryEquipment(e)
}

// QueryBookcourse queries the bookcourse edge of the Employee.
func (e *Employee) QueryBookcourse() *BookcourseQuery {
	return (&EmployeeClient{config: e.config}).QueryBookcourse(e)
}

// QueryEquipmentrental queries the equipmentrental edge of the Employee.
func (e *Employee) QueryEquipmentrental() *EquipmentrentalQuery {
	return (&EmployeeClient{config: e.config}).QueryEquipmentrental(e)
}

// QueryPromotion queries the promotion edge of the Employee.
func (e *Employee) QueryPromotion() *PromotionQuery {
	return (&EmployeeClient{config: e.config}).QueryPromotion(e)
}

// Update returns a builder for updating this Employee.
// Note that, you need to call Employee.Unwrap() before calling this method, if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return (&EmployeeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", EMPLOYEEID=")
	builder.WriteString(e.EMPLOYEEID)
	builder.WriteString(", EMPLOYEENAME=")
	builder.WriteString(e.EMPLOYEENAME)
	builder.WriteString(", EMPLOYEEADDRESS=")
	builder.WriteString(e.EMPLOYEEADDRESS)
	builder.WriteString(", IDCARDNUMBER=")
	builder.WriteString(e.IDCARDNUMBER)
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee

func (e Employees) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
