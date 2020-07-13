package tests

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
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
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		t.Errorf("Error while setting env var: \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s \n\t %s", err1, err2, err3, err4, err5, err6)
	}

	g := GlobalConfig{}
	goconfman.LoadFromEnvVars(&g, "MyApp")

	if g.IntegerValue != 21 {
		t.Errorf("")
	}
	if g.FloatValue != 2.71 {
		t.Errorf("")
	}
	if g.StringValue != "Hello There" {
		t.Errorf("")
	}

	if g.LocalConfig.IntegerValue != 42 {
		t.Errorf("")
	}
	if g.LocalConfig.FloatValue != 27.1 {
		t.Errorf("")
	}
	if g.LocalConfig.StringValue != "Goodbye" {
		t.Errorf("")
	}
}
