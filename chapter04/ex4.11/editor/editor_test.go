package editor

import (
	"testing"
)

func TestRemoveUTF8BOM(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("我是大傻瓜"), []byte("我是大傻瓜")},
		{[]byte("\xef\xbb\xbf"), []byte("")},
		{[]byte("\xef\xbb\xbfA"), []byte("A")},
		{[]byte("\xef\xbb\xbf天下无贼"), []byte("天下无贼")},
	}

	for _, test := range tests {
		got := removeUTF8BOM(test.b)
		if string(got) != string(test.want) {
			t.Errorf("removeUTF8BOM(%q) = %q, want %q", string(test.b), string(got), string(test.want))
		}
	}
}

func TestGetEditorName(t *testing.T) {
	var tests = []struct {
		want string
	}{
		{"emace"},
		{"emacs"},
		{"vi"},
		{"vim"},
	}

	for _, test := range tests {
		got := getEditorName()
		if got != test.want {
			t.Errorf("getEditorName() = %q, want %q", got, test.want)
		}
	}
}
