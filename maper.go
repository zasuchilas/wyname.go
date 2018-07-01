package main

import (
	"fmt"
	"log"
)

// auth checks access (validate client)
func auth(path string) (access bool) {
	// path length
	if len(path) < 26 {
		return false
	}
	// srv time
	// srvNowSeconds := time.Now().Unix()
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
	// befen := path[19]

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
