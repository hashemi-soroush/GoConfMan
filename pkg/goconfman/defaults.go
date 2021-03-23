package goconfman

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/internal"

type ConfigWithDefaults interface {
	BindDefaults()
}

// TODO: this method should be able to access unexported fields
func LoadFromDefaults(config interface{}) {
	c, ok := config.(ConfigWithDefaults)
	if ok {
		c.BindDefaults()
	}

	internal.LoadRecursive(config, LoadFromDefaults)
}
