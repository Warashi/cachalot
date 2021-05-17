package cachalot

import "context"

type (
	GetOptions struct{}
	SetOptions struct{}
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

// Local is used as local cache layer in MultiLayer
type Local interface {
	Get(ctx context.Context, key string, o *GetOptions) (value interface{}, found bool, err error)
	Set(ctx context.Context, key string, value interface{}, o *SetOptions) error
	Del(ctx context.Context, key string, o *DelOptions) error
}

// Remote is used as remote cache layer in MultiLayer
type Remote interface {
	Get(ctx context.Context, key string, o *GetOptions) (value []byte, found bool, err error)
	Set(ctx context.Context, key string, value []byte, o *SetOptions) error
	Del(ctx context.Context, key string, o *DelOptions) error
}

// MultiLayer is the interface you want
type MultiLayer interface {
	Get(ctx context.Context, key string, o ...GetOption) (value interface{}, from From, err error)
	Set(ctx context.Context, key string, o ...SetOption) error
	Del(ctx context.Context, key string, o ...DelOption) error
}