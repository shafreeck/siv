package siv

import "testing"

func TestVersion(t *testing.T) {
	v := Version()
	if v != "v1.1.1" {
		t.Fatal("version mismatch, expected v1.1.1 got", v)
	}
}
