package main

import (
	"reflect"
	"testing"
)

type testData struct {
	input interface{}
	want  interface{}
}

func TestGetAnagramsSet(t *testing.T) {
	tests := []testData{
		{&[]string{"Пятак", "пятка", "тЯпка", "Листок", "слиток", "столик", "тяпка", "листок", "собака"},
			&map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			}},
		{&[]string{"Пятак", "Листок", "собака"}, &map[string][]string{}},
		{&[]string{}, &map[string][]string{}},
	}
	for _, test := range tests {
		got := getAnagramsSet(test.input.(*[]string))
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("getAnagramsSet(%v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}
