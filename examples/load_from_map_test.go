package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"math"
	"testing"
)

func TestLoadFromMap(t *testing.T) {
	configMap := map[string]interface{} {
		"IntegerValue": 2,
		"FloatValue": -3.14,
		"LocalConfig": map[string]interface{} {
			"IntegerValue": 313,
		},
	}

	g := GlobalConfig{}
	err := goconfman.LoadFromMap(&g, configMap)
	if err != nil {
		t.Fatalf("Error in LoadFromMap: %s", err.Error())
	}

	if g.IntegerValue != 2 {
		t.Errorf("g.IntegerValue is loaded wrong: %v", g)
	}
	if math.Abs(float64(g.FloatValue - (-3.14))) > 1e-8 {
		t.Errorf("g.FloatValue is loaded wrong: %v", g)
	}
	if g.LocalConfig.IntegerValue != 313 {
		t.Errorf("g.LocalConfig.IntegerValue is loaded wrong: %v", g)
	}
}