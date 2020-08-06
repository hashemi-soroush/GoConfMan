package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"github.com/google/go-cmp/cmp"
	"reflect"
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

	expectedG := GlobalConfig{
		IntegerValue:       2,
		FloatValue:         -3.14,
		StringValue:        "",
		LocalConfig:        LocalConfig{
			IntegerValue:      313,
			FloatValue:        0,
			StringValue:       "",
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
				IntegerValue:      0,
				FloatValue:        0,
				StringValue:       "",
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