package hush

import "testing"

type TestSplit struct {
	in     string
	tokens []string
}

type TestHush struct {
	in   string
	want interface{}
}

func TestHushFile(t *testing.T) {
	hush := Hushfile()
	resultStr := "abcdefghijklmnopqrstuvwxyz"
	var resultInt int64 = 42
	var resultFloat float64 = 60.67

	key1 := "super_secret_key"
	key2 := "secret_app_number"
	key3 := "secret_float"
	key4 := "wrong_float"
	key5 := "wrong_int"
	key6 := "numeric_str"

	vals, _ := hush.GetString(key1)
	if vals != resultStr {
		t.Errorf("Got %s wanted %s", vals, resultStr)
	}

	vali, _ := hush.GetInt(key2)
	if vali != resultInt {
		t.Errorf("Got %d wanted %d", vali, resultInt)
	}

	valf, _ := hush.GetFloat(key3)
	if valf != resultFloat {
		t.Errorf("Got %f wanted %f", valf, resultFloat)
	}

	_, res := hush.GetFloat(key4)
	if res != false {
		t.Errorf("Got %t wanted %t", res, false)
	}

	_, res = hush.GetInt(key5)
	if res != false {
		t.Errorf("Got %t wanted %t", res, false)
	}

	_, res = hush.GetString(key6)
	if res != true {
		t.Errorf("Got %t wanted %t", res, true)
	}
}

func TestSplitLines(t *testing.T) {
	cases := []TestSplit{
		{"key1:value", []string{"key1", "value"}},
		{"key2:15", []string{"key2", "15"}},
		{" key3:abcd", []string{"key3", "abcd"}},
		{" key4: 20 ", []string{"key4", "20"}},
	}

	for _, c := range cases {
		got := processLine(c.in)
		if !compareSlices(got, c.tokens) {
			t.Errorf("processLine(%s) = %s, want %s", c.in, got, c.tokens)
		}
	}
}

func compareSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for index, val := range s1 {
			if val != s2[index] {
				return false
			}
		}
	}
	return true
}
