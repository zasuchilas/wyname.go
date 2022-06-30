package main

import "testing"

func TestNow(t *testing.T) {
	t.Error("1529048205774", "1529057006311727000", now())
}
