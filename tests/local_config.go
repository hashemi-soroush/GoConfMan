package tests

import "github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"

type LocalConfig struct {
	IntegerValue int
	FloatValue   float32
	StringValue  string
	SliceValue []string
	SliceOfSliceValue [][]float32
	MapValue map[string]float32
	ComplicatedValue []map[string][][]string
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
	goconfman.BindEnvVar(&l.SliceValue, "SliceValue", prefix)
	goconfman.BindEnvVar(&l.SliceOfSliceValue, "SliceOfSliceValue", prefix)
	goconfman.BindEnvVar(&l.MapValue, "MapValue", prefix)
	goconfman.BindEnvVar(&l.ComplicatedValue, "ComplicatedValue", prefix)
}