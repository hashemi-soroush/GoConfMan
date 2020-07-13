package goconfman

type ConfigWithDefaults interface {
	BindDefaults()
}

// TODO: this method should be able to access unexported fields
func LoadFromDefaults(config interface{}) {
	loadRecursive(config, LoadFromDefaults)

	c, ok := config.(ConfigWithDefaults)
	if ok {
		c.BindDefaults()
	}
}
