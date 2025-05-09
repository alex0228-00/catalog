// Code generated by ent, DO NOT EDIT.

package ent

import (
	"catalog/src/datastore/ent/predicate"
	"catalog/src/datastore/ent/system"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemUpdate is the builder for updating System entities.
type SystemUpdate struct {
	config
	hooks    []Hook
	mutation *SystemMutation
}

// Where appends a list predicates to the SystemUpdate builder.
func (su *SystemUpdate) Where(ps ...predicate.System) *SystemUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetType sets the "type" field.
func (su *SystemUpdate) SetType(s string) *SystemUpdate {
	su.mutation.SetType(s)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *SystemUpdate) SetNillableType(s *string) *SystemUpdate {
	if s != nil {
		su.SetType(*s)
	}
	return su
}

// SetHost sets the "host" field.
func (su *SystemUpdate) SetHost(s string) *SystemUpdate {
	su.mutation.SetHost(s)
	return su
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (su *SystemUpdate) SetNillableHost(s *string) *SystemUpdate {
	if s != nil {
		su.SetHost(*s)
	}
	return su
}

// SetUniqueIdentifier sets the "unique_identifier" field.
func (su *SystemUpdate) SetUniqueIdentifier(s string) *SystemUpdate {
	su.mutation.SetUniqueIdentifier(s)
	return su
}

// SetNillableUniqueIdentifier sets the "unique_identifier" field if the given value is not nil.
func (su *SystemUpdate) SetNillableUniqueIdentifier(s *string) *SystemUpdate {
	if s != nil {
		su.SetUniqueIdentifier(*s)
	}
	return su
}

// SetCredentials sets the "credentials" field.
func (su *SystemUpdate) SetCredentials(s string) *SystemUpdate {
	su.mutation.SetCredentials(s)
	return su
}

// SetNillableCredentials sets the "credentials" field if the given value is not nil.
func (su *SystemUpdate) SetNillableCredentials(s *string) *SystemUpdate {
	if s != nil {
		su.SetCredentials(*s)
	}
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *SystemUpdate) SetCreatedBy(s string) *SystemUpdate {
	su.mutation.SetCreatedBy(s)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *SystemUpdate) SetNillableCreatedBy(s *string) *SystemUpdate {
	if s != nil {
		su.SetCreatedBy(*s)
	}
	return su
}

// SetUpdatedBy sets the "updated_by" field.
func (su *SystemUpdate) SetUpdatedBy(s string) *SystemUpdate {
	su.mutation.SetUpdatedBy(s)
	return su
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (su *SystemUpdate) SetNillableUpdatedBy(s *string) *SystemUpdate {
	if s != nil {
		su.SetUpdatedBy(*s)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SystemUpdate) SetUpdatedAt(t time.Time) *SystemUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetDeletedBy sets the "deleted_by" field.
func (su *SystemUpdate) SetDeletedBy(s string) *SystemUpdate {
	su.mutation.SetDeletedBy(s)
	return su
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (su *SystemUpdate) SetNillableDeletedBy(s *string) *SystemUpdate {
	if s != nil {
		su.SetDeletedBy(*s)
	}
	return su
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (su *SystemUpdate) ClearDeletedBy() *SystemUpdate {
	su.mutation.ClearDeletedBy()
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SystemUpdate) SetDeletedAt(t time.Time) *SystemUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SystemUpdate) SetNillableDeletedAt(t *time.Time) *SystemUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (su *SystemUpdate) ClearDeletedAt() *SystemUpdate {
	su.mutation.ClearDeletedAt()
	return su
}

// Mutation returns the SystemMutation object of the builder.
func (su *SystemUpdate) Mutation() *SystemMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SystemUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SystemUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SystemUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SystemUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SystemUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := system.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SystemUpdate) check() error {
	if v, ok := su.mutation.GetType(); ok {
		if err := system.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "System.type": %w`, err)}
		}
	}
	if v, ok := su.mutation.Host(); ok {
		if err := system.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "System.host": %w`, err)}
		}
	}
	if v, ok := su.mutation.Credentials(); ok {
		if err := system.CredentialsValidator(v); err != nil {
			return &ValidationError{Name: "credentials", err: fmt.Errorf(`ent: validator failed for field "System.credentials": %w`, err)}
		}
	}
	if v, ok := su.mutation.CreatedBy(); ok {
		if err := system.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "System.created_by": %w`, err)}
		}
	}
	if v, ok := su.mutation.UpdatedBy(); ok {
		if err := system.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "System.updated_by": %w`, err)}
		}
	}
	return nil
}

func (su *SystemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(system.Table, system.Columns, sqlgraph.NewFieldSpec(system.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(system.FieldType, field.TypeString, value)
	}
	if value, ok := su.mutation.Host(); ok {
		_spec.SetField(system.FieldHost, field.TypeString, value)
	}
	if value, ok := su.mutation.UniqueIdentifier(); ok {
		_spec.SetField(system.FieldUniqueIdentifier, field.TypeString, value)
	}
	if value, ok := su.mutation.Credentials(); ok {
		_spec.SetField(system.FieldCredentials, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(system.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := su.mutation.UpdatedBy(); ok {
		_spec.SetField(system.FieldUpdatedBy, field.TypeString, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(system.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.DeletedBy(); ok {
		_spec.SetField(system.FieldDeletedBy, field.TypeString, value)
	}
	if su.mutation.DeletedByCleared() {
		_spec.ClearField(system.FieldDeletedBy, field.TypeString)
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.SetField(system.FieldDeletedAt, field.TypeTime, value)
	}
	if su.mutation.DeletedAtCleared() {
		_spec.ClearField(system.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SystemUpdateOne is the builder for updating a single System entity.
type SystemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SystemMutation
}

// SetType sets the "type" field.
func (suo *SystemUpdateOne) SetType(s string) *SystemUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableType(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetType(*s)
	}
	return suo
}

// SetHost sets the "host" field.
func (suo *SystemUpdateOne) SetHost(s string) *SystemUpdateOne {
	suo.mutation.SetHost(s)
	return suo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableHost(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetHost(*s)
	}
	return suo
}

// SetUniqueIdentifier sets the "unique_identifier" field.
func (suo *SystemUpdateOne) SetUniqueIdentifier(s string) *SystemUpdateOne {
	suo.mutation.SetUniqueIdentifier(s)
	return suo
}

// SetNillableUniqueIdentifier sets the "unique_identifier" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableUniqueIdentifier(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetUniqueIdentifier(*s)
	}
	return suo
}

// SetCredentials sets the "credentials" field.
func (suo *SystemUpdateOne) SetCredentials(s string) *SystemUpdateOne {
	suo.mutation.SetCredentials(s)
	return suo
}

// SetNillableCredentials sets the "credentials" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableCredentials(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetCredentials(*s)
	}
	return suo
}

// SetCreatedBy sets the "created_by" field.
func (suo *SystemUpdateOne) SetCreatedBy(s string) *SystemUpdateOne {
	suo.mutation.SetCreatedBy(s)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableCreatedBy(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetCreatedBy(*s)
	}
	return suo
}

// SetUpdatedBy sets the "updated_by" field.
func (suo *SystemUpdateOne) SetUpdatedBy(s string) *SystemUpdateOne {
	suo.mutation.SetUpdatedBy(s)
	return suo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableUpdatedBy(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetUpdatedBy(*s)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SystemUpdateOne) SetUpdatedAt(t time.Time) *SystemUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetDeletedBy sets the "deleted_by" field.
func (suo *SystemUpdateOne) SetDeletedBy(s string) *SystemUpdateOne {
	suo.mutation.SetDeletedBy(s)
	return suo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableDeletedBy(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetDeletedBy(*s)
	}
	return suo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (suo *SystemUpdateOne) ClearDeletedBy() *SystemUpdateOne {
	suo.mutation.ClearDeletedBy()
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SystemUpdateOne) SetDeletedAt(t time.Time) *SystemUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableDeletedAt(t *time.Time) *SystemUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suo *SystemUpdateOne) ClearDeletedAt() *SystemUpdateOne {
	suo.mutation.ClearDeletedAt()
	return suo
}

// Mutation returns the SystemMutation object of the builder.
func (suo *SystemUpdateOne) Mutation() *SystemMutation {
	return suo.mutation
}

// Where appends a list predicates to the SystemUpdate builder.
func (suo *SystemUpdateOne) Where(ps ...predicate.System) *SystemUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SystemUpdateOne) Select(field string, fields ...string) *SystemUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated System entity.
func (suo *SystemUpdateOne) Save(ctx context.Context) (*System, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SystemUpdateOne) SaveX(ctx context.Context) *System {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SystemUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SystemUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SystemUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := system.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SystemUpdateOne) check() error {
	if v, ok := suo.mutation.GetType(); ok {
		if err := system.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "System.type": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Host(); ok {
		if err := system.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "System.host": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Credentials(); ok {
		if err := system.CredentialsValidator(v); err != nil {
			return &ValidationError{Name: "credentials", err: fmt.Errorf(`ent: validator failed for field "System.credentials": %w`, err)}
		}
	}
	if v, ok := suo.mutation.CreatedBy(); ok {
		if err := system.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "System.created_by": %w`, err)}
		}
	}
	if v, ok := suo.mutation.UpdatedBy(); ok {
		if err := system.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "System.updated_by": %w`, err)}
		}
	}
	return nil
}

func (suo *SystemUpdateOne) sqlSave(ctx context.Context) (_node *System, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(system.Table, system.Columns, sqlgraph.NewFieldSpec(system.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "System.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, system.FieldID)
		for _, f := range fields {
			if !system.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != system.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(system.FieldType, field.TypeString, value)
	}
	if value, ok := suo.mutation.Host(); ok {
		_spec.SetField(system.FieldHost, field.TypeString, value)
	}
	if value, ok := suo.mutation.UniqueIdentifier(); ok {
		_spec.SetField(system.FieldUniqueIdentifier, field.TypeString, value)
	}
	if value, ok := suo.mutation.Credentials(); ok {
		_spec.SetField(system.FieldCredentials, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(system.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := suo.mutation.UpdatedBy(); ok {
		_spec.SetField(system.FieldUpdatedBy, field.TypeString, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(system.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DeletedBy(); ok {
		_spec.SetField(system.FieldDeletedBy, field.TypeString, value)
	}
	if suo.mutation.DeletedByCleared() {
		_spec.ClearField(system.FieldDeletedBy, field.TypeString)
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.SetField(system.FieldDeletedAt, field.TypeTime, value)
	}
	if suo.mutation.DeletedAtCleared() {
		_spec.ClearField(system.FieldDeletedAt, field.TypeTime)
	}
	_node = &System{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
