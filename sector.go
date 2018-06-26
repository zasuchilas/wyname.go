package main

import (
	"log"
)

// Sector регистрирует лайферов
type Sector struct {
	members   map[*Lifer]bool // members of sector
	subscrs   map[*Lifer]bool // sector subscribers
	broadcast chan job        // inbound messages from lifers
}

func newsector() *Sector {
	return &Sector{
		members:   make(map[*Lifer]bool, 500),
		subscrs:   make(map[*Lifer]bool, 500),
		broadcast: make(chan job),
	}
}

func (s *Sector) run() {
	for {
		select {
		case inbound := <-s.broadcast:
			switch inbound.(type) {
			case *jobCome:
				log.Println("comejob")
				newComeJob, err := inbound.(*jobCome)
				if err == false {
					s.members[newComeJob.lifer] = true
					// notify subscribers
				}
			case *jobAway:
				log.Println("awayjob")
				newAwayJob, err := inbound.(*jobAway)
				if err == false {
					delete(s.members, newAwayJob.lifer)
					// notify subscribers
				}
			case *jobSubscribe:
				log.Println("jobSubscribe")
				newSubscrJob, err := inbound.(*jobSubscribe)
				if err == false {
					s.subscrs[newSubscrJob.lifer] = true
					// get package
				}
			case *jobUnsubscribe:
				log.Println("jobUnsubscribe")
				newUnsubscribeJob, err := inbound.(*jobUnsubscribe)
				if err == false {
					delete(s.subscrs, newUnsubscribeJob.lifer)
				}
			}
		}
	}
}
