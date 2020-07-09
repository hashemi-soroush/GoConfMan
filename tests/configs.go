package tests

type GlobalConfig struct {
	IntegerValue int
	FloatValue   float32
	StringValue  string
	LocalConfig  LocalConfig
}
func (g *GlobalConfig) BindDefaults() {
	g.IntegerValue = 42
	g.FloatValue = 3.14
	g.StringValue = "in global config"
}

type LocalConfig struct {
	IntegerValue int
	FloatValue   float32
	StringValue  string
}
func (l *LocalConfig) BindDefaults() {
	l.IntegerValue = 420
	l.FloatValue = 31.4
	l.StringValue = "in local config"
}
