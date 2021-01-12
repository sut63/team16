// Code generated by entc, DO NOT EDIT.

package promotiontype

const (
	// Label holds the string label denoting the promotiontype type in the database.
	Label = "promotiontype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTYPE holds the string denoting the type field in the database.
	FieldTYPE = "type"

	// EdgePromotion holds the string denoting the promotion edge name in mutations.
	EdgePromotion = "promotion"

	// Table holds the table name of the promotiontype in the database.
	Table = "promotiontypes"
	// PromotionTable is the table the holds the promotion relation/edge.
	PromotionTable = "promotions"
	// PromotionInverseTable is the table name for the Promotion entity.
	// It exists in this package in order to avoid circular dependency with the "promotion" package.
	PromotionInverseTable = "promotions"
	// PromotionColumn is the table column denoting the promotion relation/edge.
	PromotionColumn = "promotiontype_promotion"
)

// Columns holds all SQL columns for promotiontype fields.
var Columns = []string{
	FieldID,
	FieldTYPE,
}

var (
	// TYPEValidator is a validator for the "TYPE" field. It is called by the builders before save.
	TYPEValidator func(string) error
)