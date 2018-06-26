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

func (l *Lifer) logB() {
	// -> log : part, hash, samf, sex, age, lat, lon
	log.Println("B," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
}

func (l *Lifer) logC() {
	// -> log : part, hash, samf, sex, age, lat, lon
	log.Println("C," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
}

func (l *Lifer) changeSamfData(inbsamf int, inbsamft string) {
	l.samf = inbsamf
	l.inboundSamf = inbsamft
	l.sex, l.age, l.sa, l.filter, l.mark = desamf(inbsamf)
}

// clearCalculatedData clears the computed data (for reconnect with new calculating values)
// func (l *Lifer) clearCalculatedData() {
// 	l.cmember = ""
// 	for k := range l.csubscr {
// 		delete(l.csubscr, k)
// 	}
// }
