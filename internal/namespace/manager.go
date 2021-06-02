package namespace

import (
	"context"

	pb "github.com/authzed/spicedb/pkg/REDACTEDapi/api"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("spicedb/internal/namespace")

// Manager is a subset of the datastore interface that can read (and possibly cache) namespaces.
type Manager interface {
	// ReadNamespace reads a namespace definition and version and returns it if found.
	ReadNamespace(ctx context.Context, nsName string) (*pb.NamespaceDefinition, decimal.Decimal, error)

	// CheckNamespaceAndRelation checks that the specified namespace and relation exist in the
	// datastore.
	CheckNamespaceAndRelation(ctx context.Context, namespace, relation string, allowEllipsis bool) error

	// ReadNamespaceAndTypes reads a namespace definition, version, and type system and returns it if found.
	ReadNamespaceAndTypes(ctx context.Context, nsName string) (*pb.NamespaceDefinition, *NamespaceTypeSystem, decimal.Decimal, error)
}