// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/trendanalysis"
)

// TrendAnalysisDelete is the builder for deleting a TrendAnalysis entity.
type TrendAnalysisDelete struct {
	config
	hooks    []Hook
	mutation *TrendAnalysisMutation
}

// Where appends a list predicates to the TrendAnalysisDelete builder.
func (tad *TrendAnalysisDelete) Where(ps ...predicate.TrendAnalysis) *TrendAnalysisDelete {
	tad.mutation.Where(ps...)
	return tad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tad *TrendAnalysisDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tad.hooks) == 0 {
		affected, err = tad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrendAnalysisMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tad.mutation = mutation
			affected, err = tad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tad.hooks) - 1; i >= 0; i-- {
			if tad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tad *TrendAnalysisDelete) ExecX(ctx context.Context) int {
	n, err := tad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tad *TrendAnalysisDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: trendanalysis.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: trendanalysis.FieldID,
			},
		},
	}
	if ps := tad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tad.driver, _spec)
}

// TrendAnalysisDeleteOne is the builder for deleting a single TrendAnalysis entity.
type TrendAnalysisDeleteOne struct {
	tad *TrendAnalysisDelete
}

// Exec executes the deletion query.
func (tado *TrendAnalysisDeleteOne) Exec(ctx context.Context) error {
	n, err := tado.tad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{trendanalysis.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tado *TrendAnalysisDeleteOne) ExecX(ctx context.Context) {
	tado.tad.ExecX(ctx)
}
