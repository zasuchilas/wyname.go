package main

import (
	"testing"
)

func TestPressed(t *testing.T) {
	t.Error(es)
}

// sexfrom(samf)

type sexfrompairs struct {
	val int
	res int
}

var sexfromtests = []sexfrompairs{
	{sn | ac | fc, sn},
	{sm, sm},
	{sf | sn | fc, sn}, // если несколько значений в samf, вернуться должно sn
}

func TestSexfrom(t *testing.T) {
	for _, pair := range sexfromtests {
		v := sexfrom(pair.val)
		if v != pair.res {
			t.Error("For", pair.val, "expected", pair.res, "got", v)
		}
	}
}

// ----------------
