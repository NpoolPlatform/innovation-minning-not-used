// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/capital"
	"github.com/google/uuid"
)

// CapitalCreate is the builder for creating a Capital entity.
type CapitalCreate struct {
	config
	mutation *CapitalMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cc *CapitalCreate) SetName(s string) *CapitalCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIntroduction sets the "introduction" field.
func (cc *CapitalCreate) SetIntroduction(s string) *CapitalCreate {
	cc.mutation.SetIntroduction(s)
	return cc
}

// SetLogo sets the "logo" field.
func (cc *CapitalCreate) SetLogo(s string) *CapitalCreate {
	cc.mutation.SetLogo(s)
	return cc
}

// SetCreateAt sets the "create_at" field.
func (cc *CapitalCreate) SetCreateAt(u uint32) *CapitalCreate {
	cc.mutation.SetCreateAt(u)
	return cc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cc *CapitalCreate) SetNillableCreateAt(u *uint32) *CapitalCreate {
	if u != nil {
		cc.SetCreateAt(*u)
	}
	return cc
}

// SetUpdateAt sets the "update_at" field.
func (cc *CapitalCreate) SetUpdateAt(u uint32) *CapitalCreate {
	cc.mutation.SetUpdateAt(u)
	return cc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (cc *CapitalCreate) SetNillableUpdateAt(u *uint32) *CapitalCreate {
	if u != nil {
		cc.SetUpdateAt(*u)
	}
	return cc
}

// SetDeleteAt sets the "delete_at" field.
func (cc *CapitalCreate) SetDeleteAt(u uint32) *CapitalCreate {
	cc.mutation.SetDeleteAt(u)
	return cc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cc *CapitalCreate) SetNillableDeleteAt(u *uint32) *CapitalCreate {
	if u != nil {
		cc.SetDeleteAt(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CapitalCreate) SetID(u uuid.UUID) *CapitalCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CapitalMutation object of the builder.
func (cc *CapitalCreate) Mutation() *CapitalMutation {
	return cc.mutation
}

// Save creates the Capital in the database.
func (cc *CapitalCreate) Save(ctx context.Context) (*Capital, error) {
	var (
		err  error
		node *Capital
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CapitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CapitalCreate) SaveX(ctx context.Context) *Capital {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CapitalCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CapitalCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CapitalCreate) defaults() {
	if _, ok := cc.mutation.CreateAt(); !ok {
		v := capital.DefaultCreateAt()
		cc.mutation.SetCreateAt(v)
	}
	if _, ok := cc.mutation.UpdateAt(); !ok {
		v := capital.DefaultUpdateAt()
		cc.mutation.SetUpdateAt(v)
	}
	if _, ok := cc.mutation.DeleteAt(); !ok {
		v := capital.DefaultDeleteAt()
		cc.mutation.SetDeleteAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := capital.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CapitalCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := cc.mutation.Introduction(); !ok {
		return &ValidationError{Name: "introduction", err: errors.New(`ent: missing required field "introduction"`)}
	}
	if _, ok := cc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`ent: missing required field "logo"`)}
	}
	if _, ok := cc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := cc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := cc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (cc *CapitalCreate) sqlSave(ctx context.Context) (*Capital, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (cc *CapitalCreate) createSpec() (*Capital, *sqlgraph.CreateSpec) {
	var (
		_node = &Capital{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: capital.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: capital.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Introduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldIntroduction,
		})
		_node.Introduction = value
	}
	if value, ok := cc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capital.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := cc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := cc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := cc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: capital.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Capital.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CapitalUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CapitalCreate) OnConflict(opts ...sql.ConflictOption) *CapitalUpsertOne {
	cc.conflict = opts
	return &CapitalUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Capital.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CapitalCreate) OnConflictColumns(columns ...string) *CapitalUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CapitalUpsertOne{
		create: cc,
	}
}

type (
	// CapitalUpsertOne is the builder for "upsert"-ing
	//  one Capital node.
	CapitalUpsertOne struct {
		create *CapitalCreate
	}

	// CapitalUpsert is the "OnConflict" setter.
	CapitalUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *CapitalUpsert) SetName(v string) *CapitalUpsert {
	u.Set(capital.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateName() *CapitalUpsert {
	u.SetExcluded(capital.FieldName)
	return u
}

// SetIntroduction sets the "introduction" field.
func (u *CapitalUpsert) SetIntroduction(v string) *CapitalUpsert {
	u.Set(capital.FieldIntroduction, v)
	return u
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateIntroduction() *CapitalUpsert {
	u.SetExcluded(capital.FieldIntroduction)
	return u
}

// SetLogo sets the "logo" field.
func (u *CapitalUpsert) SetLogo(v string) *CapitalUpsert {
	u.Set(capital.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateLogo() *CapitalUpsert {
	u.SetExcluded(capital.FieldLogo)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CapitalUpsert) SetCreateAt(v uint32) *CapitalUpsert {
	u.Set(capital.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateCreateAt() *CapitalUpsert {
	u.SetExcluded(capital.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *CapitalUpsert) SetUpdateAt(v uint32) *CapitalUpsert {
	u.Set(capital.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateUpdateAt() *CapitalUpsert {
	u.SetExcluded(capital.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *CapitalUpsert) SetDeleteAt(v uint32) *CapitalUpsert {
	u.Set(capital.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CapitalUpsert) UpdateDeleteAt() *CapitalUpsert {
	u.SetExcluded(capital.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Capital.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(capital.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CapitalUpsertOne) UpdateNewValues() *CapitalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(capital.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Capital.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CapitalUpsertOne) Ignore() *CapitalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CapitalUpsertOne) DoNothing() *CapitalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CapitalCreate.OnConflict
// documentation for more info.
func (u *CapitalUpsertOne) Update(set func(*CapitalUpsert)) *CapitalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CapitalUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CapitalUpsertOne) SetName(v string) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateName() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateName()
	})
}

// SetIntroduction sets the "introduction" field.
func (u *CapitalUpsertOne) SetIntroduction(v string) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetIntroduction(v)
	})
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateIntroduction() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateIntroduction()
	})
}

// SetLogo sets the "logo" field.
func (u *CapitalUpsertOne) SetLogo(v string) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateLogo() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateLogo()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CapitalUpsertOne) SetCreateAt(v uint32) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateCreateAt() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CapitalUpsertOne) SetUpdateAt(v uint32) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateUpdateAt() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CapitalUpsertOne) SetDeleteAt(v uint32) *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CapitalUpsertOne) UpdateDeleteAt() *CapitalUpsertOne {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CapitalUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CapitalCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CapitalUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CapitalUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CapitalUpsertOne.ID is not supported by MySQL driver. Use CapitalUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CapitalUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CapitalCreateBulk is the builder for creating many Capital entities in bulk.
type CapitalCreateBulk struct {
	config
	builders []*CapitalCreate
	conflict []sql.ConflictOption
}

// Save creates the Capital entities in the database.
func (ccb *CapitalCreateBulk) Save(ctx context.Context) ([]*Capital, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Capital, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CapitalMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
func (ccb *CapitalCreateBulk) SaveX(ctx context.Context) []*Capital {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CapitalCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CapitalCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Capital.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CapitalUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CapitalCreateBulk) OnConflict(opts ...sql.ConflictOption) *CapitalUpsertBulk {
	ccb.conflict = opts
	return &CapitalUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Capital.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CapitalCreateBulk) OnConflictColumns(columns ...string) *CapitalUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CapitalUpsertBulk{
		create: ccb,
	}
}

// CapitalUpsertBulk is the builder for "upsert"-ing
// a bulk of Capital nodes.
type CapitalUpsertBulk struct {
	create *CapitalCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Capital.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(capital.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CapitalUpsertBulk) UpdateNewValues() *CapitalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(capital.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Capital.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CapitalUpsertBulk) Ignore() *CapitalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CapitalUpsertBulk) DoNothing() *CapitalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CapitalCreateBulk.OnConflict
// documentation for more info.
func (u *CapitalUpsertBulk) Update(set func(*CapitalUpsert)) *CapitalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CapitalUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CapitalUpsertBulk) SetName(v string) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateName() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateName()
	})
}

// SetIntroduction sets the "introduction" field.
func (u *CapitalUpsertBulk) SetIntroduction(v string) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetIntroduction(v)
	})
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateIntroduction() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateIntroduction()
	})
}

// SetLogo sets the "logo" field.
func (u *CapitalUpsertBulk) SetLogo(v string) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateLogo() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateLogo()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CapitalUpsertBulk) SetCreateAt(v uint32) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateCreateAt() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CapitalUpsertBulk) SetUpdateAt(v uint32) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateUpdateAt() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CapitalUpsertBulk) SetDeleteAt(v uint32) *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CapitalUpsertBulk) UpdateDeleteAt() *CapitalUpsertBulk {
	return u.Update(func(s *CapitalUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CapitalUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CapitalCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CapitalCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CapitalUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
