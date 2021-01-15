// Code generated by entc, DO NOT EDIT.

package equipmentrental

const (
	// Label holds the string label denoting the equipmentrental type in the database.
	Label = "equipmentrental"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRENTALAMOUNT holds the string denoting the rentalamount field in the database.
	FieldRENTALAMOUNT = "rentalamount"
	// FieldRENTALDATE holds the string denoting the rentaldate field in the database.
	FieldRENTALDATE = "rentaldate"
	// FieldRETURNDATE holds the string denoting the returndate field in the database.
	FieldRETURNDATE = "returndate"

	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "member"
	// EdgeEquipmenttype holds the string denoting the equipmenttype edge name in mutations.
	EdgeEquipmenttype = "equipmenttype"

	// Table holds the table name of the equipmentrental in the database.
	Table = "equipmentrentals"
	// EquipmentTable is the table the holds the equipment relation/edge.
	EquipmentTable = "equipmentrentals"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_equipmentrental"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "equipmentrentals"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_equipmentrental"
	// MemberTable is the table the holds the member relation/edge.
	MemberTable = "equipmentrentals"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the member relation/edge.
	MemberColumn = "member_equipmentrental"
	// EquipmenttypeTable is the table the holds the equipmenttype relation/edge.
	EquipmenttypeTable = "equipmentrentals"
	// EquipmenttypeInverseTable is the table name for the Equipmenttype entity.
	// It exists in this package in order to avoid circular dependency with the "equipmenttype" package.
	EquipmenttypeInverseTable = "equipmenttypes"
	// EquipmenttypeColumn is the table column denoting the equipmenttype relation/edge.
	EquipmenttypeColumn = "equipmenttype_equipmentrental"
)

// Columns holds all SQL columns for equipmentrental fields.
var Columns = []string{
	FieldID,
	FieldRENTALAMOUNT,
	FieldRENTALDATE,
	FieldRETURNDATE,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Equipmentrental type.
var ForeignKeys = []string{
	"employee_equipmentrental",
	"equipment_equipmentrental",
	"equipmenttype_equipmentrental",
	"member_equipmentrental",
}

var (
	// RENTALAMOUNTValidator is a validator for the "RENTALAMOUNT" field. It is called by the builders before save.
	RENTALAMOUNTValidator func(int) error
)
