package main

import "strings"

const (
	sn = 1       // 1
	sm = 2       // 2
	sf = 4       // 3
	an = 8       // 4
	aa = 16      // 5
	ab = 32      // 6
	ac = 64      // 7
	ad = 128     // 8
	ae = 256     // 9
	af = 512     // 10
	ma = 1024    // 11
	mb = 2048    // 12
	mc = 4096    // 13
	md = 8192    // 14
	me = 16384   // 15
	mf = 32768   // 16
	fa = 65536   // 17
	fb = 131072  // 18
	fc = 262144  // 19
	fd = 524288  // 20
	fe = 1048576 // 21
	ff = 2097152 // 22
)

var se = map[string]int{"sn": sn, "sm": sm, "sf": sf, "an": an, "aa": aa, "ab": ab, "ac": ac, "ad": ad, "ae": ae, "af": af, "ma": ma, "mb": mb, "mc": mc, "md": md, "me": me, "mf": mf, "fa": fa, "fb": fb, "fc": fc, "fd": fd, "fe": fe, "ff": ff}

var es = map[int]string{sn: "n", sm: "m", sf: "f", an: "n", aa: "a", ab: "b", ac: "c", ad: "d", ae: "e", af: "f", ma: "ma", mb: "mb", mc: "mc", md: "md", me: "me", mf: "mf", fa: "fa", fb: "fb", fc: "fc", fd: "fd", fe: "fe", ff: "ff"}

var (
	ses = []int{sn, sm, sf}
	sea = []int{an, aa, ab, ac, ad, ae, af}
	sef = []int{ma, mb, mc, md, me, mf, fa, fb, fc, fd, fe, ff}
	// sem = []int{ma, mb, mc, md, me, mf}
	// sef = []int{fa, fb, fc, fd, fe, ff}
)

// desamf parses the samf
func desamf(samf int) (sex int, age int, sa int, filters int) {
	var count int
	// sex
	for v := range ses {
		if (samf & v) != 0 {
			count++
			sex = v
		}
	}
	if count > 1 || count == 0 {
		sex = sn
	}
	// age
	count = 0
	for v := range sea {
		if (samf & v) != 0 {
			count++
			age = v
		}
	}
	if count > 1 || count == 0 {
		age = an
	}
	// sa
	if sex == sn || age == an {
		sa = 0
	} else {
		sa = sex | age
	}
	// filter
	count = 0
	for v := range sef {
		if (samf & v) != 0 {
			count++
			if count == 1 {
				filters = v
			} else {
				filters = filters | v
			}
		}
	}
	if count > 9 {
		filters = 0
	}
	return
}

func pressed(samf int, i int) bool {
	// if i == nil {
	// 	return false
	// }
	return (samf & i) != 0
}

// sexfrom returns sex from samf
func sexfrom(samf int) int {
	var s int
	n := (samf & sn) != 0 // is sn
	m := (samf & sm) != 0 // is sm
	f := (samf & sf) != 0 // is sf
	if n || (m && n) {
		s = sn
	} else if m {
		s = sm
	} else if f {
		s = sf
	} else {
		s = sn
	}
	return s
}

func sexFromSamf(samf int) string {
	s := decode(sn)
	for i := 0; i < len(ses); i++ {
		if pressed(samf, ses[i]) {
			s = decode(ses[i])
			break
		}
	}
	return s
}

func ageFromSamf(samf int) string {
	a := decode(an)
	for i := 0; i < len(sea); i++ {
		if pressed(samf, sea[i]) {
			a = decode(sea[i])
			break
		}
	}
	return a
}

/// Извлекает s и a из samf
func saFromSamf(samf int) string {
	s := decode(sn)
	a := decode(an)
	sa := "nn"
	for i := 0; i < len(ses); i++ {
		if pressed(samf, ses[i]) {
			s = decode(ses[i])
			break
		}
	}
	for i := 0; i < len(sea); i++ {
		if pressed(samf, sea[i]) {
			a = decode(sea[i])
			break
		}
	}
	if s == "n" || a == "n" {
		sa = "nn"
	} else {
		sa = s + a
	}
	return sa
}

/// Извлекает все f из samf
func mfFromSamf(samf int) []string {
	f := []string{}
	for i := 0; i < len(sem); i++ {
		if pressed(samf, sem[i]) {
			f = append(f, decode(sem[i]))
		}
	}
	for i := 0; i < len(sef); i++ {
		if pressed(samf, sef[i]) {
			f = append(f, decode(sef[i]))
		}
	}
	return f
}

/// Возвращает строку значений всех фильтров
func filtersFromSamf(samf int) string {
	e := "--"
	f := make([]string, 5)
	for i := range sem {
		if pressed(samf, sem[i]) {
			f = append(f, decode(sem[i]))
		} else {
			f = append(f, e)
		}
	}
	// for i := 0; i < len(sef); i++ {
	for i := range sef {
		if pressed(samf, sef[i]) {
			f = append(f, decode(sef[i]))
		} else {
			f = append(f, e)
		}
	}
	return strings.Join(f, ",")
}

func decode(digit int) string {
	return es[digit]
}

func chats(samf int) []string {
	sa := saFromSamf(samf)
	chs := make([]string, 5)
	if sa != "nn" {
		filters := mfFromSamf(samf)
		for _, f := range filters {
			chs = append(chs, sa+f)
		}
	}
	if len(chs) == 0 {
		chs = append(chs, "nnnn")
	}
	return chs
}

func mirror(chat string) string {
	if len(chat) != 4 {
		return "nnnn"
	}
	return chat[2:] + chat[:2]
}

func underchats(chs []string) []string {
	uchs := make([]string, 5)
	for _, chat := range chs {
		uchs = append(uchs, mirror(chat))
	}
	if len(uchs) == 0 {
		uchs = append(uchs, "nnnn")
	}
	return uchs
}
