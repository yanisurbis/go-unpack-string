package main

import "testing"

func TestUnpack(t *testing.T) {
	t.Run("unpacking multiple times", func(t *testing.T) {
		got, _ := Unpack("a4bc2d11e")
		want := "aaaabccddddddddddde"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("unpacking empty string", func(t *testing.T) {
		got, _ := Unpack("")
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("handles string without any packings correctly", func(t *testing.T) {
		got, _ := Unpack("abcd")
		want := "abcd"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("generates error when string is packed incorrectly", func(t *testing.T) {
		_, err := Unpack("3x44")

		if err == nil {
			t.Errorf("should generate error")
		}
	})
}
