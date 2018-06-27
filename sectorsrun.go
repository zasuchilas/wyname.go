package main

import (
	"log"
)

// Sector регистрирует лайферов
type Sector struct {
	members   map[int]map[*Lifer]bool // members of sector
	subscrs   map[int]map[*Lifer]bool // sector subscribers
	broadcast chan job                // inbound messages from lifers
}

func newsector() *Sector {
	members := make(map[int]map[*Lifer]bool, 13)
	subscrs := make(map[int]map[*Lifer]bool, 13)
	members[0] = make(map[*Lifer]bool, 101)
	subscrs[0] = make(map[*Lifer]bool, 101)
	for _, sefsa := range sef {
		members[sefsa] = make(map[*Lifer]bool, 101)
		subscrs[sefsa] = make(map[*Lifer]bool, 101)
	}
	return &Sector{
		members:   members,
		subscrs:   subscrs,
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
					s.members[newComeJob.lifer.sa][newComeJob.lifer] = true
					// notify subscribers about come = move

					// for i, k := range s.subscrs {

					// }
				}
			case *jobMove:
				newMoveJob, err := inbound.(*jobMove)
				if err == false {
					log.Println("movejob", newMoveJob.lifer)
					// notify subscribers about move
				}
			case *jobAway:
				log.Println("awayjob")
				newAwayJob, err := inbound.(*jobAway)
				if err == false {
					delete(s.members[newAwayJob.lifer.sa], newAwayJob.lifer)
					// notify subscribers about away
				}
			case *jobSubscribe:
				log.Println("jobSubscribe")
				newSubscrJob, err := inbound.(*jobSubscribe)
				if err == false {
					s.subscrs[newSubscrJob.lifer.sa][newSubscrJob.lifer] = true
					// get package
				}
			case *jobUnsubscribe:
				log.Println("jobUnsubscribe")
				newUnsubscribeJob, err := inbound.(*jobUnsubscribe)
				if err == false {
					delete(s.subscrs[newUnsubscribeJob.lifer.sa], newUnsubscribeJob.lifer)
					// send unsubscribe sector
				}
			}
		}
	}
}
