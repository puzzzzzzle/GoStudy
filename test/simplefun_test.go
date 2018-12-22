package test

import (
	"fmt"
	"testing"
)

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatal("true !=true")
	}
}

type AddTestCel struct {
	LValue int `json:"lValue"`
	RValue int `json:"rValue"`
	Want   int `json:"want"`
}
var testValues = []AddTestCel{
	{5,5,10},
	{12,15,27},
	{4,8,12},
}
func TestAdd(t *testing.T) {
	for _,v :=range testValues{
		got := MAdd(v.LValue, v.RValue)
		want := v.Want
		if got != want {
			t.Fatal(fmt.Sprintf("want %d got %d", want, got))
		}
	}

}
func BenchmarkAdd(t *testing.B) {
	for _,v :=range testValues{
		got := MAdd(v.LValue, v.RValue)
		want := v.Want
		if got != want {
			t.Fatal(fmt.Sprintf("want %d got %d", want, got))
		}
	}

}
