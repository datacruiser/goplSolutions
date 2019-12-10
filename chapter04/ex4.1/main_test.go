package main

import (
	"testing"
)

func TestSha256PopCount(t *testing.T) {
	var tests = []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"学习go语言", "学习go语言", 0},
		{"a", "b", 126},
		{"测试一下", "再测试一下", 124},
	}

	for _, test := range tests {
		got := sha256PopCount(test.a, test.b)
		if got != test.want {
			t.Errorf("sha256PopCount(%s, %s) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
