package tests

type GlobalConfig struct {
	IntegerValue int
	FloatValue   float32
	StringValue  string
	LocalConfig  LocalConfig
	NonGoConfManConfig NonGoConfManConfig
}

func (g *GlobalConfig) BindDefaults() {
	g.IntegerValue = 42
	g.FloatValue = 3.14
	g.StringValue = "in global config"
}

func (g *GlobalConfig) BindAliases() {
	g.LocalConfig.IntegerValue = g.IntegerValue
	g.LocalConfig.StringValue = g.StringValue
	g.NonGoConfManConfig.LocalConfig.FloatValue = g.FloatValue
}
