package siv

import "testing"

func TestVersion(t *testing.T) {
	v := Version()
	if v != "v2.1.2" {
		t.Fatal("version mismatch, expected v2.1.2 got", v)
	}
}
