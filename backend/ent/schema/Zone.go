package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
)
 
// Zone schema.
type Zone struct {
   ent.Schema
}
 
// Fields of the Zone.
func (Zone) Fields() []ent.Field {
   return []ent.Field{
       field.String("EQUIPMENTZONE"),
   }
}
 
// Edges of the Zone.
func (Zone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type),
	}
}
