package main

import (
	"log"
	"strconv"
)

// connectFirst connect first and l.started = true and log B
func (l *Lifer) connectFirst() {
	l.secache[l.cmember].broadcast <- createComeJob(l)
	for secname := range l.csubscr {
		l.secache[secname].broadcast <- createSubscribeJob(l)
	}
	l.started = true
	// -> log : part, hash, samf, sex, age, lat, lon
	log.Println("B," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
}

// awayFromMembers removes lifer from members in his sector
func (l *Lifer) awayFromMembers() {
	if l.cmember != "" {
		if awaySector, found := l.secache[l.cmember]; found {
			awaySector.broadcast <- createAwayJob(l)
		}
	}
}

// unsubscribeEverywhere removes all subscribtions in sectors
func (l *Lifer) unsubscribeEverywhere() {
	if len(l.csubscr) > 0 {
		awayUnsubscrJob := createUnsubscribeJob(l)
		for awayname := range l.csubscr {
			if awaySubs, found := l.secache[awayname]; found {
				awaySubs.broadcast <- awayUnsubscrJob
			}
		}
	}
}
