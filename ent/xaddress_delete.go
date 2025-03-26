// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yumo001/simple-learn-rpc/ent/predicate"
	"github.com/yumo001/simple-learn-rpc/ent/xaddress"
)

// XAddressDelete is the builder for deleting a XAddress entity.
type XAddressDelete struct {
	config
	hooks    []Hook
	mutation *XAddressMutation
}

// Where appends a list predicates to the XAddressDelete builder.
func (xd *XAddressDelete) Where(ps ...predicate.XAddress) *XAddressDelete {
	xd.mutation.Where(ps...)
	return xd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (xd *XAddressDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, xd.sqlExec, xd.mutation, xd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (xd *XAddressDelete) ExecX(ctx context.Context) int {
	n, err := xd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (xd *XAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(xaddress.Table, sqlgraph.NewFieldSpec(xaddress.FieldID, field.TypeUint64))
	if ps := xd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, xd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	xd.mutation.done = true
	return affected, err
}

// XAddressDeleteOne is the builder for deleting a single XAddress entity.
type XAddressDeleteOne struct {
	xd *XAddressDelete
}

// Where appends a list predicates to the XAddressDelete builder.
func (xdo *XAddressDeleteOne) Where(ps ...predicate.XAddress) *XAddressDeleteOne {
	xdo.xd.mutation.Where(ps...)
	return xdo
}

// Exec executes the deletion query.
func (xdo *XAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := xdo.xd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{xaddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (xdo *XAddressDeleteOne) ExecX(ctx context.Context) {
	if err := xdo.Exec(ctx); err != nil {
		panic(err)
	}
}
