package main

import (
	"testing"
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

var (
	ti = [2]float64{57.63389587402344, 39.83459854125977}
	ta = [2]float64{57.62788764527776, 39.82338690757751}
	tc = [2]float64{57.63990310499766, 39.845810174942024}
)

type sectorspair struct {
	mla int
	mlo int
	// secs map[int]string
}

var sectorstests = []sectorspair{
	{57633896, 39834599},
}

// var megatests = []megapair{
// 	{57.63389587402344, 57633896},
// 	{39.83459854125977, 39834599},
// 	{0, 0},
// 	{-57.63389587402344, -57633896},
// 	{-39.83459854125977, -39834599},
// 	{9.834599432, 9834599},
// 	{1.834592453, 1834592}, // 1834593 ?
// }

// map[2:57640|3983 6:57640|3984 3:57640|3985 7:57635|3985 4:57630|3985 1:57630|3983 8:57630|3984 5:57635|3983]

func TestSectors(t *testing.T) {
	secs, _ := sectors(57633896, 39834599)
	t.Error(secs)
	// fmt.Println("secs")
}

func TestLatab(t *testing.T) {
	lt := latable()
	t.Error(lt[0], lt[1], lt[2], lt[89], lt[91])
}
func TestDistable(t *testing.T) {
	dt := distable()
	for i := 0; i < 90; i++ {
		ipar := float64(i)
		t.Error(ipar, dt[ipar])
	}
}

func TestDist589(t *testing.T) {
	g, _ := newGps(57.633895, 39.834598)
	_, d589 := g.dist589()
	dt := distable()
	t.Error(g, d589, dt[57], dt[58])
}
