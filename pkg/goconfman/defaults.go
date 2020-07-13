package goconfman

// TODO: this interface should be declared on Config not *Config. why???
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
