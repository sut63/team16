// Code generated by entc, DO NOT EDIT.

package promotionamount

const (
	// Label holds the string label denoting the promotionamount type in the database.
	Label = "promotionamount"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAMOUNT holds the string denoting the amount field in the database.
	FieldAMOUNT = "amount"

	// EdgePromotion holds the string denoting the promotion edge name in mutations.
	EdgePromotion = "promotion"

	// Table holds the table name of the promotionamount in the database.
	Table = "promotionamounts"
	// PromotionTable is the table the holds the promotion relation/edge.
	PromotionTable = "promotions"
	// PromotionInverseTable is the table name for the Promotion entity.
	// It exists in this package in order to avoid circular dependency with the "promotion" package.
	PromotionInverseTable = "promotions"
	// PromotionColumn is the table column denoting the promotion relation/edge.
	PromotionColumn = "promotionamount_promotion"
)

// Columns holds all SQL columns for promotionamount fields.
var Columns = []string{
	FieldID,
	FieldAMOUNT,
}

var (
	// AMOUNTValidator is a validator for the "AMOUNT" field. It is called by the builders before save.
	AMOUNTValidator func(int) error
)