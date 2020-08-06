package internal

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
	"testing"
)

func TestSetBool(t *testing.T) {
	var myVar bool
	myVarValue := reflect.ValueOf(&myVar)

	var err interface{}

	myVar = true
	setBool(myVarValue, "false")
	err = recover()
	if err != nil {
		t.Errorf("error in setBool: %s", err)
	}
	if myVar != false {
		t.Errorf("myVar is %t", myVar)
	}
}

func TestSetInt(t *testing.T) {
	var myVar int
	myVarValue := reflect.ValueOf(&myVar)

	var err interface{}

	myVar = 0
	setInt(myVarValue, "27")
	err = recover()
	if err != nil {
		t.Errorf("error in setInt: %s", err)
	}
	if myVar != 27 {
		t.Errorf("myVar is %d", myVar)
	}
}

func TestSetUint(t *testing.T) {
	var myVar uint
	myVarValue := reflect.ValueOf(&myVar)

	var err interface{}

	myVar = 0
	setUint(myVarValue, "13")
	err = recover()
	if err != nil {
		t.Errorf("error in setUint: %s", err)
	}
	if myVar != 13 {
		t.Errorf("myVar is %d", myVar)
	}
}

func TestSetFloat(t *testing.T) {
	var myVar float64
	myVarValue := reflect.ValueOf(&myVar)

	var err interface{}

	myVar = 0.0
	setFloat(myVarValue, "3.14")
	err = recover()
	if err != nil {
		t.Errorf("error in setFloat: %s", err)
	}
	if math.Abs(myVar - 3.14) > 1e-8 {
		t.Errorf("myVar is %f", myVar)
	}
}

func TestSetString(t *testing.T) {
	var myVar string
	myVarValue := reflect.ValueOf(&myVar)

	var err interface{}

	myVar = ""
	setString(myVarValue, "hello")
	err = recover()
	if err != nil {
		t.Errorf("error in setString: %s", err)
	}
	if myVar != "hello" {
		t.Errorf("myVar is %s", myVar)
	}
}

func TestSetSlice(t *testing.T) {
	var err interface{}

	var sliceOfInt []int
	sliceOfIntValue := reflect.ValueOf(&sliceOfInt)
	sliceOfIntExpectedValue := []int{23, 17}
	setSlice(sliceOfIntValue, "[23, 17]")
	err = recover()
	if err != nil {
		t.Errorf("error in setSlice: %s", err)
	}
	if reflect.DeepEqual(sliceOfInt, sliceOfIntExpectedValue) == false {
		t.Errorf("difference between sliceOfInt and its expected value: %s", cmp.Diff(sliceOfInt, sliceOfIntExpectedValue))
	}

	var sliceOfString []string
	sliceOfStringValue := reflect.ValueOf(&sliceOfString)
	sliceOfStringExpectedValue := []string{"hello", "there"}
	setSlice(sliceOfStringValue, "[\"hello\", \"there\"]")
	err = recover()
	if err != nil {
		t.Errorf("error in setSlice: %s", err)
	}
	if reflect.DeepEqual(sliceOfString, sliceOfStringExpectedValue) == false {
		t.Errorf("difference between sliceOfString and its expected value: %s", cmp.Diff(sliceOfString, sliceOfStringExpectedValue))
	}

	var sliceOfSliceOfFloat [][]float64
	sliceOfSliceOfFloatValue := reflect.ValueOf(&sliceOfSliceOfFloat)
	sliceOfSliceOfFloatExpectedValue := [][]float64{{3.14, 2.71}, {9.8}}
	setSlice(sliceOfSliceOfFloatValue, "[[3.14, 2.71], [9.8]]")
	err = recover()
	if err != nil {
		t.Errorf("error in setSlice: %s", err)
	}
	if reflect.DeepEqual(sliceOfSliceOfFloat, sliceOfSliceOfFloatExpectedValue) == false {
		t.Errorf("difference between sliceOfSliceOfFloat and its expected value: %s", cmp.Diff(sliceOfSliceOfFloat, sliceOfSliceOfFloatExpectedValue))
	}
}

func TestSetMap(t *testing.T) {
	var err interface{}

	var mapOfStringToInt map[string]int
	mapOfStringToIntValue := reflect.ValueOf(&mapOfStringToInt)
	mapOfStringToIntExpectedValue := map[string]int{"hello": 42, "you": 13}
	setMap(mapOfStringToIntValue, "{\"hello\": 42, \"you\": 13}")
	err = recover()
	if err != nil {
		t.Errorf("error in setMap: %s", err)
	}
	if reflect.DeepEqual(mapOfStringToInt, mapOfStringToIntExpectedValue) == false {
		t.Errorf("difference between mapOfStringToInt and its expected value: %s", cmp.Diff(mapOfStringToInt, mapOfStringToIntExpectedValue))
	}
}