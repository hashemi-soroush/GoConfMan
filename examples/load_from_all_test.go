package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestLoadFromAllGlobalConfig(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromAll(&g, "", "")

	expectedG := GlobalConfig{
		IntegerValue:       42,
		FloatValue:         3.14,
		StringValue:        "in global config",
		LocalConfig:        LocalConfig{
			IntegerValue:      42,
			FloatValue:        31.4,
			StringValue:       "in global config",
			SliceValue:        nil,
			SliceOfSliceValue: nil,
			MapValue:          nil,
			ComplicatedValue:  nil,
		},
		NonGoConfManConfig: NonGoConfManConfig{
			IntegerValue: 0,
			FloatValue:   0,
			StringValue:  "",
			LocalConfig:  LocalConfig{
				IntegerValue:      420,
				FloatValue:        3.14,
				StringValue:       "in local config",
				SliceValue:        nil,
				SliceOfSliceValue: nil,
				MapValue:          nil,
				ComplicatedValue:  nil,
			},
		},
	}

	if reflect.DeepEqual(g, expectedG) == false {
		t.Errorf("g is different from the expectedG. here's the diff: \n%s", cmp.Diff(g, expectedG))
	}
}

func TestLoadFromAllNonGoConfManConfig(t *testing.T) {
	ng := NonGoConfManConfig{}
	goconfman.LoadFromAll(&ng, "", "")

	expectedNG := NonGoConfManConfig{
		IntegerValue: 0,
		FloatValue:   0,
		StringValue:  "",
		LocalConfig:  LocalConfig{
			IntegerValue:      420,
			FloatValue:        31.4,
			StringValue:       "in local config",
			SliceValue:        nil,
			SliceOfSliceValue: nil,
			MapValue:          nil,
			ComplicatedValue:  nil,
		},
	}

	if reflect.DeepEqual(ng, expectedNG) == false {
		t.Errorf("ng is different from the expectedNG. here's the diff: \n%s", cmp.Diff(ng, expectedNG))
	}
}
