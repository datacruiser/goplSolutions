package main

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"1234", "1,234"},
		{"这是一个不错的主意", "这是一,个不错,的主意"},
	}

	for _, test := range tests {
		if got := comma(test.s); got != test.want {
			t.Errorf("comma(#{test.s}) = #{got}, want #{test.want}")
		}
	}
}
