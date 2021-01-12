package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
)
 
// Classifier schema.
type Classifier struct {
   ent.Schema
}
 
// Fields of the Classifier.
func (Classifier) Fields() []ent.Field {
   return []ent.Field{
       field.String("EQUIPMENTCLASSIFIER"),
   }
}
 
// Edges of the Classifier.
func (Classifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type),
	}
}