package tests

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"testing"
)

func TestLoadFromAll(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromAll(&g)

	if g.IntegerValue != 42 {
		t.Errorf("g.IntegerValue = %d != 42", g.IntegerValue)
	}
	if g.FloatValue != 3.14 {
		t.Errorf("g.FloatValue = %f != 3.14", g.FloatValue)
	}
	if g.StringValue != "in global config" {
		t.Errorf("g.StringValue = %s != \"in global config\"", g.StringValue)
	}

	if g.LocalConfig.IntegerValue != 42 {
		t.Errorf("g.LocalConfig.IntegerValue = %d != 42", g.LocalConfig.IntegerValue)
	}
	if g.LocalConfig.FloatValue != 31.4 {
		t.Errorf("g.LocalConfig.FloatValue = %f != 31.4", g.LocalConfig.FloatValue)
	}
	if g.LocalConfig.StringValue != "in global config" {
		t.Errorf("g.LocalConfig.StringValue = %s != \"in global config\"", g.LocalConfig.StringValue)
	}
}
