package goconfman

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/internal"

type ConfigWithAliases interface {
	BindAliases()
}

func LoadFromAliases(config interface{}) {
	c, ok := config.(ConfigWithAliases)
	if ok {
		c.BindAliases()
	}

	internal.LoadRecursive(config, LoadFromAliases)
}
