package main

import (
	"reflect"
	"testing"
)

type testData struct {
	input interface{}
	want  interface{}
}

func TestSortStringsWithNandK1(t *testing.T) {
	tests := []testData{
		{[][]string{{"3"}, {"2"}, {"1"}}, [][]string{{"1"}, {"2"}, {"3"}}},
		{[][]string{{"3a"}, {"2"}, {"1"}}, [][]string{{"3a"}, {"1"}, {"2"}}},
		{[][]string{{"3a"}, {"2"}, {"test"}, {"1"}}, [][]string{{"3a"}, {"test"}, {"1"}, {"2"}}},
		{[][]string{{"3a"}, {"2"}, {"1"}, {"aaa"}, {"bbb"}, {"Test"}}, [][]string{{"3a"}, {"Test"}, {"aaa"}, {"bbb"}, {"1"}, {"2"}}},
		{[][]string{{"3a"}, {"10", "30"}, {"2"}, {"1"}, {"12", "45"}, {"aaa", "ccc"}, {"bbb", "ggg"}, {"Test"}}, [][]string{{"3a"}, {"Test"}, {"aaa", "ccc"}, {"bbb", "ggg"}, {"1"}, {"2"}, {"10", "30"}, {"12", "45"}}},
	}
	for _, test := range tests {
		got := SortStrings(test.input.([][]string), 1, true, false)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("SortStrings(%v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}

func TestSortStringsWithNandK2(t *testing.T) {
	tests := []testData{
		{[][]string{{"3"}, {"2"}, {"1"}}, [][]string{{"1"}, {"2"}, {"3"}}},
		{[][]string{{"3a"}, {"2"}, {"1"}}, [][]string{{"1"}, {"2"}, {"3a"}}},
		{[][]string{{"3a"}, {"2"}, {"test"}, {"1"}}, [][]string{{"1"}, {"2"}, {"3a"}, {"test"}}},
		{[][]string{{"3a"}, {"2"}, {"1"}, {"aaa"}, {"bbb"}, {"Test"}}, [][]string{{"1"}, {"2"}, {"3a"}, {"Test"}, {"aaa"}, {"bbb"}}},
		{[][]string{{"3a"}, {"10", "30"}, {"2"}, {"1"}, {"12", "45"}, {"aaa", "ccc"}, {"bbb", "ggg"}, {"Test"}}, [][]string{{"1"}, {"2"}, {"3a"}, {"Test"}, {"aaa", "ccc"}, {"bbb", "ggg"}, {"10", "30"}, {"12", "45"}}},
	}
	for _, test := range tests {
		got := SortStrings(test.input.([][]string), 2, true, false)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("SortStrings(%v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}
