package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
)
 
// Equipmenttype schema.
type Equipmenttype struct {
   ent.Schema
}
 
// Fields of the Equipmenttype.
func (Equipmenttype) Fields() []ent.Field {
   return []ent.Field{
		field.String("EQUIPMENTTYPE"),
   }
}
 
// Edges of the Equipmenttype.
func (Equipmenttype) Edges() []ent.Edge {
	return []ent.Edge{
      edge.To("equipment", Equipment.Type),
      edge.To("equipmentrental", Equipmentrental.Type),
	}
}