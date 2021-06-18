package sopts

import (
	"time"

	"github.com/Warashi/cachalot"
)

func Serializer(s cachalot.Serializer) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.Remote.Serializer = s
	}
}

func TTL(ttl time.Duration) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.TTL = ttl
	}
}

func Cost(c int) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.Local.Cost = c
	}
}
