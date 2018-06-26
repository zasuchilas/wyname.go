package main

import (
	"log"
	"strconv"
)

// connectFirst connect first and l.started = true and log B
func (l *Lifer) connectFirst() {
	l.secache[l.cmember].broadcast <- newcomejob(l)
	for secname := range l.csubscr {
		l.secache[secname].broadcast <- newSubscribeJob(l)
	}
	l.started = true
	// -> log : part, hash, samf, sex, age, lat, lon
	log.Println("B," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
}
