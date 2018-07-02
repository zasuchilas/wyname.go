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
		broadcast: make(chan job, 2048),
	}
}

func (s *Sector) run() {
	for {
		select {
		case inbound := <-s.broadcast:
			switch inbound.(type) {
			case *jobCome:
				j, errCome := inbound.(*jobCome)
				log.Println("comejob", j, "err:", errCome)
				// if errCome == false {
				s.members[j.sa][j.lifer] = true
				log.Println("come members", s.members)
				s.move(j.lifer, j.lat, j.lon, j.mark, j.sa, j.filter, j.filters) // notify subscribers about come (move)
				// }
			case *jobMove:
				j, errMove := inbound.(*jobMove)
				log.Println("movejob", j, "err:", errMove)
				// if errMove == false {
				s.move(j.lifer, j.lat, j.lon, j.mark, j.sa, j.filter, j.filters) // notify subscribers about move
				// }
			case *jobAway:
				j, err := inbound.(*jobAway)
				log.Println("awayjob", j, "err:", err)
				// if err == false {
				delete(s.members[j.sa], j.lifer)
				s.away(j.lifer, j.sa, j.filter, j.filters) // notify subscribers about away
				// }
			case *jobSubscribe:
				j, err := inbound.(*jobSubscribe)
				log.Println("jobSubscribe", j, "err:", err)
				// if err == false {
				s.subscrs[j.sa][j.lifer] = true
				s.pack(j.lifer, j.sa, j.filter, j.filters) // send sector package to lifer
				// }
			case *jobUnsubscribe:
				j, err := inbound.(*jobUnsubscribe)
				log.Println("jobUnsubscribe", j, "err:", err)
				// if err == false {
				l := j.lifer
				delete(s.subscrs[j.sa], l)
				l.send <- []byte(codeSectorUnpack + "," + s.name) // send remove sector points
				// }
			case *jobGlob:
				j, err := inbound.(*jobGlob)
				log.Println("jobGlob", j, "err:", err)
				// if err == false {
				s.glob(j.lifer, j.sa, j.filter, j.filters, j.globReqCode)
				// }
			}
		}
	}
}
