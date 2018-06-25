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
			case *comejob:
				log.Println("comejob")
				newComeJob, err := inbound.(*comejob)
				if err == false {
					s.members[newComeJob.lifer] = true
				}
			case *awayjob:
				log.Println("awayjob")
			}
		}
	}
}
