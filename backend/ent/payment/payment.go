// Code generated by entc, DO NOT EDIT.

package payment

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPAYMENTAMOUNT holds the string denoting the paymentamount field in the database.
	FieldPAYMENTAMOUNT = "paymentamount"
	// FieldPAYMENTDATE holds the string denoting the paymentdate field in the database.
	FieldPAYMENTDATE = "paymentdate"

	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "member"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgePaymenttype holds the string denoting the paymenttype edge name in mutations.
	EdgePaymenttype = "paymenttype"
	// EdgePromotion holds the string denoting the promotion edge name in mutations.
	EdgePromotion = "promotion"

	// Table holds the table name of the payment in the database.
	Table = "payments"
	// MemberTable is the table the holds the member relation/edge.
	MemberTable = "payments"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the member relation/edge.
	MemberColumn = "member_payment"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "payments"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_payment"
	// PaymenttypeTable is the table the holds the paymenttype relation/edge.
	PaymenttypeTable = "payments"
	// PaymenttypeInverseTable is the table name for the Paymenttype entity.
	// It exists in this package in order to avoid circular dependency with the "paymenttype" package.
	PaymenttypeInverseTable = "paymenttypes"
	// PaymenttypeColumn is the table column denoting the paymenttype relation/edge.
	PaymenttypeColumn = "paymenttype_payment"
	// PromotionTable is the table the holds the promotion relation/edge.
	PromotionTable = "payments"
	// PromotionInverseTable is the table name for the Promotion entity.
	// It exists in this package in order to avoid circular dependency with the "promotion" package.
	PromotionInverseTable = "promotions"
	// PromotionColumn is the table column denoting the promotion relation/edge.
	PromotionColumn = "promotion_payment"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldPAYMENTAMOUNT,
	FieldPAYMENTDATE,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Payment type.
var ForeignKeys = []string{
	"employee_payment",
	"member_payment",
	"paymenttype_payment",
	"promotion_payment",
}
