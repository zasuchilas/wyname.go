package main

import (
	"fmt"
	"testing"
)

// desamf

type desamfpairs struct {
	samf    int
	sex     int
	age     int
	sa      int
	filter  int
	filters []int
	mark    string
}

var desamftests = []desamfpairs{
	{sn | ac | fc, sn, ac, 0, 0, []int{0}, "6"},
	{sm, sm, an, 0, 0, []int{0}, "6"},
	{sf | sn | fc, sn, an, 0, 0, []int{0}, "6"},
	{sf | sm | ac | fc, sn, ac, 0, 0, []int{0}, "6"},
	{sf | ac | fc, sf, ac, fc, fc, []int{fc}, "0"},
	{sf | ac | fc | fd, sf, ac, fc, fc | fd, []int{fc, fd}, "0"},
}

func TestDesamf(t *testing.T) {
	fmt.Println(sn | ac | fc)
	fmt.Println(sm)
	fmt.Println(sf | sn | fc)
	for _, pair := range desamftests {
		s, a, sa, filter, filters, m := desamf(pair.samf)
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
		filtersChecks := len(pair.filters)
		for _, v := range filters {
			for _, ev := range pair.filters {
				if v == ev {
					filtersChecks--
					break
				}
			}
		}
		if filtersChecks != 0 {
			t.Error("For", pair.samf, "expected filters", pair.filters, "got", filters)
		}
		// t.Error("filters", filters)
		if m != pair.mark {
			t.Error("For", pair.samf, "expected mark", pair.mark, "got", m)
		}
	}
}

// chat

type chatpair struct {
	sa1       int
	f1        int
	sa2       int
	f2        int
	intersect bool
}

var chattests = []chatpair{
	{0, 0, 0, 0, true},
	{0, 0, sf | ac, fc | fd, false},
	{fc, md, fc, fc | fd, false},
	{fc, fc | md, fc, fc | fd, true},
	{mc, fc | fd, fc, fc | fd, false},
	{md, fc | fd, fc, mc | md, true},
	{fd, fc | md, fc, fc | fd, true},
}

func TestChat(t *testing.T) {
	for num, pair := range chattests {
		res := chat(pair.sa1, pair.f1, pair.sa2, pair.f2)
		if res != pair.intersect {
			t.Error("For pair num", num, "expected", pair.intersect, "got", res)
		}
	}
}

//
