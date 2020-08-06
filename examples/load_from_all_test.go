package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"github.com/google/go-cmp/cmp"
	"os"
	"reflect"
	"testing"
)

func TestLoadFromAllGlobalConfig(t *testing.T) {
	envVars := map[string]string {
		"MyApp_StringValue": "StringValue in envVars",
		"MyApp_LocalConfig_IntegerValue": "721",
		"MyApp_LocalConfig_SliceValue": "[\"hello\", \"you\"]",
		"MyApp_LocalConfig_ComplicatedValue": "[{\"hello\": [[\"s11\", \"s12\", \"s13\"], [\"s21\"]], \"you\": []}, {}, {\"are\": [[], [\"you\"], [], [\"ok?\"]]}]",
	}

	configMap := map[string]interface{}{
		"LocalConfig": map[string]interface{}{
			"StringValue": "Goodbye",
			"MapValue": map[string]float32{"hello": 3.14, "you": 2.71},
		},
		"NonGoConfManConfig": map[string]interface{}{
			"LocalConfig": map[string]interface{}{
				"IntegerValue": 123,
			},
		},
	}

	for envVar, value := range envVars {
		err := os.Setenv(envVar, value)
		if err != nil {
			t.Errorf("error while setting env var %s to %s: %s", envVar, value, err)
		}
	}

	g := GlobalConfig{}
	goconfman.LoadFromAll(&g, configMap, "global_config.json", "MyApp")

	expectedG := GlobalConfig{
		IntegerValue:       1234,
		FloatValue:         987.3,
		StringValue:        "StringValue in envVars",
		LocalConfig:        LocalConfig{
			IntegerValue:      1234,
			FloatValue:        31.4,
			StringValue:       "StringValue in envVars",
			SliceValue:        []string{"hello", "you"},
			SliceOfSliceValue: nil,
			MapValue:          map[string]float32{"hello": 3.14, "you": 2.71},
			ComplicatedValue:  []map[string][][]string{{"hello": {{"s11", "s12", "s13"}, {"s21"}}, "you": {}}, {}, {"are": {{}, {"you"}, {}, {"ok?"}}}},
		},
		NonGoConfManConfig: NonGoConfManConfig{
			IntegerValue: 0,
			FloatValue:   0,
			StringValue:  "",
			LocalConfig:  LocalConfig{
				IntegerValue:      71,
				FloatValue:        987.3,
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

//func TestLoadFromAllNonGoConfManConfig(t *testing.T) {
//	ng := NonGoConfManConfig{}
//	goconfman.LoadFromAll(&ng, nil, "", "")
//
//	expectedNG := NonGoConfManConfig{
//		IntegerValue: 0,
//		FloatValue:   0,
//		StringValue:  "",
//		LocalConfig:  LocalConfig{
//			IntegerValue:      420,
//			FloatValue:        31.4,
//			StringValue:       "in local config",
//			SliceValue:        nil,
//			SliceOfSliceValue: nil,
//			MapValue:          nil,
//			ComplicatedValue:  nil,
//		},
//	}
//
//	if reflect.DeepEqual(ng, expectedNG) == false {
//		t.Errorf("ng is different from the expectedNG. here's the diff: \n%s", cmp.Diff(ng, expectedNG))
//	}
//}
