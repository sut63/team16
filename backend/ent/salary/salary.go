// Code generated by entc, DO NOT EDIT.

package salary

const (
	// Label holds the string label denoting the salary type in the database.
	Label = "salary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSALARY holds the string denoting the salary field in the database.
	FieldSALARY = "salary"

	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"

	// Table holds the table name of the salary in the database.
	Table = "salaries"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "employees"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "salary_employee"
)

// Columns holds all SQL columns for salary fields.
var Columns = []string{
	FieldID,
	FieldSALARY,
}
