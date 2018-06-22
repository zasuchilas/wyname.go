package main

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

var (
	ses = []int{sn, sm, sf}
	sea = []int{an, aa, ab, ac, ad, ae, af}
	sef = []int{ma, mb, mc, md, me, mf, fa, fb, fc, fd, fe, ff}
)

// desamf parses the samf
func desamf(samf int) (sex int, age int, sa int, filter int, mark string) {
	var count int
	// sex
	for _, v := range ses {
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
	for _, v := range sea {
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
		// sa = sex | age -> filters representation
		switch {
		case sex == sm && age == aa:
			sa = ma
		case sex == sm && age == ab:
			sa = mb
		case sex == sm && age == ac:
			sa = mc
		case sex == sm && age == ad:
			sa = md
		case sex == sm && age == ae:
			sa = me
		case sex == sm && age == af:
			sa = mf
		case sex == sf && age == aa:
			sa = fa
		case sex == sf && age == ab:
			sa = fb
		case sex == sf && age == ac:
			sa = fc
		case sex == sf && age == ad:
			sa = fd
		case sex == sf && age == ae:
			sa = fe
		case sex == sf && age == af:
			sa = ff
		}
	}
	// filter
	count = 0
	for _, v := range sef {
		if (samf & v) != 0 {
			count++
			if count == 1 {
				filter = v
			} else {
				filter = filter | v
			}
		}
	}
	if count > 9 || count == 0 {
		filter = 0
	}
	// zero
	if sa == 0 || filter == 0 {
		sa = 0
		filter = 0
	}
	// mark
	switch {
	case sa == 0 || filter == 0:
		mark = "6"
	case sex == sf:
		mark = "0"
	case sex == sm:
		mark = "1"
	default:
		mark = "6"
	}
	return
}

// chat calculates the intersection of lifers and return true if yes
func chat(sa1 int, f1 int, sa2 int, f2 int) (intersect bool) {
	// zero
	if sa1 == 0 || sa2 == 0 {
		if sa1 == 0 && sa2 == 0 {
			intersect = true
		} else {
			intersect = false
		}
		return
	}
	// non zero
	if ((sa1 & f2) != 0) && ((f1 & sa2) != 0) {
		intersect = true
	} else {
		intersect = false
	}
	return
}
