// Code generated by entc, DO NOT EDIT.

package promotion

const (
	// Label holds the string label denoting the promotion type in the database.
	Label = "promotion"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNAME holds the string denoting the name field in the database.
	FieldNAME = "name"
	// FieldDESC holds the string denoting the desc field in the database.
	FieldDESC = "desc"
	// FieldCODE holds the string denoting the code field in the database.
	FieldCODE = "code"
	// FieldDATE holds the string denoting the date field in the database.
	FieldDATE = "date"

	// EdgePromotiontype holds the string denoting the promotiontype edge name in mutations.
	EdgePromotiontype = "promotiontype"
	// EdgePromotionamount holds the string denoting the promotionamount edge name in mutations.
	EdgePromotionamount = "promotionamount"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"

	// Table holds the table name of the promotion in the database.
	Table = "promotions"
	// PromotiontypeTable is the table the holds the promotiontype relation/edge.
	PromotiontypeTable = "promotions"
	// PromotiontypeInverseTable is the table name for the Promotiontype entity.
	// It exists in this package in order to avoid circular dependency with the "promotiontype" package.
	PromotiontypeInverseTable = "promotiontypes"
	// PromotiontypeColumn is the table column denoting the promotiontype relation/edge.
	PromotiontypeColumn = "promotiontype_promotion"
	// PromotionamountTable is the table the holds the promotionamount relation/edge.
	PromotionamountTable = "promotions"
	// PromotionamountInverseTable is the table name for the Promotionamount entity.
	// It exists in this package in order to avoid circular dependency with the "promotionamount" package.
	PromotionamountInverseTable = "promotionamounts"
	// PromotionamountColumn is the table column denoting the promotionamount relation/edge.
	PromotionamountColumn = "promotionamount_promotion"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "promotions"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_promotion"
	// PaymentTable is the table the holds the payment relation/edge.
	PaymentTable = "payments"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payments"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "promotion_payment"
)

// Columns holds all SQL columns for promotion fields.
var Columns = []string{
	FieldID,
	FieldNAME,
	FieldDESC,
	FieldCODE,
	FieldDATE,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Promotion type.
var ForeignKeys = []string{
	"employee_promotion",
	"promotionamount_promotion",
	"promotiontype_promotion",
}

var (
	// NAMEValidator is a validator for the "NAME" field. It is called by the builders before save.
	NAMEValidator func(string) error
	// DESCValidator is a validator for the "DESC" field. It is called by the builders before save.
	DESCValidator func(string) error
	// CODEValidator is a validator for the "CODE" field. It is called by the builders before save.
	CODEValidator func(string) error
)
