// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/predicate"
)

// DinosaurUpdate is the builder for updating Dinosaur entities.
type DinosaurUpdate struct {
	config
	hooks    []Hook
	mutation *DinosaurMutation
}

// Where appends a list predicates to the DinosaurUpdate builder.
func (du *DinosaurUpdate) Where(ps ...predicate.Dinosaur) *DinosaurUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DinosaurUpdate) SetUpdatedAt(t time.Time) *DinosaurUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetName sets the "name" field.
func (du *DinosaurUpdate) SetName(s string) *DinosaurUpdate {
	du.mutation.SetName(s)
	return du
}

// SetSpecies sets the "species" field.
func (du *DinosaurUpdate) SetSpecies(s string) *DinosaurUpdate {
	du.mutation.SetSpecies(s)
	return du
}

// SetGroup sets the "group" field.
func (du *DinosaurUpdate) SetGroup(d dinosaur.Group) *DinosaurUpdate {
	du.mutation.SetGroup(d)
	return du
}

// SetCageID sets the "cage" edge to the Cage entity by ID.
func (du *DinosaurUpdate) SetCageID(id uuid.UUID) *DinosaurUpdate {
	du.mutation.SetCageID(id)
	return du
}

// SetNillableCageID sets the "cage" edge to the Cage entity by ID if the given value is not nil.
func (du *DinosaurUpdate) SetNillableCageID(id *uuid.UUID) *DinosaurUpdate {
	if id != nil {
		du = du.SetCageID(*id)
	}
	return du
}

// SetCage sets the "cage" edge to the Cage entity.
func (du *DinosaurUpdate) SetCage(c *Cage) *DinosaurUpdate {
	return du.SetCageID(c.ID)
}

// Mutation returns the DinosaurMutation object of the builder.
func (du *DinosaurUpdate) Mutation() *DinosaurMutation {
	return du.mutation
}

// ClearCage clears the "cage" edge to the Cage entity.
func (du *DinosaurUpdate) ClearCage() *DinosaurUpdate {
	du.mutation.ClearCage()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DinosaurUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DinosaurUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DinosaurUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DinosaurUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DinosaurUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := dinosaur.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DinosaurUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := dinosaur.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.name": %w`, err)}
		}
	}
	if v, ok := du.mutation.Species(); ok {
		if err := dinosaur.SpeciesValidator(v); err != nil {
			return &ValidationError{Name: "species", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.species": %w`, err)}
		}
	}
	if v, ok := du.mutation.Group(); ok {
		if err := dinosaur.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.group": %w`, err)}
		}
	}
	return nil
}

func (du *DinosaurUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dinosaur.Table, dinosaur.Columns, sqlgraph.NewFieldSpec(dinosaur.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(dinosaur.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(dinosaur.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Species(); ok {
		_spec.SetField(dinosaur.FieldSpecies, field.TypeString, value)
	}
	if value, ok := du.mutation.Group(); ok {
		_spec.SetField(dinosaur.FieldGroup, field.TypeEnum, value)
	}
	if du.mutation.CageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dinosaur.CageTable,
			Columns: []string{dinosaur.CageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.CageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dinosaur.CageTable,
			Columns: []string{dinosaur.CageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dinosaur.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DinosaurUpdateOne is the builder for updating a single Dinosaur entity.
type DinosaurUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DinosaurMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DinosaurUpdateOne) SetUpdatedAt(t time.Time) *DinosaurUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetName sets the "name" field.
func (duo *DinosaurUpdateOne) SetName(s string) *DinosaurUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetSpecies sets the "species" field.
func (duo *DinosaurUpdateOne) SetSpecies(s string) *DinosaurUpdateOne {
	duo.mutation.SetSpecies(s)
	return duo
}

// SetGroup sets the "group" field.
func (duo *DinosaurUpdateOne) SetGroup(d dinosaur.Group) *DinosaurUpdateOne {
	duo.mutation.SetGroup(d)
	return duo
}

// SetCageID sets the "cage" edge to the Cage entity by ID.
func (duo *DinosaurUpdateOne) SetCageID(id uuid.UUID) *DinosaurUpdateOne {
	duo.mutation.SetCageID(id)
	return duo
}

// SetNillableCageID sets the "cage" edge to the Cage entity by ID if the given value is not nil.
func (duo *DinosaurUpdateOne) SetNillableCageID(id *uuid.UUID) *DinosaurUpdateOne {
	if id != nil {
		duo = duo.SetCageID(*id)
	}
	return duo
}

// SetCage sets the "cage" edge to the Cage entity.
func (duo *DinosaurUpdateOne) SetCage(c *Cage) *DinosaurUpdateOne {
	return duo.SetCageID(c.ID)
}

// Mutation returns the DinosaurMutation object of the builder.
func (duo *DinosaurUpdateOne) Mutation() *DinosaurMutation {
	return duo.mutation
}

// ClearCage clears the "cage" edge to the Cage entity.
func (duo *DinosaurUpdateOne) ClearCage() *DinosaurUpdateOne {
	duo.mutation.ClearCage()
	return duo
}

// Where appends a list predicates to the DinosaurUpdate builder.
func (duo *DinosaurUpdateOne) Where(ps ...predicate.Dinosaur) *DinosaurUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DinosaurUpdateOne) Select(field string, fields ...string) *DinosaurUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dinosaur entity.
func (duo *DinosaurUpdateOne) Save(ctx context.Context) (*Dinosaur, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DinosaurUpdateOne) SaveX(ctx context.Context) *Dinosaur {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DinosaurUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DinosaurUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DinosaurUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := dinosaur.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DinosaurUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := dinosaur.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.name": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Species(); ok {
		if err := dinosaur.SpeciesValidator(v); err != nil {
			return &ValidationError{Name: "species", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.species": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Group(); ok {
		if err := dinosaur.GroupValidator(v); err != nil {
			return &ValidationError{Name: "group", err: fmt.Errorf(`ent: validator failed for field "Dinosaur.group": %w`, err)}
		}
	}
	return nil
}

func (duo *DinosaurUpdateOne) sqlSave(ctx context.Context) (_node *Dinosaur, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dinosaur.Table, dinosaur.Columns, sqlgraph.NewFieldSpec(dinosaur.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dinosaur.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dinosaur.FieldID)
		for _, f := range fields {
			if !dinosaur.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dinosaur.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(dinosaur.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(dinosaur.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Species(); ok {
		_spec.SetField(dinosaur.FieldSpecies, field.TypeString, value)
	}
	if value, ok := duo.mutation.Group(); ok {
		_spec.SetField(dinosaur.FieldGroup, field.TypeEnum, value)
	}
	if duo.mutation.CageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dinosaur.CageTable,
			Columns: []string{dinosaur.CageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.CageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dinosaur.CageTable,
			Columns: []string{dinosaur.CageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dinosaur{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dinosaur.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
