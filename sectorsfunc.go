package main

// move notify subscribers about move
func (s *Sector) move(l *Lifer) {
	hash := l.hash
	lat := l.inboundLat
	lon := l.inboundLon
	mark := l.mark
	for _, lifersa := range l.filters {
		for subscriber := range s.subscrs[lifersa] {
			if l != subscriber && chat(l.sa, l.filter, subscriber.sa, subscriber.filter) {
				subscriber.send <- []byte(codeMove + "," + hash + "," + lat + "," + lon + "," + mark)
			}
		}
	}
}
