package main

import "testing"

var (
	ti = [2]float64{57.63389587402344, 39.83459854125977}
	ta = [2]float64{57.62788764527776, 39.82338690757751}
	tc = [2]float64{57.63990310499766, 39.845810174942024}
)

type megapair struct {
	value  float64
	result int
}

var megatests = []megapair{
	{57.63389587402344, 57633896},
	{39.83459854125977, 39834599},
	{0, 0},
	{-57.63389587402344, -57633896},
	{-39.83459854125977, -39834599},
	{9.834599432, 9834599},
	{1.834592453, 1834592}, // 1834593 ?
}

func TestMega(t *testing.T) {
	for _, pair := range megatests {
		v := mega(pair.value)
		if v != pair.result {
			t.Error("For", pair.value, "expected", pair.result, "got", v)
			// t.Errorf("results not match\n %v %v", r0, okti[0])
		}
	}
}

type micropair struct {
	value  int
	result float64
}

var microtests = []micropair{
	{57633896, 57.633896},
	{39834599, 39.834599},
	{0, 0},
	{-57633896, -57.633896},
	{-39834599, -39.834599},
	{9834599, 9.834599},
	{1834592, 1.834592},
}

func TestMicro(t *testing.T) {
	for _, pair := range microtests {
		v := micro(pair.value)
		if v != pair.result {
			t.Error("For", pair.value, "expected", pair.result, "got", v)
		}
	}
}

type partspair struct {
	value  float64
	result int
}

var latparttests = []partspair{
	{57.633895, 57635},
	{57.633395, 57635},
	{57.635395, 57640},
	{57.653895, 57655},
	{157.653895, 157655},
	{0, 0},
	{-57.633895, -57630},
	{-57.633395, -57630},
	{-57.635395, -57635},
	{-57.653895, -57650},
	{9.834599, 9835},
	{1.834592, 1835},
	{1.836592, 1840},
	{1.8, 1800},
	{1.79999, 1800},
	{1.79, 1790},
	{1.80001, 1805},
}

func TestSectorsLatPart(t *testing.T) {
	for _, pair := range latparttests {
		v, _ := sectorsLatPart(pair.value)
		if v != pair.result {
			t.Error("For", pair.value, "expected", pair.result, "got", v)
		}
	}
}

var lonparttests = []partspair{
	{57.633895, 5764},
	{57.633395, 5764},
	{57.625395, 5763},
	{57.653895, 5766},
}

func TestSectorsLonPart(t *testing.T) {
	for _, pair := range lonparttests {
		v, _ := sectorsLonPart(pair.value)
		if v != pair.result {
			t.Error("For", pair.value, "expected", pair.result, "got", v)
		}
	}
}

type sectorpair struct {
	lat float64
	lon float64
	sec string
}

var sectortests = []sectorpair{
	{57.63389587402344, 39.83459854125977, "57635|3984"},
	{57.62788764527776, 39.82338690757751, "57630|3983"},
	{57.63990310499766, 39.845810174942024, "57640|3985"},
}

func TestSector(t *testing.T) {
	for _, pair := range sectortests {
		v, _ := sector(pair.lat, pair.lon)
		if v != pair.sec {
			t.Error("For", pair.lat, pair.lon, "expected", pair.sec, "got", v)
		}
	}
}

func TestSectors(t *testing.T) {

}
