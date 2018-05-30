package main

import (
	"math"
	"testing"
)

func TestDist589(t *testing.T) {
	points := [][2]float64{
		{57.633895, 39.834598},
		{-57.633895, 39.834598},
		{157.633895, 39.834598},
		{-157.633895, 39.834598},
		{1.633895, 39.834598},
		{0, 39.834598},
	}

	dt := distable()
	for _, point := range points {
		g, _ := newGps(point[0], point[1])
		_, d589 := g.dist589()
		p := math.Floor(g.lat)
		if d589 < dt[p] || d589 > dt[p+1] {
			t.Error(g, d589, dt[p], dt[p+1])
		}
	}
}

// Вычисляет таблицу для всех параллелей
// насколько градусов нужно сместиться на восток по параллели
// чтобы преодолеть 589 метров
// (for develop only)
func distable() map[float64]float64 {
	ditab := make(map[float64]float64, 90)
	equator := 0.0052910052910053
	// var dig float64 = 100000
	for parallel := 0; parallel < 900; parallel++ {
		ipar := float64(parallel)
		ditab[ipar] = equator / math.Cos(ipar*(math.Pi/180))
		// ditab[ipar] = math.Round((equator/math.Cos(ipar*(math.Pi/180)))*dig) / dig
	}
	return ditab
}

// Округление float64 до dig знаков после запятой
func roundto(val float64, dig int) float64 {
	d := float64(dig)
	return math.Round(val*d) / d
}

// Выводит таблицу distable
// func TestDistable(t *testing.T) {
// 	dt := distable()
// 	for i := 0; i < 90; i++ {
// 		ipar := float64(i)
// 		t.Error(ipar, dt[ipar])
// 	}
// }
