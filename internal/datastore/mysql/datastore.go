package mysql

import (
	"context"
	"errors"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

	"github.com/authzed/spicedb/internal/datastore"
	"github.com/authzed/spicedb/internal/datastore/options"
)

var errUnimplemented = errors.New("unimplemented")

func NewMysqlDatastore() (datastore.Datastore, error) {
	return &mysqlDatastore{}, nil
}

type mysqlDatastore struct{}

// QueryTuples reads relationships starting from the resource side.
func (mds *mysqlDatastore) QueryTuples(
	ctx context.Context,
	filter *v1.RelationshipFilter,
	revision datastore.Revision,
	options ...options.QueryOptionsOption,
) (datastore.TupleIterator, error) {
	return nil, errUnimplemented
}

// ReverseQueryRelationships reads relationships starting from the subject.
func (mds *mysqlDatastore) ReverseQueryTuples(
	ctx context.Context,
	subjectFilter *v1.SubjectFilter,
	revision datastore.Revision,
	options ...options.ReverseQueryOptionsOption,
) (datastore.TupleIterator, error) {
	return nil, errUnimplemented
}

// CheckRevision checks the specified revision to make sure it's valid and
// hasn't been garbage collected.
func (mds *mysqlDatastore) CheckRevision(ctx context.Context, revision datastore.Revision) error {
	return errUnimplemented
}

// WriteTuples takes a list of existing tuples that must exist, and a list of
// tuple mutations and applies it to the datastore for the specified
// namespace.
func (mds *mysqlDatastore) WriteTuples(ctx context.Context, preconditions []*v1.Precondition, mutations []*v1.RelationshipUpdate) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// DeleteRelationships deletes all Relationships that match the provided
// filter if all preconditions are met.
func (mds *mysqlDatastore) DeleteRelationships(ctx context.Context, preconditions []*v1.Precondition, filter *v1.RelationshipFilter) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// OptimizedRevision gets a revision that will likely already be replicated
// and will likely be shared amongst many queries.
func (mds *mysqlDatastore) OptimizedRevision(ctx context.Context) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// HeadRevision gets a revision that is guaranteed to be at least as fresh as
// right now.
func (mds *mysqlDatastore) HeadRevision(ctx context.Context) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// Watch notifies the caller about all changes to tuples.
//
// All events following afterRevision will be sent to the caller.
func (mds *mysqlDatastore) Watch(ctx context.Context, afterRevision datastore.Revision) (<-chan *datastore.RevisionChanges, <-chan error) {
	errs := make(chan error)
	errs <- errUnimplemented
	return nil, errs
}

// WriteNamespace takes a proto namespace definition and persists it,
// returning the version of the namespace that was created.
func (mds *mysqlDatastore) WriteNamespace(ctx context.Context, newConfig *v0.NamespaceDefinition) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// ReadNamespace reads a namespace definition and version and returns it, and the revision at
// which it was created or last written, if found.
func (mds *mysqlDatastore) ReadNamespace(ctx context.Context, nsName string, revision datastore.Revision) (ns *v0.NamespaceDefinition, lastWritten datastore.Revision, err error) {
	return nil, datastore.NoRevision, errUnimplemented
}

// DeleteNamespace deletes a namespace and any associated tuples.
func (mds *mysqlDatastore) DeleteNamespace(ctx context.Context, nsName string) (datastore.Revision, error) {
	return datastore.NoRevision, errUnimplemented
}

// ListNamespaces lists all namespaces defined.
func (mds *mysqlDatastore) ListNamespaces(ctx context.Context, revision datastore.Revision) ([]*v0.NamespaceDefinition, error) {
	return nil, errUnimplemented
}

// IsReady returns whether the datastore is ready to accept data. Datastores that require
// database schema creation will return false until the migrations have been run to create
// the necessary tables.
func (mds *mysqlDatastore) IsReady(ctx context.Context) (bool, error) {
	return false, errUnimplemented
}

// Close closes the data store.
func (mds *mysqlDatastore) Close() error {
	return errUnimplemented
}