package goconfman

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/internal"

type ConfigWithAliases interface {
	BindAliases()
}

func LoadFromAliases(config interface{}) {
	internal.LoadRecursive(config, LoadFromAliases)

	c, ok := config.(ConfigWithAliases)
	if ok {
		c.BindAliases()
	}
}
