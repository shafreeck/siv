package siv

import "testing"

func TestVersion(t *testing.T) {
	v := Version()
	if v != "v0.1.1" {
		t.Fatal("version mismatch, expected v0.1.1 got", v)
	}
}
