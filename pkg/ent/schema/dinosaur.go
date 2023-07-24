package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Dinosaur holds the schema definition for the Dinosaur entity.
type Dinosaur struct {
	ent.Schema
}

func (Dinosaur) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Dinosaur.
func (Dinosaur) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Comment("unique uuid for each dinosaur entity"),
		field.String("name").
			NotEmpty().
			Unique().
			Comment("name of dinosaur. must not be empty"),
		field.String("species").
			NotEmpty().
			Comment("species dinosaur belongs. must not be empty"),
		field.Enum("group").
			Values("HERBIVORE", "CARNIVORE").
			Comment("group defines eating class based on species i.e HERBIVORE, CARNIVORE"),
	}
}

// Edges of the Dinosaur.
func (Dinosaur) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cage", Cage.Type).Ref("dinosaurs").Unique(),
	}
}
