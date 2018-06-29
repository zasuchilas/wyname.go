package main

import (
	"log"
)

// Sector регистрирует лайферов
type Sector struct {
	name      string                  // sector name
	members   map[int]map[*Lifer]bool // members of sector
	subscrs   map[int]map[*Lifer]bool // sector subscribers
	broadcast chan job                // inbound messages from lifers
}

func newsector(name string) *Sector {
	members := make(map[int]map[*Lifer]bool, 13)
	subscrs := make(map[int]map[*Lifer]bool, 13)
	members[0] = make(map[*Lifer]bool, 101)
	subscrs[0] = make(map[*Lifer]bool, 101)
	for _, sefsa := range sef {
		members[sefsa] = make(map[*Lifer]bool, 101)
		subscrs[sefsa] = make(map[*Lifer]bool, 101)
	}
	return &Sector{
		name:      name,
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
				job, err := inbound.(*jobCome)
				if err == false {
					l := job.lifer
					l.mutex.RLock()

					l.mutex.RUnlock()
					s.members[l.sa][l] = true
					s.move(l) // notify subscribers about come (move)
				}
			case *jobMove:
				job, err := inbound.(*jobMove)
				if err == false {
					log.Println("movejob", job.lifer)
					s.move(job.lifer) // notify subscribers about move
				}
			case *jobAway:
				log.Println("awayjob")
				job, err := inbound.(*jobAway)
				if err == false {
					delete(s.members[job.sa], job.lifer)
					s.away(job.lifer, job.sa, job.filter, job.filters) // notify subscribers about away
				}
			case *jobSubscribe:
				log.Println("jobSubscribe")
				job, err := inbound.(*jobSubscribe)
				if err == false {
					l := job.lifer
					s.subscrs[l.sa][l] = true
					s.pack(l) // send sector package to lifer
				}
			case *jobUnsubscribe:
				log.Println("jobUnsubscribe")
				job, err := inbound.(*jobUnsubscribe)
				if err == false {
					l := job.lifer
					delete(s.subscrs[job.sa], l)
					l.send <- []byte(codeSectorUnpack + "," + s.name) // send remove sector points
				}
			case *jobGlob:
				log.Println("jobGlob")
				job, err := inbound.(*jobGlob)
				if err == false {
					s.glob(job.lifer, job.globReqCode)
				}
			}
		}
	}
}
