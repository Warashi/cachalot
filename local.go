package cachalot

import (
	"context"
	"fmt"

	"github.com/dgraph-io/ristretto"
)

var _ Local = (*Ristretto)(nil)

type Ristretto struct {
	base *ristretto.Cache
}

func NewRistretto(config *ristretto.Config) (*Ristretto, error) {
	c, err := ristretto.NewCache(config)
	if err != nil {
		return nil, fmt.Errorf("creating ristretto.Cache: %w", err)
	}
	return &Ristretto{
		base: c,
	}, nil
}

func (c *Ristretto) Get(ctx context.Context, key string, o *LocalGetOptions) (value interface{}, found bool, err error) {
	value, found = c.base.Get(key)
	return value, found, nil
}

func (c *Ristretto) Set(ctx context.Context, key string, value interface{}, o *LocalSetOptions) error {
	if !c.base.SetWithTTL(key, value, o.Cost, o.TTL) {
		return ErrSetFailed
	}
	return nil
}

func (c *Ristretto) Del(ctx context.Context, key string, o *LocalDelOptions) error {
	c.base.Del(key)
	return nil
}
