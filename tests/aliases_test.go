package tests

import (
	"git.cafebazaar.ir/SayedSoroushHashemi/goconfman.git/pkg/goconfman"
	"testing"
)

func TestLoadFromAliases(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromDefaults(&g)
	goconfman.LoadFromAliases(&g)

	if g.LocalConfig.IntegerValue != g.IntegerValue {
		t.Errorf("g.LocalConfig.IntegerValue = %d != %d = g.IntegerValue", g.LocalConfig.IntegerValue, g.IntegerValue)
	}
	if g.LocalConfig.StringValue != g.StringValue {
		t.Errorf("g.LocalConfig.StringValue = %s != %s = g.StringValue", g.LocalConfig.StringValue, g.StringValue)
	}
}
