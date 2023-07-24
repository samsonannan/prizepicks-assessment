package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cage holds the schema definition for the Cage entity.
type Cage struct {
	ent.Schema
}

func (Cage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Cage.
func (Cage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Comment("unique uuid for each cage entity"),
		field.Int64("size").
			Default(0).NonNegative().
			Comment("number of dinosaurs held in cage. defaults to zero"),
		field.Int64("capacity").
			Default(0).NonNegative().
			Comment("capacity for each cage entity. defaults to zero"),
		field.Enum("status").
			Values("ACTIVE", "DOWN").
			Default("ACTIVE").
			Comment("status for cage entity. can be ACTIVE or DOWN only"),
	}
}

// Edges of the Cage.
func (Cage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dinosaurs", Dinosaur.Type).StorageKey(edge.Column("cage_id")),
	}
}
