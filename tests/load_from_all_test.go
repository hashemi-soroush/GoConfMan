package tests

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"math"
	"testing"
)

func TestLoadFromAllGlobalConfig(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromAll(&g, "")

	if g.IntegerValue != 42 {
		t.Errorf("g.IntegerValue = %d != 42", g.IntegerValue)
	}
	if math.Abs(float64(g.FloatValue - 3.14)) > 1e-8 {
		t.Errorf("g.FloatValue = %f != 3.14", g.FloatValue)
	}
	if g.StringValue != "in global config" {
		t.Errorf("g.StringValue = %s != \"in global config\"", g.StringValue)
	}

	if g.LocalConfig.IntegerValue != 42 {
		t.Errorf("g.LocalConfig.IntegerValue = %d != 42", g.LocalConfig.IntegerValue)
	}
	if math.Abs(float64(g.LocalConfig.FloatValue - 31.4)) > 1e-8 {
		t.Errorf("g.LocalConfig.FloatValue = %f != 31.4", g.LocalConfig.FloatValue)
	}
	if g.LocalConfig.StringValue != "in global config" {
		t.Errorf("g.LocalConfig.StringValue = %s != \"in global config\"", g.LocalConfig.StringValue)
	}

	if g.NonGoConfManConfig.LocalConfig.IntegerValue != 420 {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.IntegerValue = %d != 420", g.NonGoConfManConfig.LocalConfig.IntegerValue)
	}
	if math.Abs(float64(g.NonGoConfManConfig.LocalConfig.FloatValue - 3.14)) > 1e-8 {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.FloatValue = %f != 3.14", g.NonGoConfManConfig.LocalConfig.FloatValue)
	}
	if g.NonGoConfManConfig.LocalConfig.StringValue != "in local config" {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.StringValue = %s != \"in local config\"", g.NonGoConfManConfig.LocalConfig.StringValue)
	}
}

func TestLoadFromAllNonGoConfManConfig(t *testing.T) {
	ng := NonGoConfManConfig{}
	goconfman.LoadFromAll(&ng, "")

	if ng.LocalConfig.IntegerValue != 420 {
		t.Errorf("ng.LocalConfig.IntegerValue = %d != 420", ng.LocalConfig.IntegerValue)
	}
	if math.Abs(float64(ng.LocalConfig.FloatValue - 31.4)) > 1e-8 {
		t.Errorf("ng.LocalConfig.FloatValue = %f != 31.4", ng.LocalConfig.FloatValue)
	}
	if ng.LocalConfig.StringValue != "in local config" {
		t.Errorf("ng.LocalConfig.StringValue = %s != \"in local config\"", ng.LocalConfig.StringValue)
	}
}
