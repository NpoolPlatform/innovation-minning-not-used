// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/project"
	"github.com/google/uuid"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetGoodID sets the "good_id" field.
func (pu *ProjectUpdate) SetGoodID(u uuid.UUID) *ProjectUpdate {
	pu.mutation.SetGoodID(u)
	return pu
}

// SetTeamID sets the "team_id" field.
func (pu *ProjectUpdate) SetTeamID(u uuid.UUID) *ProjectUpdate {
	pu.mutation.SetTeamID(u)
	return pu
}

// SetCapitalIds sets the "capital_ids" field.
func (pu *ProjectUpdate) SetCapitalIds(u []uuid.UUID) *ProjectUpdate {
	pu.mutation.SetCapitalIds(u)
	return pu
}

// SetIntroduction sets the "introduction" field.
func (pu *ProjectUpdate) SetIntroduction(s string) *ProjectUpdate {
	pu.mutation.SetIntroduction(s)
	return pu
}

// SetLogo sets the "logo" field.
func (pu *ProjectUpdate) SetLogo(s string) *ProjectUpdate {
	pu.mutation.SetLogo(s)
	return pu
}

// SetCreateAt sets the "create_at" field.
func (pu *ProjectUpdate) SetCreateAt(u uint32) *ProjectUpdate {
	pu.mutation.ResetCreateAt()
	pu.mutation.SetCreateAt(u)
	return pu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreateAt(u *uint32) *ProjectUpdate {
	if u != nil {
		pu.SetCreateAt(*u)
	}
	return pu
}

// AddCreateAt adds u to the "create_at" field.
func (pu *ProjectUpdate) AddCreateAt(u uint32) *ProjectUpdate {
	pu.mutation.AddCreateAt(u)
	return pu
}

// SetUpdateAt sets the "update_at" field.
func (pu *ProjectUpdate) SetUpdateAt(u uint32) *ProjectUpdate {
	pu.mutation.ResetUpdateAt()
	pu.mutation.SetUpdateAt(u)
	return pu
}

// AddUpdateAt adds u to the "update_at" field.
func (pu *ProjectUpdate) AddUpdateAt(u uint32) *ProjectUpdate {
	pu.mutation.AddUpdateAt(u)
	return pu
}

// SetDeleteAt sets the "delete_at" field.
func (pu *ProjectUpdate) SetDeleteAt(u uint32) *ProjectUpdate {
	pu.mutation.ResetDeleteAt()
	pu.mutation.SetDeleteAt(u)
	return pu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDeleteAt(u *uint32) *ProjectUpdate {
	if u != nil {
		pu.SetDeleteAt(*u)
	}
	return pu
}

// AddDeleteAt adds u to the "delete_at" field.
func (pu *ProjectUpdate) AddDeleteAt(u uint32) *ProjectUpdate {
	pu.mutation.AddDeleteAt(u)
	return pu
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() {
	if _, ok := pu.mutation.UpdateAt(); !ok {
		v := project.UpdateDefaultUpdateAt()
		pu.mutation.SetUpdateAt(v)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: project.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := pu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: project.FieldGoodID,
		})
	}
	if value, ok := pu.mutation.TeamID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: project.FieldTeamID,
		})
	}
	if value, ok := pu.mutation.CapitalIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: project.FieldCapitalIds,
		})
	}
	if value, ok := pu.mutation.Introduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldIntroduction,
		})
	}
	if value, ok := pu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldLogo,
		})
	}
	if value, ok := pu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldDeleteAt,
		})
	}
	if value, ok := pu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetGoodID sets the "good_id" field.
func (puo *ProjectUpdateOne) SetGoodID(u uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetGoodID(u)
	return puo
}

// SetTeamID sets the "team_id" field.
func (puo *ProjectUpdateOne) SetTeamID(u uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetTeamID(u)
	return puo
}

// SetCapitalIds sets the "capital_ids" field.
func (puo *ProjectUpdateOne) SetCapitalIds(u []uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetCapitalIds(u)
	return puo
}

// SetIntroduction sets the "introduction" field.
func (puo *ProjectUpdateOne) SetIntroduction(s string) *ProjectUpdateOne {
	puo.mutation.SetIntroduction(s)
	return puo
}

// SetLogo sets the "logo" field.
func (puo *ProjectUpdateOne) SetLogo(s string) *ProjectUpdateOne {
	puo.mutation.SetLogo(s)
	return puo
}

// SetCreateAt sets the "create_at" field.
func (puo *ProjectUpdateOne) SetCreateAt(u uint32) *ProjectUpdateOne {
	puo.mutation.ResetCreateAt()
	puo.mutation.SetCreateAt(u)
	return puo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreateAt(u *uint32) *ProjectUpdateOne {
	if u != nil {
		puo.SetCreateAt(*u)
	}
	return puo
}

// AddCreateAt adds u to the "create_at" field.
func (puo *ProjectUpdateOne) AddCreateAt(u uint32) *ProjectUpdateOne {
	puo.mutation.AddCreateAt(u)
	return puo
}

// SetUpdateAt sets the "update_at" field.
func (puo *ProjectUpdateOne) SetUpdateAt(u uint32) *ProjectUpdateOne {
	puo.mutation.ResetUpdateAt()
	puo.mutation.SetUpdateAt(u)
	return puo
}

// AddUpdateAt adds u to the "update_at" field.
func (puo *ProjectUpdateOne) AddUpdateAt(u uint32) *ProjectUpdateOne {
	puo.mutation.AddUpdateAt(u)
	return puo
}

// SetDeleteAt sets the "delete_at" field.
func (puo *ProjectUpdateOne) SetDeleteAt(u uint32) *ProjectUpdateOne {
	puo.mutation.ResetDeleteAt()
	puo.mutation.SetDeleteAt(u)
	return puo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDeleteAt(u *uint32) *ProjectUpdateOne {
	if u != nil {
		puo.SetDeleteAt(*u)
	}
	return puo
}

// AddDeleteAt adds u to the "delete_at" field.
func (puo *ProjectUpdateOne) AddDeleteAt(u uint32) *ProjectUpdateOne {
	puo.mutation.AddDeleteAt(u)
	return puo
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateAt(); !ok {
		v := project.UpdateDefaultUpdateAt()
		puo.mutation.SetUpdateAt(v)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: project.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Project.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if value, ok := puo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: project.FieldGoodID,
		})
	}
	if value, ok := puo.mutation.TeamID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: project.FieldTeamID,
		})
	}
	if value, ok := puo.mutation.CapitalIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: project.FieldCapitalIds,
		})
	}
	if value, ok := puo.mutation.Introduction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldIntroduction,
		})
	}
	if value, ok := puo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldLogo,
		})
	}
	if value, ok := puo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldDeleteAt,
		})
	}
	if value, ok := puo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: project.FieldDeleteAt,
		})
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
