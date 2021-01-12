// Code generated by entc, DO NOT EDIT.

package classifier

const (
	// Label holds the string label denoting the classifier type in the database.
	Label = "classifier"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEQUIPMENTCLASSIFIER holds the string denoting the equipmentclassifier field in the database.
	FieldEQUIPMENTCLASSIFIER = "equipmentclassifier"

	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"

	// Table holds the table name of the classifier in the database.
	Table = "classifiers"
	// EquipmentTable is the table the holds the equipment relation/edge.
	EquipmentTable = "equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "classifier_equipment"
)

// Columns holds all SQL columns for classifier fields.
var Columns = []string{
	FieldID,
	FieldEQUIPMENTCLASSIFIER,
}
