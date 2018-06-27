package main

// move notify subscribers about move
func (s *Sector) move(l *Lifer) {
	hash := l.hash
	lat := l.inboundLat
	lon := l.inboundLon
	mark := l.mark
	for _, lf := range l.filters {
		for subscriber := range s.subscrs[lf] {
			if l != subscriber && chat(l.sa, l.filter, subscriber.sa, subscriber.filter) {
				subscriber.send <- []byte(codeMove + "," + hash + "," + lat + "," + lon + "," + mark)
			}
		}
	}
}

func (s *Sector) away(l *Lifer, sa int, filter int, filters []int) {
	hash := l.hash
	sector := s.name
	for _, lf := range filters {
		for subscriber := range s.subscrs[lf] {
			if l != subscriber && chat(sa, filter, subscriber.sa, subscriber.filter) {
				subscriber.send <- []byte(codeRemove + "," + hash + "," + sector)
			}
		}
	}
}

func (s *Sector) pack(l *Lifer) {
	var pack string
	for _, lf := range l.filters {
		for member := range s.members[lf] {
			if l != member && chat(l.sa, l.filter, member.sa, member.filter) {
				pack += "," + member.hash + "," + member.inboundLat + "," + member.inboundLon + "," + member.mark
			}
		}
	}
	if pack != "" {
		l.send <- []byte(codeSectorPackage + pack)
	}
}
