package cachalot

import (
	"context"
	"time"
)

type (
	LocalGetOptions  struct{}
	RemoteGetOptions struct {
		Deserializer Deserializer
	}
	GetOptions struct {
		Local  LocalGetOptions
		Remote RemoteGetOptions
	}

	LocalSetOptions struct {
		Cost int
		TTL  time.Duration
	}
	RemoteSetOptions struct {
		Serializer Serializer
		TTL        time.Duration
	}
	SetOptions struct {
		Local  LocalSetOptions
		Remote RemoteSetOptions
	}
	DelOptions struct{}
)

type (
	GetOption func(*GetOptions)
	SetOption func(*SetOptions)
	DelOption func(*DelOptions)
)

type From int

const (
	NotFound From = iota
	LocalCache
	RemoteCache
)

type (
	Deserializer interface {
		Deserialize([]byte) (interface{}, error)
	}
	Serializer interface {
		Serialize(interface{}) ([]byte, error)
	}
)

// Local is used as local cache layer in MultiLayer
type Local interface {
	Get(ctx context.Context, key string, o *GetOptions) (value interface{}, found bool, err error)
	Set(ctx context.Context, key string, value interface{}, o *LocalSetOptions) error
	Del(ctx context.Context, key string, o *DelOptions) error
}

// Remote is used as remote cache layer in MultiLayer
type Remote interface {
	Get(ctx context.Context, key string, o *GetOptions) (value []byte, found bool, err error)
	Set(ctx context.Context, key string, value interface{}, o *RemoteSetOptions) error
	Del(ctx context.Context, key string, o *DelOptions) error

	// InvalidatedKeys returns channel which sends key of invalidated cache
	InvalidatedKeys() <-chan []string
}

// MultiLevel is the interface you want
type MultiLevel interface {
	Get(ctx context.Context, key string, o ...GetOption) (value interface{}, from From, err error)
	Set(ctx context.Context, key string, o ...SetOption) error
	Del(ctx context.Context, key string, o ...DelOption) error
}
