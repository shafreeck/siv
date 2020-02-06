package siv

import "testing"

func TestVersion(t *testing.T) {
	v := Version()
	if v != "v3.0.0" {
		t.Fatal("version mismatch, expected v3.0.0 got", v)
	}
}
