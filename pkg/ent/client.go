// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cage is the client for interacting with the Cage builders.
	Cage *CageClient
	// Dinosaur is the client for interacting with the Dinosaur builders.
	Dinosaur *DinosaurClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cage = NewCageClient(c.config)
	c.Dinosaur = NewDinosaurClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Cage:     NewCageClient(cfg),
		Dinosaur: NewDinosaurClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Cage:     NewCageClient(cfg),
		Dinosaur: NewDinosaurClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cage.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Cage.Use(hooks...)
	c.Dinosaur.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Cage.Intercept(interceptors...)
	c.Dinosaur.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CageMutation:
		return c.Cage.mutate(ctx, m)
	case *DinosaurMutation:
		return c.Dinosaur.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CageClient is a client for the Cage schema.
type CageClient struct {
	config
}

// NewCageClient returns a client for the Cage from the given config.
func NewCageClient(c config) *CageClient {
	return &CageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cage.Hooks(f(g(h())))`.
func (c *CageClient) Use(hooks ...Hook) {
	c.hooks.Cage = append(c.hooks.Cage, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cage.Intercept(f(g(h())))`.
func (c *CageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cage = append(c.inters.Cage, interceptors...)
}

// Create returns a builder for creating a Cage entity.
func (c *CageClient) Create() *CageCreate {
	mutation := newCageMutation(c.config, OpCreate)
	return &CageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cage entities.
func (c *CageClient) CreateBulk(builders ...*CageCreate) *CageCreateBulk {
	return &CageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cage.
func (c *CageClient) Update() *CageUpdate {
	mutation := newCageMutation(c.config, OpUpdate)
	return &CageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CageClient) UpdateOne(ca *Cage) *CageUpdateOne {
	mutation := newCageMutation(c.config, OpUpdateOne, withCage(ca))
	return &CageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CageClient) UpdateOneID(id uuid.UUID) *CageUpdateOne {
	mutation := newCageMutation(c.config, OpUpdateOne, withCageID(id))
	return &CageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cage.
func (c *CageClient) Delete() *CageDelete {
	mutation := newCageMutation(c.config, OpDelete)
	return &CageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CageClient) DeleteOne(ca *Cage) *CageDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CageClient) DeleteOneID(id uuid.UUID) *CageDeleteOne {
	builder := c.Delete().Where(cage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CageDeleteOne{builder}
}

// Query returns a query builder for Cage.
func (c *CageClient) Query() *CageQuery {
	return &CageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCage},
		inters: c.Interceptors(),
	}
}

// Get returns a Cage entity by its id.
func (c *CageClient) Get(ctx context.Context, id uuid.UUID) (*Cage, error) {
	return c.Query().Where(cage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CageClient) GetX(ctx context.Context, id uuid.UUID) *Cage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDinosaurs queries the dinosaurs edge of a Cage.
func (c *CageClient) QueryDinosaurs(ca *Cage) *DinosaurQuery {
	query := (&DinosaurClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cage.Table, cage.FieldID, id),
			sqlgraph.To(dinosaur.Table, dinosaur.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cage.DinosaursTable, cage.DinosaursColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CageClient) Hooks() []Hook {
	return c.hooks.Cage
}

// Interceptors returns the client interceptors.
func (c *CageClient) Interceptors() []Interceptor {
	return c.inters.Cage
}

func (c *CageClient) mutate(ctx context.Context, m *CageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Cage mutation op: %q", m.Op())
	}
}

// DinosaurClient is a client for the Dinosaur schema.
type DinosaurClient struct {
	config
}

// NewDinosaurClient returns a client for the Dinosaur from the given config.
func NewDinosaurClient(c config) *DinosaurClient {
	return &DinosaurClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dinosaur.Hooks(f(g(h())))`.
func (c *DinosaurClient) Use(hooks ...Hook) {
	c.hooks.Dinosaur = append(c.hooks.Dinosaur, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `dinosaur.Intercept(f(g(h())))`.
func (c *DinosaurClient) Intercept(interceptors ...Interceptor) {
	c.inters.Dinosaur = append(c.inters.Dinosaur, interceptors...)
}

// Create returns a builder for creating a Dinosaur entity.
func (c *DinosaurClient) Create() *DinosaurCreate {
	mutation := newDinosaurMutation(c.config, OpCreate)
	return &DinosaurCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dinosaur entities.
func (c *DinosaurClient) CreateBulk(builders ...*DinosaurCreate) *DinosaurCreateBulk {
	return &DinosaurCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dinosaur.
func (c *DinosaurClient) Update() *DinosaurUpdate {
	mutation := newDinosaurMutation(c.config, OpUpdate)
	return &DinosaurUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DinosaurClient) UpdateOne(d *Dinosaur) *DinosaurUpdateOne {
	mutation := newDinosaurMutation(c.config, OpUpdateOne, withDinosaur(d))
	return &DinosaurUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DinosaurClient) UpdateOneID(id uuid.UUID) *DinosaurUpdateOne {
	mutation := newDinosaurMutation(c.config, OpUpdateOne, withDinosaurID(id))
	return &DinosaurUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dinosaur.
func (c *DinosaurClient) Delete() *DinosaurDelete {
	mutation := newDinosaurMutation(c.config, OpDelete)
	return &DinosaurDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DinosaurClient) DeleteOne(d *Dinosaur) *DinosaurDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DinosaurClient) DeleteOneID(id uuid.UUID) *DinosaurDeleteOne {
	builder := c.Delete().Where(dinosaur.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DinosaurDeleteOne{builder}
}

// Query returns a query builder for Dinosaur.
func (c *DinosaurClient) Query() *DinosaurQuery {
	return &DinosaurQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDinosaur},
		inters: c.Interceptors(),
	}
}

// Get returns a Dinosaur entity by its id.
func (c *DinosaurClient) Get(ctx context.Context, id uuid.UUID) (*Dinosaur, error) {
	return c.Query().Where(dinosaur.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DinosaurClient) GetX(ctx context.Context, id uuid.UUID) *Dinosaur {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCage queries the cage edge of a Dinosaur.
func (c *DinosaurClient) QueryCage(d *Dinosaur) *CageQuery {
	query := (&CageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dinosaur.Table, dinosaur.FieldID, id),
			sqlgraph.To(cage.Table, cage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dinosaur.CageTable, dinosaur.CageColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DinosaurClient) Hooks() []Hook {
	return c.hooks.Dinosaur
}

// Interceptors returns the client interceptors.
func (c *DinosaurClient) Interceptors() []Interceptor {
	return c.inters.Dinosaur
}

func (c *DinosaurClient) mutate(ctx context.Context, m *DinosaurMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DinosaurCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DinosaurUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DinosaurUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DinosaurDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Dinosaur mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Cage, Dinosaur []ent.Hook
	}
	inters struct {
		Cage, Dinosaur []ent.Interceptor
	}
)
