package examples

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"github.com/google/go-cmp/cmp"
	"os"
	"reflect"
	"testing"
)

func TestGlobalConfigLoadFromEnvVars(t *testing.T) {
	envVars := map[string]string {
		"MyApp_IntegerValue": "21",
		"MyApp_FloatValue": "2.71",
		"MyApp_StringValue": "Hello There",
		"MyApp_LocalConfig_IntegerValue": "42",
		"MyApp_LocalConfig_FloatValue": "27.1",
		"MyApp_LocalConfig_StringValue": "Goodbye",
		"MyApp_LocalConfig_SliceValue": "[\"hello\", \"you\"]",
		"MyApp_LocalConfig_SliceOfSliceValue": "[[1, 2, 3], [3.14, 2.71]]",
		"MyApp_LocalConfig_MapValue": "{\"hello\": 3.14, \"you\": 2.71}",
		"MyApp_LocalConfig_ComplicatedValue": "[{\"hello\": [[\"s11\", \"s12\", \"s13\"], [\"s21\"]], \"you\": []}, {}, {\"are\": [[], [\"you\"], [], [\"ok?\"]]}]",
	}
	for envVar, value := range envVars {
		err := os.Setenv(envVar, value)
		if err != nil {
			t.Errorf("error while setting env var %s to %s: %s", envVar, value, err)
		}
	}

	g := GlobalConfig{}
	goconfman.LoadFromEnvVars(&g, "MyApp")

	expectedG := GlobalConfig{
		IntegerValue:       21,
		FloatValue:         2.71,
		StringValue:        "Hello There",
		LocalConfig:        LocalConfig{
			IntegerValue:      42,
			FloatValue:        27.1,
			StringValue:       "Goodbye",
			SliceValue:        []string{"hello", "you"},
			SliceOfSliceValue: [][]float32{{1, 2, 3}, {3.14, 2.71}},
			MapValue:          map[string]float32{"hello": 3.14, "you": 2.71},
			ComplicatedValue:  []map[string][][]string{{"hello": {{"s11", "s12", "s13"}, {"s21"}}, "you": {}}, {}, {"are": {{}, {"you"}, {}, {"ok?"}}}},
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
