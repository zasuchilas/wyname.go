package main

import (
	"fmt"
	"math"
)

const (
	// соотношение градусов и метров
	la1 = 0.0000085 // градусов = 1м
	lo1 = 0.0000170 // градусов = 1м

	// sector sides (отрезки в градусах)
	las = 0.005 // lat 0.005 ~588m
	los = 0.010 // lon 0.010 ~588m

	// square Iam (отрезки в градусах)
	laq = las
	loq = los

	// пределы mla mlo
	mlamax = 90000000  // max la
	mlomax = 180000000 // max lo
)

/// Получить секторы (для подписки)
/// по точке gps (mla, mlo)
func sectors(mla, mlo int) (map[int]string, error) {
	// возвратить нужно Set (без повторений)
	// Set s = new Set();
	s := make(map[int]string, 8)
	var e error

	la := micro(mla)
	lo := micro(mlo)
	lam := la - laq
	lap := la + laq
	lom := lo - loq
	lop := lo + loq

	// 8 points
	points := make(map[int][2]float64, 8) // [2]int это массив размером 2
	points[1] = [2]float64{lam, lom}      // tA
	points[8] = [2]float64{lam, lo}       // tH
	points[5] = [2]float64{la, lom}       // tE
	points[2] = [2]float64{lap, lom}      // tB
	points[6] = [2]float64{lap, lo}       // tF
	points[3] = [2]float64{lap, lop}      // tC
	points[7] = [2]float64{la, lop}       // tG
	points[4] = [2]float64{lam, lop}      // tD

	for i, v := range points {
		fmt.Println(i, v)
		s[i], _ = sector(v[0], v[1])
	}
	// for (var i = 1; i <= 8; i++) {
	// 	// print('$i: ${points[i]}');
	// 	s.add(sector(points[i][0], points[i][1]));
	// }

	return s, e
}

/// Получение сектора точки по её координатам
/// т.е. точку Y (max, max) для данных констант las и los
func sector(lat, lon float64) (string, error) {
	lonPart, _ := sectorsLonPart(lon)
	latPart, _ := sectorsLatPart(lat)
	return string(latPart) + "|" + string(lonPart), nil
}

/// Получить latPart ключа сектора
func sectorsLatPart(la float64) (int, error) {
	var latPart int
	lat100 := la * 100
	latCeil := int(math.Ceil(lat100))
	latRound := int(math.Round(lat100))
	if latCeil == latRound {
		latPart = latCeil * 10
	} else {
		latPart = latRound*10 + 5 // '${latRound.toString()}5';
	}
	// ? для отрицательных lat должно нормально работать todo test!
	// т.е. тот же принцип Y (правая верхняя точка)
	return latPart, nil
}

/// Получить lonPart ключа сектора
func sectorsLonPart(lo float64) (int, error) {
	return int(math.Ceil(lo * 100)), nil
}

// Взять целое число (представление градусов)
// вернуть градусы
func micro(v int) float64 {
	r := float64(v) / 1000000
	return r
}
