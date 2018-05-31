package main

import (
	"testing"
)

func TestGetsector(t *testing.T) {
	camp := newcamp()
	// [5763:3990 5763:3989 5763:3991 5764:3989 5764:3990 5764:3991]
	name := "5763:3990"
	name3 := "5763:3989"
	s := camp.sector(name)
	s2 := camp.sector(name)
	s3 := camp.sector(name3)
	if s != s2 {
		t.Error("sector name", name, "s", s, "s2", s2)
	}
	if s == s3 {
		t.Error("sector name", name, "s", s, "s3", s3)
	}

}
