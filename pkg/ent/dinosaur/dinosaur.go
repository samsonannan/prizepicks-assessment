// Code generated by ent, DO NOT EDIT.

package dinosaur

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the dinosaur type in the database.
	Label = "dinosaur"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSpecies holds the string denoting the species field in the database.
	FieldSpecies = "species"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// EdgeCage holds the string denoting the cage edge name in mutations.
	EdgeCage = "cage"
	// Table holds the table name of the dinosaur in the database.
	Table = "dinosaurs"
	// CageTable is the table that holds the cage relation/edge.
	CageTable = "dinosaurs"
	// CageInverseTable is the table name for the Cage entity.
	// It exists in this package in order to avoid circular dependency with the "cage" package.
	CageInverseTable = "cages"
	// CageColumn is the table column denoting the cage relation/edge.
	CageColumn = "cage_id"
)

// Columns holds all SQL columns for dinosaur fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldSpecies,
	FieldGroup,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "dinosaurs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"cage_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SpeciesValidator is a validator for the "species" field. It is called by the builders before save.
	SpeciesValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Group defines the type for the "group" enum field.
type Group string

// Group values.
const (
	GroupHERBIVORE Group = "HERBIVORE"
	GroupCARNIVORE Group = "CARNIVORE"
)

func (gr Group) String() string {
	return string(gr)
}

// GroupValidator is a validator for the "group" field enum values. It is called by the builders before save.
func GroupValidator(gr Group) error {
	switch gr {
	case GroupHERBIVORE, GroupCARNIVORE:
		return nil
	default:
		return fmt.Errorf("dinosaur: invalid enum value for group field: %q", gr)
	}
}

// OrderOption defines the ordering options for the Dinosaur queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySpecies orders the results by the species field.
func BySpecies(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpecies, opts...).ToFunc()
}

// ByGroup orders the results by the group field.
func ByGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroup, opts...).ToFunc()
}

// ByCageField orders the results by cage field.
func ByCageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCageStep(), sql.OrderByField(field, opts...))
	}
}
func newCageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CageTable, CageColumn),
	)
}
