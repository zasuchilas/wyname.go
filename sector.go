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
		members:   make(map[*Lifer]bool, 200),
		subscrs:   make(map[*Lifer]bool, 300),
		broadcast: make(chan job),
	}
}

func (s *Sector) run() {
	for {
		select {
		case inbjob := <-s.broadcast:
			switch inbjob.(type) {
			case *comejob:
				log.Println("comejob")
			case *awayjob:
				log.Println("awayjob")
			}
		}
	}
}
