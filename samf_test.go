package main

import (
	"testing"
)

func TestPressed(t *testing.T) {
	t.Error(es)
}

func TestSexfrom(t *testing.T) {
	snsamf := sn | ac | fc
	snval := sexfrom(snsamf)
	if snval != sn {
		t.Error("sexfrom(", snsamf, ") returnd ", snval, "expect", sn)
	}
}
