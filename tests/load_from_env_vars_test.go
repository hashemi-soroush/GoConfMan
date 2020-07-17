package tests

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"math"
	"os"
	"testing"
)

func TestGlobalConfigLoadFromEnvVars(t *testing.T) {
	err1 := os.Setenv("MyApp__IntegerValue", "21")
	err2 := os.Setenv("MyApp__FloatValue", "2.71")
	err3 := os.Setenv("MyApp__StringValue", "Hello There")
	err4 := os.Setenv("MyApp__LocalConfig__IntegerValue", "42")
	err5 := os.Setenv("MyApp__LocalConfig__FloatValue", "27.1")
	err6 := os.Setenv("MyApp__LocalConfig__StringValue", "Goodbye")
	err7 := os.Setenv("MyApp__LocalConfig__SliceValue", "[\"hello\", \"you\"]")
	err8 := os.Setenv("MyApp__LocalConfig__SliceOfSliceValue", "[[1, 2, 3], [3.14, 2.71]]")
	err9 := os.Setenv("MyApp__LocalConfig__MapValue", "{\"hello\": 3.14, \"you\": 2.71}")
	err10 := os.Setenv("MyApp__LocalConfig__ComplicatedValue", "[{\"hello\": [[\"s11\", \"s12\", \"s13\"], [\"s21\"]], \"you\": []}, {}, {\"are\": [[], [\"you\"], [], [\"ok?\"]]}]")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil ||
		err8 != nil || err9 != nil || err10 != nil {
		t.Errorf(
			"Error while setting env var: \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s",
			err1, err2, err3, err4, err5, err6, err7, err8, err9, err10)
	}

	g := GlobalConfig{}
	goconfman.LoadFromEnvVars(&g, "MyApp")

	if g.IntegerValue != 21 {
		t.Errorf("g.IntegerValue is loaded wrong: %d", g.IntegerValue)
	}
	if g.FloatValue != 2.71 {
		t.Errorf("g.FloatValue is loaded wrong: %f", g.FloatValue)
	}
	if g.StringValue != "Hello There" {
		t.Errorf("g.StringValue is loaded wrong: %s", g.StringValue)
	}

	if g.LocalConfig.IntegerValue != 42 {
		t.Errorf("g.LocalConfig.IntegerValue is loaded wrong: %d", g.LocalConfig.IntegerValue)
	}
	if g.LocalConfig.FloatValue != 27.1 {
		t.Errorf("g.LocalConfig.FloatValue is loaded wrong: %f", g.LocalConfig.FloatValue)
	}
	if g.LocalConfig.StringValue != "Goodbye" {
		t.Errorf("g.LocalConfig.StringValue is loaded wrong: %s", g.LocalConfig.StringValue)
	}
	if g.LocalConfig.SliceValue[0] != "hello" || g.LocalConfig.SliceValue[1] != "you" {
		t.Errorf("g.SliceValue is loaded wrong: %v vs %v", []byte(g.LocalConfig.SliceValue[1][:]), []byte("you"))
	}

	if math.Abs(float64(g.LocalConfig.SliceOfSliceValue[0][0] - 1.0)) > 1e-8 ||
		math.Abs(float64(g.LocalConfig.SliceOfSliceValue[0][1] - 2.0)) > 1e-8  ||
		math.Abs(float64(g.LocalConfig.SliceOfSliceValue[0][2] - 3.0)) > 1e-8  ||
		math.Abs(float64(g.LocalConfig.SliceOfSliceValue[1][0] - 3.14)) > 1e-8 ||
		math.Abs(float64(g.LocalConfig.SliceOfSliceValue[1][1] - 2.71)) > 1e-8 {
		t.Errorf("g.LocalConfig.SliceOfSliceValue is loaded wrong: %v", g.LocalConfig.SliceOfSliceValue)
	}

	if  math.Abs(float64(g.LocalConfig.MapValue["hello"] - 3.14)) > 1e-8 ||
		math.Abs(float64(g.LocalConfig.MapValue["you"] - 2.71)) > 1e-8 {
		t.Errorf("g.LocalConfig.MapValue is loaded wrong: %v", g.LocalConfig.MapValue)
	}

	if g.LocalConfig.ComplicatedValue[0]["hello"][0][0] != "s11" ||
		g.LocalConfig.ComplicatedValue[0]["hello"][0][1] != "s12" ||
		g.LocalConfig.ComplicatedValue[0]["hello"][0][2] != "s13" ||
		g.LocalConfig.ComplicatedValue[0]["hello"][1][0] != "s21" ||
		len(g.LocalConfig.ComplicatedValue[0]["you"]) != 0 ||
		len(g.LocalConfig.ComplicatedValue[1]) != 0 ||
		len(g.LocalConfig.ComplicatedValue[2]["are"][0]) != 0 ||
		g.LocalConfig.ComplicatedValue[2]["are"][1][0] != "you" ||
		len(g.LocalConfig.ComplicatedValue[2]["are"][2]) != 0 ||
		g.LocalConfig.ComplicatedValue[2]["are"][3][0] != "ok?" {
		t.Errorf("g.LocalConfig.ComplicatedValue is loaded wrong %v", g.LocalConfig.ComplicatedValue)
	}

}
