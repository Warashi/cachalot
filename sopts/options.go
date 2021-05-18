package sopts

import (
	"time"

	"github.com/Warashi/cachalot"
)

func Serializer(s cachalot.Serializer) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.Serializer = s
	}
}

func TTL(ttl time.Duration) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.TTL = ttl
	}
}

func Custom(opts interface{}) cachalot.SetOption {
	return func(o *cachalot.SetOptions) {
		o.Custom = opts
	}
}
