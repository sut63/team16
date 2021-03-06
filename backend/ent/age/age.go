// Code generated by entc, DO NOT EDIT.

package age

const (
	// Label holds the string label denoting the age type in the database.
	Label = "age"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAGE holds the string denoting the age field in the database.
	FieldAGE = "age"

	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"

	// Table holds the table name of the age in the database.
	Table = "ages"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "employees"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "age_employee"
)

// Columns holds all SQL columns for age fields.
var Columns = []string{
	FieldID,
	FieldAGE,
}
