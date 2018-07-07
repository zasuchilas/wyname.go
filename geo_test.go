package main

import (
	"math"
	"testing"
)

var placies = [][2]float64{
	{57.626569, 39.893787},   // Ярославль
	{55.755814, 37.617635},   // Москва
	{59.939095, 30.315868},   // Санкт-Петербург
	{41.388698, 2.169018},    // Барселона
	{41.696583, 2.834054},    // Льорет де Мар
	{48.856651, 2.351691},    // Париж
	{51.511118, -0.085708},   // Лондон
	{52.208652, 0.126709},    // Кембридж
	{-33.926706, 18.416187},  // Кейптаун
	{-15.795027, -47.900866}, // Бразилиа
	{40.814742, -73.940556},  // Нью-Йорк
	{41.029760, -73.764490},  // Уайт-Плейнс
	{34.031620, -118.282082}, // Лос-Анджелес
	{1.331412, 103.804368},   // Сингапур
	{-33.851008, 151.176392}, // Сидней
	{-41.302729, 174.778453}, // Веллингтон
}

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
func TestDistable(t *testing.T) {
	dt := distable()
	for i := 0; i < 90; i++ {
		ipar := float64(i)
		t.Error(ipar, dt[ipar])
	}
}

func TestCalculate(t *testing.T) {
	yarplace := [2]float64{57.626569, 39.893787}
	yar, _ := newGps(yarplace[0], yarplace[1])
	m, a, _ := yar.calculate()
	t.Error("membersec", m)
	t.Error("allsectors", a)
}
