package main

import "testing"

type testData struct {
	input string
	want  string
}

func TestUnpackValidStr(t *testing.T) {
	tests := []testData{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
		{`\45`, `44444`},
		{`\5`, `5`},
		{`\\`, `\`},
		{`\\3`, `\\\`},
		{"", ""},
		{"a", "a"},
		{"abcd3", "abcddd"},
		{"ab\ncd3", "ab\ncddd"},
		{"ab\tcd3", "ab\tcddd"},
	}
	for _, test := range tests {
		got, err := Unpack(test.input)
		if err != nil {
			t.Errorf("Unpack returned error: %s, want \"%s\"", err, test.want)
		}
		if got != test.want {
			t.Errorf("Unpack(%s) = \"%s\", want \"%s\"", test.input, test.want, got)
		}
	}
}

func TestUnpackInvalidStr(t *testing.T) {
	tests := []testData{
		{"45", ""},
		{"\\", ""},
		{`ab\`, ""},
		{`5`, ""},
		{`\q`, ``},
		{`a\q`, ``},
	}
	for _, test := range tests {
		got, err := Unpack(test.input)
		if err != ErrInvalidString {
			t.Errorf("Unpack must retur error: \"некорректная строка\"")
		}
		if got != test.want {
			t.Errorf("Unpack(%s) = \"%s\", want \"%s\"", test.input, test.want, got)
		}
	}
}
