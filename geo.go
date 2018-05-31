package main

import (
	"fmt"
	"math"
	"strconv"
)

const equator float64 = 0.0052910052910053

// dist588 возвращает смещение в градусах на 589 метров для данных координат
func (g *Gps) dist589() (north, east float64) {
	north = equator
	east = math.Abs(equator / math.Cos(g.lat*(math.Pi/180)))
	return
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

// Gps координаты
type Gps struct {
	lat float64
	lon float64
}

func newGps(lat, lon float64) (gps *Gps, err error) {
	if lat > 180 || lat < -180 || lon > 90 || lon < -90 {
		err = fmt.Errorf("coordinates out of range")
	}
	// TODO: для крайних случаев делаем обработку (180, -180, 90, -90, 0)
	// для большого превышения ? возвращаем ошибку, но крайние случаи - норма, обрабатываем
	return &Gps{lat, lon}, err
}

// Вычислить все секторы
// [0] member sector
func (g *Gps) compute() (secs []string, err error) {
	// main sector isec (Iam)
	ila, ilo := g.sectornums()
	isec := strconv.Itoa(ila) + ":" + strconv.Itoa(ilo)

	// нужно вычислить точки A и C
	north, east := g.dist589()
	a, err := newGps(g.lat-north, g.lon-east)
	if err != nil {
		err = fmt.Errorf("failed to create point A")
	}
	c, err := newGps(g.lat+north, g.lon+east)
	if err != nil {
		err = fmt.Errorf("failed to create point C")
	}
	ala, alo := a.sectornums()
	cla, clo := c.sectornums()
	cap := (cla - ala + 1) * (clo - alo + 1) // TODO test this with different placies & 180
	secs = make([]string, cap)
	secs[0] = isec
	idx := 1
	var n string
	for i := ala; i <= cla; i++ {
		for j := alo; j <= clo; j++ {
			n = strconv.Itoa(i) + ":" + strconv.Itoa(j)
			if n != isec {
				secs[idx] = n
				idx++
			}
		}
	}

	return secs, nil
}

// Получить числовые части названия сектора точки
// секторы делятся по 0.01 т.е. 57.01 57.02 ..
// сектор определяется верхней правой точкой
// т.е. 57.633895, 39.834598 -> 5764, 3984
func (g *Gps) sectornums() (lat int, lon int) {
	lat = int(math.Ceil(g.lat * 100))
	lon = int(math.Ceil(g.lon * 100))
	return
}
