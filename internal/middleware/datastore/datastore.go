package datastore

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"google.golang.org/grpc"

	"github.com/authzed/spicedb/internal/datastore"
)

type ctxKeyType struct{}

var datastoreKey ctxKeyType = struct{}{}

type datastoreHandle struct {
	datastore datastore.Datastore
}

// ContextWithHandle adds a placeholder to a context that will later be
// filled by the datastore
func ContextWithHandle(ctx context.Context) context.Context {
	return context.WithValue(ctx, datastoreKey, &datastoreHandle{})
}

// FromContext reads the selected datastore out of a context.Context
// and returns nil if it does not exist.
func FromContext(ctx context.Context) datastore.Datastore {
	if c := ctx.Value(datastoreKey); c != nil {
		handle := c.(*datastoreHandle)
		return handle.datastore
	}
	return nil
}

// MustFromContext reads the selected datastore out of a context.Context, computes a zedtoken
// from it, and panics if it has not been set on the context.
func MustFromContext(ctx context.Context) datastore.Datastore {
	datastore := FromContext(ctx)
	if datastore == nil {
		panic("datastore middleware did not inject datastore")
	}

	return datastore
}

// SetInContext adds a datastore to the given context
func SetInContext(ctx context.Context, datastore datastore.Datastore) error {
	handle := ctx.Value(datastoreKey)
	if handle == nil {
		return nil
	}
	handle.(*datastoreHandle).datastore = datastore
	return nil
}

// UnaryServerInterceptor returns a new unary server interceptor that adds the
// datastore to the context
func UnaryServerInterceptor(datastore datastore.Datastore) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := ContextWithHandle(ctx)
		if err := SetInContext(newCtx, datastore); err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

// StreamServerInterceptor returns a new stream server interceptor that adds the
// datastore to the context
func StreamServerInterceptor(datastore datastore.Datastore) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ContextWithHandle(wrapped.WrappedContext)
		if err := SetInContext(wrapped.WrappedContext, datastore); err != nil {
			return err
		}
		return handler(srv, wrapped)
	}
}
