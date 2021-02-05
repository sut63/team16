// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/G16/app/ent/course"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Course is the model entity for the Course schema.
type Course struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// COURSE holds the value of the "COURSE" field.
	COURSE string `json:"COURSE,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseQuery when eager-loading is set.
	Edges CourseEdges `json:"edges"`
}

// CourseEdges holds the relations/edges for other nodes in the graph.
type CourseEdges struct {
	// Bookcourse holds the value of the bookcourse edge.
	Bookcourse []*Bookcourse
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BookcourseOrErr returns the Bookcourse value or an error if the edge
// was not loaded in eager-loading.
func (e CourseEdges) BookcourseOrErr() ([]*Bookcourse, error) {
	if e.loadedTypes[0] {
		return e.Bookcourse, nil
	}
	return nil, &NotLoadedError{edge: "bookcourse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Course) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // COURSE
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Course fields.
func (c *Course) assignValues(values ...interface{}) error {
	if m, n := len(values), len(course.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field COURSE", values[0])
	} else if value.Valid {
		c.COURSE = value.String
	}
	return nil
}

// QueryBookcourse queries the bookcourse edge of the Course.
func (c *Course) QueryBookcourse() *BookcourseQuery {
	return (&CourseClient{config: c.config}).QueryBookcourse(c)
}

// Update returns a builder for updating this Course.
// Note that, you need to call Course.Unwrap() before calling this method, if this Course
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Course) Update() *CourseUpdateOne {
	return (&CourseClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Course) Unwrap() *Course {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Course is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Course) String() string {
	var builder strings.Builder
	builder.WriteString("Course(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", COURSE=")
	builder.WriteString(c.COURSE)
	builder.WriteByte(')')
	return builder.String()
}

// Courses is a parsable slice of Course.
type Courses []*Course

func (c Courses) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
