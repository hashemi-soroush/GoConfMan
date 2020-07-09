package tests

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
