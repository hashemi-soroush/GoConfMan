package tests

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"

type GlobalConfig struct {
	IntegerValue       int
	FloatValue         float32
	StringValue        string
	LocalConfig        LocalConfig
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

func (g *GlobalConfig) BindEnvVars(prefix string) {
	goconfman.BindEnvVar(&g.IntegerValue, "IntegerValue", prefix)
	goconfman.BindEnvVar(&g.FloatValue, "FloatValue", prefix)
	goconfman.BindEnvVar(&g.StringValue, "StringValue", prefix)
}
