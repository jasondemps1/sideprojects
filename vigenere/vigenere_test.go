package vigenere

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	cases := []struct {
		plainText, keyword, want string
	}{
		{"attack at dawn", "lemon", "LXFOPVEFRNHR"},
	}

	for _, c := range cases {
		got := Encrypt(c.plainText, c.keyword)

		if got != c.want {
			t.Errorf("TestEncrypt(%q, %q) == %q, want %q", c.plainText, c.keyword, got, c.want)
		}
	}
}

func TestDecrypt(t *testing.T) {
	cases := []struct {
		cipherText, keyword, want string
	}{
		{"LXFOPVEFRNHR", "lemon", "ATTACKATDAWN"},
	}

	for _, c := range cases {
		got := Decrypt(c.cipherText, c.keyword)

		if got != c.want {
			t.Errorf("TestDecrypt(%q, %q) == %q, want %q", c.cipherText, c.keyword, got, c.want)
		}
	}
}
