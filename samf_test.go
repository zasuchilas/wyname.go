package main

import (
	"fmt"
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

// func TestSexfrom(t *testing.T) {
// 	for _, pair := range sexfromtests {
// 		v := sexfrom(pair.val)
// 		if v != pair.res {
// 			t.Error("For", pair.val, "expected", pair.res, "got", v)
// 		}
// 	}
// }

// desamf

type desamfpairs struct {
	samf   int
	sex    int
	age    int
	sa     int
	filter int
}

var desamftests = []desamfpairs{
	{sn | ac | fc, sn, ac, 0, 0},
	{sm, sm, an, 0, 0},
	{sf | sn | fc, sn, an, 0, 0},
	{sf | sm | ac | fc, sn, ac, 0, 0},
	{sf | ac | fc, sf, ac, sf | ac, fc},
	{sf | ac | fc | fd, sf, ac, sf | ac, fc | fd},
}

func TestDesamf(t *testing.T) {
	fmt.Println(sn | ac | fc)
	fmt.Println(sm)
	fmt.Println(sf | sn | fc)
	for _, pair := range desamftests {
		s, a, sa, filter := desamf(pair.samf)
		if s != pair.sex {
			t.Error("For", pair.samf, "expected sex", pair.sex, "got", s)
		}
		if a != pair.age {
			t.Error("For", pair.samf, "expected age", pair.age, "got", a)
		}
		if sa != pair.sa {
			t.Error("For", pair.samf, "expected sa", pair.sa, "got", sa)
		}
		if filter != pair.filter {
			t.Error("For", pair.samf, "expected filter", pair.filter, "got", filter)
		}
	}
}

//
