// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCage     = "Cage"
	TypeDinosaur = "Dinosaur"
)

// CageMutation represents an operation that mutates the Cage nodes in the graph.
type CageMutation struct {
	config
	op               Op
	typ              string
	id               *uuid.UUID
	created_at       *time.Time
	updated_at       *time.Time
	size             *int64
	addsize          *int64
	capacity         *int64
	addcapacity      *int64
	status           *cage.Status
	clearedFields    map[string]struct{}
	dinosaurs        map[uuid.UUID]struct{}
	removeddinosaurs map[uuid.UUID]struct{}
	cleareddinosaurs bool
	done             bool
	oldValue         func(context.Context) (*Cage, error)
	predicates       []predicate.Cage
}

var _ ent.Mutation = (*CageMutation)(nil)

// cageOption allows management of the mutation configuration using functional options.
type cageOption func(*CageMutation)

// newCageMutation creates new mutation for the Cage entity.
func newCageMutation(c config, op Op, opts ...cageOption) *CageMutation {
	m := &CageMutation{
		config:        c,
		op:            op,
		typ:           TypeCage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCageID sets the ID field of the mutation.
func withCageID(id uuid.UUID) cageOption {
	return func(m *CageMutation) {
		var (
			err   error
			once  sync.Once
			value *Cage
		)
		m.oldValue = func(ctx context.Context) (*Cage, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Cage.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCage sets the old Cage of the mutation.
func withCage(node *Cage) cageOption {
	return func(m *CageMutation) {
		m.oldValue = func(context.Context) (*Cage, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Cage entities.
func (m *CageMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CageMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CageMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Cage.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *CageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Cage entity.
// If the Cage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *CageMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *CageMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Cage entity.
// If the Cage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CageMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *CageMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetSize sets the "size" field.
func (m *CageMutation) SetSize(i int64) {
	m.size = &i
	m.addsize = nil
}

// Size returns the value of the "size" field in the mutation.
func (m *CageMutation) Size() (r int64, exists bool) {
	v := m.size
	if v == nil {
		return
	}
	return *v, true
}

// OldSize returns the old "size" field's value of the Cage entity.
// If the Cage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CageMutation) OldSize(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSize is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSize requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSize: %w", err)
	}
	return oldValue.Size, nil
}

// AddSize adds i to the "size" field.
func (m *CageMutation) AddSize(i int64) {
	if m.addsize != nil {
		*m.addsize += i
	} else {
		m.addsize = &i
	}
}

// AddedSize returns the value that was added to the "size" field in this mutation.
func (m *CageMutation) AddedSize() (r int64, exists bool) {
	v := m.addsize
	if v == nil {
		return
	}
	return *v, true
}

// ResetSize resets all changes to the "size" field.
func (m *CageMutation) ResetSize() {
	m.size = nil
	m.addsize = nil
}

// SetCapacity sets the "capacity" field.
func (m *CageMutation) SetCapacity(i int64) {
	m.capacity = &i
	m.addcapacity = nil
}

// Capacity returns the value of the "capacity" field in the mutation.
func (m *CageMutation) Capacity() (r int64, exists bool) {
	v := m.capacity
	if v == nil {
		return
	}
	return *v, true
}

// OldCapacity returns the old "capacity" field's value of the Cage entity.
// If the Cage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CageMutation) OldCapacity(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCapacity is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCapacity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCapacity: %w", err)
	}
	return oldValue.Capacity, nil
}

// AddCapacity adds i to the "capacity" field.
func (m *CageMutation) AddCapacity(i int64) {
	if m.addcapacity != nil {
		*m.addcapacity += i
	} else {
		m.addcapacity = &i
	}
}

// AddedCapacity returns the value that was added to the "capacity" field in this mutation.
func (m *CageMutation) AddedCapacity() (r int64, exists bool) {
	v := m.addcapacity
	if v == nil {
		return
	}
	return *v, true
}

// ResetCapacity resets all changes to the "capacity" field.
func (m *CageMutation) ResetCapacity() {
	m.capacity = nil
	m.addcapacity = nil
}

// SetStatus sets the "status" field.
func (m *CageMutation) SetStatus(c cage.Status) {
	m.status = &c
}

// Status returns the value of the "status" field in the mutation.
func (m *CageMutation) Status() (r cage.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Cage entity.
// If the Cage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CageMutation) OldStatus(ctx context.Context) (v cage.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *CageMutation) ResetStatus() {
	m.status = nil
}

// AddDinosaurIDs adds the "dinosaurs" edge to the Dinosaur entity by ids.
func (m *CageMutation) AddDinosaurIDs(ids ...uuid.UUID) {
	if m.dinosaurs == nil {
		m.dinosaurs = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.dinosaurs[ids[i]] = struct{}{}
	}
}

// ClearDinosaurs clears the "dinosaurs" edge to the Dinosaur entity.
func (m *CageMutation) ClearDinosaurs() {
	m.cleareddinosaurs = true
}

// DinosaursCleared reports if the "dinosaurs" edge to the Dinosaur entity was cleared.
func (m *CageMutation) DinosaursCleared() bool {
	return m.cleareddinosaurs
}

// RemoveDinosaurIDs removes the "dinosaurs" edge to the Dinosaur entity by IDs.
func (m *CageMutation) RemoveDinosaurIDs(ids ...uuid.UUID) {
	if m.removeddinosaurs == nil {
		m.removeddinosaurs = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.dinosaurs, ids[i])
		m.removeddinosaurs[ids[i]] = struct{}{}
	}
}

// RemovedDinosaurs returns the removed IDs of the "dinosaurs" edge to the Dinosaur entity.
func (m *CageMutation) RemovedDinosaursIDs() (ids []uuid.UUID) {
	for id := range m.removeddinosaurs {
		ids = append(ids, id)
	}
	return
}

// DinosaursIDs returns the "dinosaurs" edge IDs in the mutation.
func (m *CageMutation) DinosaursIDs() (ids []uuid.UUID) {
	for id := range m.dinosaurs {
		ids = append(ids, id)
	}
	return
}

// ResetDinosaurs resets all changes to the "dinosaurs" edge.
func (m *CageMutation) ResetDinosaurs() {
	m.dinosaurs = nil
	m.cleareddinosaurs = false
	m.removeddinosaurs = nil
}

// Where appends a list predicates to the CageMutation builder.
func (m *CageMutation) Where(ps ...predicate.Cage) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Cage, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Cage).
func (m *CageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CageMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, cage.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, cage.FieldUpdatedAt)
	}
	if m.size != nil {
		fields = append(fields, cage.FieldSize)
	}
	if m.capacity != nil {
		fields = append(fields, cage.FieldCapacity)
	}
	if m.status != nil {
		fields = append(fields, cage.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case cage.FieldCreatedAt:
		return m.CreatedAt()
	case cage.FieldUpdatedAt:
		return m.UpdatedAt()
	case cage.FieldSize:
		return m.Size()
	case cage.FieldCapacity:
		return m.Capacity()
	case cage.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case cage.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case cage.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case cage.FieldSize:
		return m.OldSize(ctx)
	case cage.FieldCapacity:
		return m.OldCapacity(ctx)
	case cage.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown Cage field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case cage.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case cage.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case cage.FieldSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSize(v)
		return nil
	case cage.FieldCapacity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCapacity(v)
		return nil
	case cage.FieldStatus:
		v, ok := value.(cage.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Cage field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CageMutation) AddedFields() []string {
	var fields []string
	if m.addsize != nil {
		fields = append(fields, cage.FieldSize)
	}
	if m.addcapacity != nil {
		fields = append(fields, cage.FieldCapacity)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case cage.FieldSize:
		return m.AddedSize()
	case cage.FieldCapacity:
		return m.AddedCapacity()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case cage.FieldSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSize(v)
		return nil
	case cage.FieldCapacity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCapacity(v)
		return nil
	}
	return fmt.Errorf("unknown Cage numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Cage nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CageMutation) ResetField(name string) error {
	switch name {
	case cage.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case cage.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case cage.FieldSize:
		m.ResetSize()
		return nil
	case cage.FieldCapacity:
		m.ResetCapacity()
		return nil
	case cage.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown Cage field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CageMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.dinosaurs != nil {
		edges = append(edges, cage.EdgeDinosaurs)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case cage.EdgeDinosaurs:
		ids := make([]ent.Value, 0, len(m.dinosaurs))
		for id := range m.dinosaurs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeddinosaurs != nil {
		edges = append(edges, cage.EdgeDinosaurs)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CageMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case cage.EdgeDinosaurs:
		ids := make([]ent.Value, 0, len(m.removeddinosaurs))
		for id := range m.removeddinosaurs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareddinosaurs {
		edges = append(edges, cage.EdgeDinosaurs)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CageMutation) EdgeCleared(name string) bool {
	switch name {
	case cage.EdgeDinosaurs:
		return m.cleareddinosaurs
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CageMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Cage unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CageMutation) ResetEdge(name string) error {
	switch name {
	case cage.EdgeDinosaurs:
		m.ResetDinosaurs()
		return nil
	}
	return fmt.Errorf("unknown Cage edge %s", name)
}

// DinosaurMutation represents an operation that mutates the Dinosaur nodes in the graph.
type DinosaurMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	species       *string
	group         *dinosaur.Group
	clearedFields map[string]struct{}
	cage          *uuid.UUID
	clearedcage   bool
	done          bool
	oldValue      func(context.Context) (*Dinosaur, error)
	predicates    []predicate.Dinosaur
}

var _ ent.Mutation = (*DinosaurMutation)(nil)

// dinosaurOption allows management of the mutation configuration using functional options.
type dinosaurOption func(*DinosaurMutation)

// newDinosaurMutation creates new mutation for the Dinosaur entity.
func newDinosaurMutation(c config, op Op, opts ...dinosaurOption) *DinosaurMutation {
	m := &DinosaurMutation{
		config:        c,
		op:            op,
		typ:           TypeDinosaur,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDinosaurID sets the ID field of the mutation.
func withDinosaurID(id uuid.UUID) dinosaurOption {
	return func(m *DinosaurMutation) {
		var (
			err   error
			once  sync.Once
			value *Dinosaur
		)
		m.oldValue = func(ctx context.Context) (*Dinosaur, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Dinosaur.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDinosaur sets the old Dinosaur of the mutation.
func withDinosaur(node *Dinosaur) dinosaurOption {
	return func(m *DinosaurMutation) {
		m.oldValue = func(context.Context) (*Dinosaur, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DinosaurMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DinosaurMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Dinosaur entities.
func (m *DinosaurMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DinosaurMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DinosaurMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Dinosaur.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *DinosaurMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *DinosaurMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Dinosaur entity.
// If the Dinosaur object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DinosaurMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *DinosaurMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *DinosaurMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *DinosaurMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Dinosaur entity.
// If the Dinosaur object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DinosaurMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *DinosaurMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *DinosaurMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *DinosaurMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Dinosaur entity.
// If the Dinosaur object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DinosaurMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *DinosaurMutation) ResetName() {
	m.name = nil
}

// SetSpecies sets the "species" field.
func (m *DinosaurMutation) SetSpecies(s string) {
	m.species = &s
}

// Species returns the value of the "species" field in the mutation.
func (m *DinosaurMutation) Species() (r string, exists bool) {
	v := m.species
	if v == nil {
		return
	}
	return *v, true
}

// OldSpecies returns the old "species" field's value of the Dinosaur entity.
// If the Dinosaur object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DinosaurMutation) OldSpecies(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSpecies is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSpecies requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpecies: %w", err)
	}
	return oldValue.Species, nil
}

// ResetSpecies resets all changes to the "species" field.
func (m *DinosaurMutation) ResetSpecies() {
	m.species = nil
}

// SetGroup sets the "group" field.
func (m *DinosaurMutation) SetGroup(d dinosaur.Group) {
	m.group = &d
}

// Group returns the value of the "group" field in the mutation.
func (m *DinosaurMutation) Group() (r dinosaur.Group, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroup returns the old "group" field's value of the Dinosaur entity.
// If the Dinosaur object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DinosaurMutation) OldGroup(ctx context.Context) (v dinosaur.Group, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGroup is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGroup requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroup: %w", err)
	}
	return oldValue.Group, nil
}

// ResetGroup resets all changes to the "group" field.
func (m *DinosaurMutation) ResetGroup() {
	m.group = nil
}

// SetCageID sets the "cage" edge to the Cage entity by id.
func (m *DinosaurMutation) SetCageID(id uuid.UUID) {
	m.cage = &id
}

// ClearCage clears the "cage" edge to the Cage entity.
func (m *DinosaurMutation) ClearCage() {
	m.clearedcage = true
}

// CageCleared reports if the "cage" edge to the Cage entity was cleared.
func (m *DinosaurMutation) CageCleared() bool {
	return m.clearedcage
}

// CageID returns the "cage" edge ID in the mutation.
func (m *DinosaurMutation) CageID() (id uuid.UUID, exists bool) {
	if m.cage != nil {
		return *m.cage, true
	}
	return
}

// CageIDs returns the "cage" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// CageID instead. It exists only for internal usage by the builders.
func (m *DinosaurMutation) CageIDs() (ids []uuid.UUID) {
	if id := m.cage; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCage resets all changes to the "cage" edge.
func (m *DinosaurMutation) ResetCage() {
	m.cage = nil
	m.clearedcage = false
}

// Where appends a list predicates to the DinosaurMutation builder.
func (m *DinosaurMutation) Where(ps ...predicate.Dinosaur) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the DinosaurMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *DinosaurMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Dinosaur, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *DinosaurMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *DinosaurMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Dinosaur).
func (m *DinosaurMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DinosaurMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, dinosaur.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, dinosaur.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, dinosaur.FieldName)
	}
	if m.species != nil {
		fields = append(fields, dinosaur.FieldSpecies)
	}
	if m.group != nil {
		fields = append(fields, dinosaur.FieldGroup)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DinosaurMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case dinosaur.FieldCreatedAt:
		return m.CreatedAt()
	case dinosaur.FieldUpdatedAt:
		return m.UpdatedAt()
	case dinosaur.FieldName:
		return m.Name()
	case dinosaur.FieldSpecies:
		return m.Species()
	case dinosaur.FieldGroup:
		return m.Group()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DinosaurMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case dinosaur.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case dinosaur.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case dinosaur.FieldName:
		return m.OldName(ctx)
	case dinosaur.FieldSpecies:
		return m.OldSpecies(ctx)
	case dinosaur.FieldGroup:
		return m.OldGroup(ctx)
	}
	return nil, fmt.Errorf("unknown Dinosaur field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DinosaurMutation) SetField(name string, value ent.Value) error {
	switch name {
	case dinosaur.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case dinosaur.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case dinosaur.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case dinosaur.FieldSpecies:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpecies(v)
		return nil
	case dinosaur.FieldGroup:
		v, ok := value.(dinosaur.Group)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroup(v)
		return nil
	}
	return fmt.Errorf("unknown Dinosaur field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DinosaurMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DinosaurMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DinosaurMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Dinosaur numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DinosaurMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DinosaurMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DinosaurMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Dinosaur nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DinosaurMutation) ResetField(name string) error {
	switch name {
	case dinosaur.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case dinosaur.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case dinosaur.FieldName:
		m.ResetName()
		return nil
	case dinosaur.FieldSpecies:
		m.ResetSpecies()
		return nil
	case dinosaur.FieldGroup:
		m.ResetGroup()
		return nil
	}
	return fmt.Errorf("unknown Dinosaur field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DinosaurMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cage != nil {
		edges = append(edges, dinosaur.EdgeCage)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DinosaurMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case dinosaur.EdgeCage:
		if id := m.cage; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DinosaurMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DinosaurMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DinosaurMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcage {
		edges = append(edges, dinosaur.EdgeCage)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DinosaurMutation) EdgeCleared(name string) bool {
	switch name {
	case dinosaur.EdgeCage:
		return m.clearedcage
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DinosaurMutation) ClearEdge(name string) error {
	switch name {
	case dinosaur.EdgeCage:
		m.ClearCage()
		return nil
	}
	return fmt.Errorf("unknown Dinosaur unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DinosaurMutation) ResetEdge(name string) error {
	switch name {
	case dinosaur.EdgeCage:
		m.ResetCage()
		return nil
	}
	return fmt.Errorf("unknown Dinosaur edge %s", name)
}
