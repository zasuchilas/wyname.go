package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// auth checks access (validate client)
func auth(path string) (access bool) {
	// path length
	if len(path) < 26 {
		return false
	}

	srvNowSeconds := time.Now().Unix() // srv time remember

	// /mls3.sec3.rnd3.mls3c.sec3c.rnd3c.befen
	// /1112223334445556667777777
	// 0123456789012345678901234567
	// 0000000000111111111122222222

	// mls3
	mls3 := path[1:4]
	mls3c := path[10:13]
	mls3calc, err := chsum(mls3)
	if err != nil || mls3c != mls3calc {
		log.Println("mls3 check failure")
		return false
	}
	// sec3
	sec3 := path[4:7]
	sec3c := path[13:16]
	sec3calc, err := chsum(sec3)
	if err != nil || sec3c != sec3calc {
		log.Println("sec3 check failure")
		return false
	}
	// rnd3
	rnd3 := path[7:10]
	rnd3c := path[16:19]
	rnd3calc, err := chsum(rnd3)
	if err != nil || rnd3c != rnd3calc {
		log.Println("rnd3 check failure")
		return false
	}
	// synchr check
	befen := string(path[19])
	befkey := mls3 + sec3 + rnd3 + mls3c + sec3c + rnd3c
	bef, err := befdecode(befen, befkey)
	if err != nil {
		log.Println("bef check failure")
		return false
	}
	srvcalc, err := strconv.ParseInt(bef+sec3, 10, 64)
	if err != nil {
		log.Println("srvcalc parse failure")
		return false
	}
	diff := srvNowSeconds - srvcalc
	if diff < -1 || diff > 2 {
		log.Println("srv time check failure, diff:", diff)
		return false
	}
	return true
}

var chtab = map[int]map[string]string{
	1: {"0": "7", "1": "4", "2": "6", "3": "5", "4": "9", "5": "0", "6": "8", "7": "1", "8": "2", "9": "3"},
	2: {"0": "4", "1": "7", "2": "0", "3": "6", "4": "8", "5": "1", "6": "3", "7": "5", "8": "9", "9": "2"},
	3: {"0": "5", "1": "6", "2": "4", "3": "8", "4": "0", "5": "3", "6": "9", "7": "2", "8": "7", "9": "1"},
}

// chsum return 3 digits
func chsum(d3 string) (ch3 string, err error) {
	if len(d3) != 3 {
		err = fmt.Errorf("maper chsum d3 != 3")
		return
	}
	succs := 0
	for i := 1; i <= 3; i++ {
		b := d3[i-1 : i]
		if n, found := chtab[i][b]; found {
			succs++
			ch3 += n
		}
	}
	if succs != 3 {
		err = fmt.Errorf("maper chsum succs != 3")
		ch3 = ""
	}
	return
}

func befdecode(ben, key string) (bef string, err error) {
	benlen := len(ben)
	keylen := len(key)
	if benlen < 7 || keylen < 7 || benlen > keylen {
		err = fmt.Errorf("maper befdecode benlen or keylen wrong")
		return
	}
	succs := 0
	for i := 1; i <= benlen; i++ {
		b, eb := strconv.Atoi(ben[i-1 : i])
		k, ek := strconv.Atoi(key[i-1 : i])
		if eb == nil && ek == nil {
			succs++
			if b < k {
				b = b + 10
			}
			n := b - k
			bef += strconv.Itoa(n)
		}
	}
	if succs != benlen {
		err = fmt.Errorf("maper befdecode succs != benlen")
		return
	}
	bef, err = befdeconst(bef, 7)
	return
}

func befdeconst(ben string, key int) (bef string, err error) {
	benlen := len(ben)
	if key < 0 || key > 9 || benlen < 7 {
		err = fmt.Errorf("maper befdeconst benlen or key failure")
		return
	}
	succs := 0
	for i := 1; i <= benlen; i++ {
		b, eb := strconv.Atoi(ben[i-1 : i])
		if eb == nil {
			succs++
			if b < key {
				b = b + 10
			}
			n := b - key
			bef += strconv.Itoa(n)
		}
	}
	if succs != benlen {
		err = fmt.Errorf("maper befdeconst succs != benlen")
		bef = ""
		return
	}
	return
}
