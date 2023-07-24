// Code generated by ent, DO NOT EDIT.

package cage

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cage type in the database.
	Label = "cage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldCapacity holds the string denoting the capacity field in the database.
	FieldCapacity = "capacity"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeDinosaurs holds the string denoting the dinosaurs edge name in mutations.
	EdgeDinosaurs = "dinosaurs"
	// Table holds the table name of the cage in the database.
	Table = "cages"
	// DinosaursTable is the table that holds the dinosaurs relation/edge.
	DinosaursTable = "dinosaurs"
	// DinosaursInverseTable is the table name for the Dinosaur entity.
	// It exists in this package in order to avoid circular dependency with the "dinosaur" package.
	DinosaursInverseTable = "dinosaurs"
	// DinosaursColumn is the table column denoting the dinosaurs relation/edge.
	DinosaursColumn = "cage_id"
)

// Columns holds all SQL columns for cage fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSize,
	FieldCapacity,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize int64
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int64) error
	// DefaultCapacity holds the default value on creation for the "capacity" field.
	DefaultCapacity int64
	// CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	CapacityValidator func(int64) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusACTIVE is the default value of the Status enum.
const DefaultStatus = StatusACTIVE

// Status values.
const (
	StatusACTIVE Status = "ACTIVE"
	StatusDOWN   Status = "DOWN"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusACTIVE, StatusDOWN:
		return nil
	default:
		return fmt.Errorf("cage: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Cage queries.
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

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByCapacity orders the results by the capacity field.
func ByCapacity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapacity, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDinosaursCount orders the results by dinosaurs count.
func ByDinosaursCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDinosaursStep(), opts...)
	}
}

// ByDinosaurs orders the results by dinosaurs terms.
func ByDinosaurs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDinosaursStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDinosaursStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DinosaursInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DinosaursTable, DinosaursColumn),
	)
}