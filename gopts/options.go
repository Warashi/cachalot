package gopts

import "github.com/Warashi/cachalot"

func Deserializer(d cachalot.Deserializer) cachalot.GetOption {
	return func(o *cachalot.GetOptions) {
		o.Deserializer = d
	}
}

func Custom(opts interface{}) cachalot.GetOption {
	return func(o *cachalot.GetOptions) {
		o.Custom = opts
	}
}
