package main

import (
	"log"
	"strconv"
)

// connectFirst connect first and l.started = true and log B
func (l *Lifer) connectFirst() {
	l.secache[l.membership].broadcast <- createComeJob(l)
	for secname := range l.subscriptions {
		l.secache[secname].broadcast <- createSubscribeJob(l)
	}
	l.started = true
}

// awayFromMembers removes lifer from members in his sector
func (l *Lifer) awayFromMembers() {
	if l.membership != "" {
		if awaySector, found := l.secache[l.membership]; found {
			filters := make([]int, len(l.filters))
			copy(filters, l.filters)
			awaySector.broadcast <- createAwayJob(l, l.sa, l.filter, filters)
		}
	}
}

// unsubscribeEverywhere removes all subscribtions in sectors
func (l *Lifer) unsubscribeEverywhere() {
	if len(l.subscriptions) > 0 {
		awayUnsubscrJob := createUnsubscribeJob(l)
		for awayname := range l.subscriptions {
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
	l.sex, l.age, l.sa, l.filter, l.filters, l.mark = desamf(inbsamf)
}
