package main

import (
	"bytes"
	"testing"
)

func TestReverseUTF8(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("ABC"), []byte("CBA")},
		{[]byte("这个男人来自地球"), []byte("球地自来人男个这")},
	}

	for _, test := range tests {
		got := reverseUTF8(test.b)
		if string(got) != string(test.want) {
			t.Errorf("reverseUTF8(%q) = %q, want %q", string(test.b), string(got), string(test.want))
		}
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("ABC"), []byte("CBA")},
		{[]byte("这个男人来自地球"), []byte("球地自来人男个这")},
	}

	for _, test := range tests {
		if reverse(test.b); !bytes.Equal(test.b, test.want) {
			t.Errorf("got %s, want %s", test.b, test.want)
		}
	}
}
