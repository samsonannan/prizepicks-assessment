// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
)

// CageCreate is the builder for creating a Cage entity.
type CageCreate struct {
	config
	mutation *CageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CageCreate) SetCreatedAt(t time.Time) *CageCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CageCreate) SetNillableCreatedAt(t *time.Time) *CageCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CageCreate) SetUpdatedAt(t time.Time) *CageCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CageCreate) SetNillableUpdatedAt(t *time.Time) *CageCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetSize sets the "size" field.
func (cc *CageCreate) SetSize(i int64) *CageCreate {
	cc.mutation.SetSize(i)
	return cc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (cc *CageCreate) SetNillableSize(i *int64) *CageCreate {
	if i != nil {
		cc.SetSize(*i)
	}
	return cc
}

// SetCapacity sets the "capacity" field.
func (cc *CageCreate) SetCapacity(i int64) *CageCreate {
	cc.mutation.SetCapacity(i)
	return cc
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (cc *CageCreate) SetNillableCapacity(i *int64) *CageCreate {
	if i != nil {
		cc.SetCapacity(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CageCreate) SetStatus(c cage.Status) *CageCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CageCreate) SetNillableStatus(c *cage.Status) *CageCreate {
	if c != nil {
		cc.SetStatus(*c)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CageCreate) SetID(u uuid.UUID) *CageCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CageCreate) SetNillableID(u *uuid.UUID) *CageCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddDinosaurIDs adds the "dinosaurs" edge to the Dinosaur entity by IDs.
func (cc *CageCreate) AddDinosaurIDs(ids ...uuid.UUID) *CageCreate {
	cc.mutation.AddDinosaurIDs(ids...)
	return cc
}

// AddDinosaurs adds the "dinosaurs" edges to the Dinosaur entity.
func (cc *CageCreate) AddDinosaurs(d ...*Dinosaur) *CageCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cc.AddDinosaurIDs(ids...)
}

// Mutation returns the CageMutation object of the builder.
func (cc *CageCreate) Mutation() *CageMutation {
	return cc.mutation
}

// Save creates the Cage in the database.
func (cc *CageCreate) Save(ctx context.Context) (*Cage, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CageCreate) SaveX(ctx context.Context) *Cage {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CageCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CageCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CageCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := cage.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := cage.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Size(); !ok {
		v := cage.DefaultSize
		cc.mutation.SetSize(v)
	}
	if _, ok := cc.mutation.Capacity(); !ok {
		v := cage.DefaultCapacity
		cc.mutation.SetCapacity(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := cage.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := cage.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CageCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Cage.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Cage.updated_at"`)}
	}
	if _, ok := cc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Cage.size"`)}
	}
	if v, ok := cc.mutation.Size(); ok {
		if err := cage.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "Cage.size": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "Cage.capacity"`)}
	}
	if v, ok := cc.mutation.Capacity(); ok {
		if err := cage.CapacityValidator(v); err != nil {
			return &ValidationError{Name: "capacity", err: fmt.Errorf(`ent: validator failed for field "Cage.capacity": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Cage.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := cage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Cage.status": %w`, err)}
		}
	}
	return nil
}

func (cc *CageCreate) sqlSave(ctx context.Context) (*Cage, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CageCreate) createSpec() (*Cage, *sqlgraph.CreateSpec) {
	var (
		_node = &Cage{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(cage.Table, sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(cage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(cage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Size(); ok {
		_spec.SetField(cage.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := cc.mutation.Capacity(); ok {
		_spec.SetField(cage.FieldCapacity, field.TypeInt64, value)
		_node.Capacity = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(cage.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := cc.mutation.DinosaursIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cage.DinosaursTable,
			Columns: []string{cage.DinosaursColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dinosaur.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CageCreateBulk is the builder for creating many Cage entities in bulk.
type CageCreateBulk struct {
	config
	builders []*CageCreate
}

// Save creates the Cage entities in the database.
func (ccb *CageCreateBulk) Save(ctx context.Context) ([]*Cage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cage, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CageCreateBulk) SaveX(ctx context.Context) []*Cage {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CageCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CageCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
