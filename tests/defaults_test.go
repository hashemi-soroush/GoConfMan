package tests

import (
	"git.cafebazaar.ir/SayedSoroushHashemi/goconfman.git/pkg/goconfman"
	"testing"
)

func TestLoadFromDefaults(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromDefaults(&g)

	if g.IntegerValue != 42 {
		t.Errorf("g.IntegerValue = %d != 42", g.IntegerValue)
	}
	if g.FloatValue != 3.14 {
		t.Errorf("g.FloatValue = %f != 3.14", g.FloatValue)
	}
	if g.StringValue != "in global config" {
		t.Errorf("g.StringValue = %s != \"in global config\"", g.StringValue)
	}

	if g.LocalConfig.IntegerValue != 420 {
		t.Errorf("g.LocalConfig.IntegerValue = %d != 420", g.LocalConfig.IntegerValue)
	}
	if g.LocalConfig.FloatValue != 31.4 {
		t.Errorf("g.LocalConfig.FloatValue = %f != 31.4", g.LocalConfig.FloatValue)
	}
	if g.LocalConfig.StringValue != "in local config" {
		t.Errorf("g.LocalConfig.StringValue = %s != \"in local config\"", g.LocalConfig.StringValue)
	}
}
