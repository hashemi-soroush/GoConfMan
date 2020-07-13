package tests

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"

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

func (l *LocalConfig) BindEnvVars(prefix string) {
	goconfman.BindEnvVar(&l.IntegerValue, "IntegerValue", prefix)
	goconfman.BindEnvVar(&l.FloatValue, "FloatValue", prefix)
	goconfman.BindEnvVar(&l.StringValue, "StringValue", prefix)
}