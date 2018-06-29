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
			awaySector.broadcast <- createAwayJob(l)
		}
	}
}

// unsubscribeEverywhere removes all subscribtions in sectors
func (l *Lifer) unsubscribeEverywhere(notify bool) {
	if len(l.subscriptions) > 0 {
		awayUnsubscrJob := createUnsubscribeJob(l, notify)
		for awayname := range l.subscriptions {
			if awaySubs, found := l.secache[awayname]; found {
				awaySubs.broadcast <- awayUnsubscrJob
			}
		}
	}
}

func (l *Lifer) logB() {
	// -> log : part, hash, samf, sex, age, lat, lon
	l.mutex.RLock()
	samf := l.inboundSamf
	sex := l.sex
	age := l.age
	lat := l.inboundLat
	lon := l.inboundLon
	l.mutex.RUnlock()
	log.Println("B," + l.hash + "," + samf + "," + strconv.Itoa(sex) + "," + strconv.Itoa(age) + "," + lat + "," + lon)
}

func (l *Lifer) logC() {
	// -> log : part, hash, samf, sex, age, lat, lon
	l.mutex.RLock()
	samf := l.inboundSamf
	sex := l.sex
	age := l.age
	lat := l.inboundLat
	lon := l.inboundLon
	l.mutex.RUnlock()
	log.Println("C," + l.hash + "," + samf + "," + strconv.Itoa(sex) + "," + strconv.Itoa(age) + "," + lat + "," + lon)
}

func (l *Lifer) changeSamfData(inbsamf int, inbsamft string) {
	lsex, lage, lsa, lfilter, lfilters, lmark := desamf(inbsamf)
	l.mutex.Lock()
	l.samf = inbsamf
	l.inboundSamf = inbsamft
	l.sex = lsex
	l.age = lage
	l.sa = lsa
	l.filter = lfilter
	l.filters = lfilters
	l.mark = lmark
	l.mutex.Unlock()
}
