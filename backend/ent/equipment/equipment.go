// Code generated by entc, DO NOT EDIT.

package equipment

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEQUIPMENTNAME holds the string denoting the equipmentname field in the database.
	FieldEQUIPMENTNAME = "equipmentname"
	// FieldEQUIPMENTAMOUNT holds the string denoting the equipmentamount field in the database.
	FieldEQUIPMENTAMOUNT = "equipmentamount"
	// FieldEQUIPMENTDETAIL holds the string denoting the equipmentdetail field in the database.
	FieldEQUIPMENTDETAIL = "equipmentdetail"
	// FieldEQUIPMENTDATE holds the string denoting the equipmentdate field in the database.
	FieldEQUIPMENTDATE = "equipmentdate"

	// EdgeClassifier holds the string denoting the classifier edge name in mutations.
	EdgeClassifier = "classifier"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgeEquipmenttype holds the string denoting the equipmenttype edge name in mutations.
	EdgeEquipmenttype = "equipmenttype"
	// EdgeZone holds the string denoting the zone edge name in mutations.
	EdgeZone = "zone"
	// EdgeEquipmentrental holds the string denoting the equipmentrental edge name in mutations.
	EdgeEquipmentrental = "equipmentrental"

	// Table holds the table name of the equipment in the database.
	Table = "equipment"
	// ClassifierTable is the table the holds the classifier relation/edge.
	ClassifierTable = "equipment"
	// ClassifierInverseTable is the table name for the Classifier entity.
	// It exists in this package in order to avoid circular dependency with the "classifier" package.
	ClassifierInverseTable = "classifiers"
	// ClassifierColumn is the table column denoting the classifier relation/edge.
	ClassifierColumn = "classifier_equipment"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "equipment"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_equipment"
	// EquipmenttypeTable is the table the holds the equipmenttype relation/edge.
	EquipmenttypeTable = "equipment"
	// EquipmenttypeInverseTable is the table name for the Equipmenttype entity.
	// It exists in this package in order to avoid circular dependency with the "equipmenttype" package.
	EquipmenttypeInverseTable = "equipmenttypes"
	// EquipmenttypeColumn is the table column denoting the equipmenttype relation/edge.
	EquipmenttypeColumn = "equipmenttype_equipment"
	// ZoneTable is the table the holds the zone relation/edge.
	ZoneTable = "equipment"
	// ZoneInverseTable is the table name for the Zone entity.
	// It exists in this package in order to avoid circular dependency with the "zone" package.
	ZoneInverseTable = "zones"
	// ZoneColumn is the table column denoting the zone relation/edge.
	ZoneColumn = "zone_equipment"
	// EquipmentrentalTable is the table the holds the equipmentrental relation/edge.
	EquipmentrentalTable = "equipmentrentals"
	// EquipmentrentalInverseTable is the table name for the Equipmentrental entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentrental" package.
	EquipmentrentalInverseTable = "equipmentrentals"
	// EquipmentrentalColumn is the table column denoting the equipmentrental relation/edge.
	EquipmentrentalColumn = "equipment_equipmentrental"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldEQUIPMENTNAME,
	FieldEQUIPMENTAMOUNT,
	FieldEQUIPMENTDETAIL,
	FieldEQUIPMENTDATE,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Equipment type.
var ForeignKeys = []string{
	"classifier_equipment",
	"employee_equipment",
	"equipmenttype_equipment",
	"zone_equipment",
}
