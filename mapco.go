package main

import (
	"fmt"
	"math"
	"strconv"
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
		s[i], _ = sector(v[0], v[1])
	}

	return s, e
}

/// Получение сектора точки по её координатам
/// т.е. точку Y (max, max) для данных констант las и los
func sector(lat, lon float64) (string, error) {
	lonPart, _ := sectorsLonPart(lon)
	latPart, _ := sectorsLatPart(lat)
	return strconv.Itoa(latPart) + "|" + strconv.Itoa(lonPart), nil
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

// Взять градусы, вернуть целое число
func mega(v float64) int {
	r := int(math.Round(v * 1000000))
	return r
}

// Gps координаты
type Gps struct {
	lat float64
	lon float64
}

func newGps(lat, lon float64) (*Gps, error) {
	var e error
	if lat > 180 || lat < -180 || lon > 90 || lon < -90 {
		e = fmt.Errorf("coordinates out og range")
	}
	return &Gps{lat, lon}, e
}

// Вычислить все секторы
// [0] member sector
func (g *Gps) compute() ([]string, error) {
	// нужно вычислить точки A и C

	return []string{}, nil
}

// var latab [90]float64

func latable() map[int]int {
	latab := make(map[int]int, 90)
	var equator float64 = 111321
	for parallel := 0; parallel < 90; parallel++ {
		ipar := float64(parallel)
		latab[parallel] = int(math.Round(equator * math.Cos(ipar*(math.Pi/180))))
	}
	return latab
}

func distable() map[float64]float64 {
	ditab := make(map[float64]float64, 90)
	equator := 0.0052910052910053
	var dig float64 = 100000
	for parallel := 0; parallel < 900; parallel++ {
		ipar := float64(parallel)
		ditab[ipar] = math.Round((equator/math.Cos(ipar*(math.Pi/180)))*dig) / dig
	}
	return ditab
}

const equator float64 = 0.0052910052910053

// dist588 возвращает смещение в градусах на 588 метров для данных координат
func (g *Gps) dist589() float64 {
	return equator / math.Cos(g.lat*(math.Pi/180))
	/*
		вдоль мередиана смешение в градусах постоянно
		вдоль параллели - изменяется в зависимости от широты
		мне нужно вычислить точки А и С от точки I
		т.е. произвести смещение вдоль мередианов и парралелей
		и нужно знать на сколько градусов смещаться для данной широты
		(при движении вдоль параллели, т.е. когда изменяю долготу)
		l = le * math.cos(lat)
		le = // длина дуги на экваторе (1" = 31 метр)
		111321 м в 1 градусе на экваторе (при движении вдоль параллели)
		это 189 раз по 589 метров
		это значит 0.0052910052910053 градусов = 589 метров
		это то число, на которое для точки на экваторе нужно увеличивать или уменьшать
		исходную долготу, чтобы получить долготу точек А и С
		по таблице на широте 57 в 1 градусе 60773 метров
		т.е. 60773/589 = 103,1799660441426 раз, т.е. 0,0096918101296861, т.е. примероно 0,01
		поэтому мне нужно было умножать на 2, чтобы сделать сектор квадратным
		по формуле: 0.0052910052910053 / cos(57) = 0.0052910052910053 / 0,54463903501502708222408369208157 = 0,0097146917499481307802048696476 ~ 0.01
	*/

	/*
		Шаг сектора:
		вариант 1: 0.005 на экваторе будут квадратные в этом случае
		вариант 2: 0.01
		вариант 3: 0.001 секторов будет больше (т.е. подписок тоже), но ближе к границе выступать квадраты будут - (1000*180)^2 = 32.400.000.000 секторов будет - не подходит,
		т.к. даже на экваторе будет (5+5)*(5+5) = 100 подписок
		Значит, вариант 2
		Поскольку секторы больше дистанции смещения при поиске А и С, то количество секторов может быть меньше 4 - протестируем на разных точках земной поверхности
	*/

	/*
		Северный и южный полюсы
		Все кто выше 89 и ниже -89 - в одном секторе (не беда, там всего может 1000 человек)
	*/

	/*
		Гринвич - нулевой меридиан
		?
	*/

	// случай когда параллель -> 90 или -90 (будет очень много секторов)
	// return 3.14
}

// Получить секторы экрана
// вернее, области включающей 2 точки
// A нижняя левая, C верхняя правая
/*
static Set screenSectors(List tA, List tC) {
	Set sects = new Set();

	// checks
	if ((tA[0] > tC[0]) || (tA[1] > tC[1])) return sects;

	// lons
	Set losSet = new Set();
	int losC = sectorsLonPart(tC[1]); // lo крайнего справа сектора C
	int losA = sectorsLonPart(tA[1]); // lo крайнего слева сектора A
	int computedSectLon = losA;
	while (computedSectLon <= losC) {
		losSet.add(computedSectLon);
		computedSectLon += 1; // прибавляем по одному сектору
	}
	// print('losSet: $losSet');

	// lats
	Set lasSet = new Set();
	int lasC = sectorsLatPart(tC[0]); // la крайнего сверху сектора C
	int lasA = sectorsLatPart(tA[0]); // la крайнего снизу сектора A
	int computedSectLat = lasA;
	while (computedSectLat <= lasC) {
		lasSet.add(computedSectLat);
		computedSectLat += 5; // прибавляем по одному секторы (у lat шаг 5)
	}
	// print('lasSet: $lasSet');

	lasSet.forEach((e1) {
		losSet.forEach((e2) {
			sects.add('${e1}|${e2}');
		});
	});
	// print('sects: $sects');

	return sects;
}

  /// Возвращает дистанцию между точками в метрах
  static num distance(List pa, List pb) {
    if (pa.length != 2 || pb.length != 2) return 0;
    num ac = (pa[0]-pb[0])/la1; // АС в метрах
    num bc = (pa[1]-pb[1])/lo1; // BC в метрах
    num ab = sqrt(pow(ac, 2) + pow(bc, 2));
    return ab;
	}

  static bool validate(dynamic mla, dynamic mlo) {
    bool valid = false;
    if (mla is int && mlo is int) { // if null -> false
      if (mla < mlamax && mla > -(mlamax)
          && mlo < mlomax && mlo > -(mlomax)) valid = true;
    }
    return valid;
  }
*/
