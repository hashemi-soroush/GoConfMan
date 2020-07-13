package goconfman

type ConfigWithAliases interface {
	BindAliases()
}

func LoadFromAliases(config interface{}) {
	loadRecursive(config, LoadFromAliases)

	c, ok := config.(ConfigWithAliases)
	if ok {
		c.BindAliases()
	}
}
