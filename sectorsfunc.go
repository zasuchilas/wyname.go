package main

import "fmt"

// move notify subscribers about move
func (s *Sector) move(l *Lifer, lat, lon, mark string, sa, filter int, filters []int) {
	hash := l.hash
	for _, lf := range filters {
		for subscriber := range s.subscrs[lf] {
			subscriber.mutex.RLock()
			ssa := subscriber.sa
			sfilter := subscriber.filter
			subscriber.mutex.RUnlock()
			if l != subscriber && chat(sa, filter, ssa, sfilter) {
				subscriber.send <- []byte(codeMove + "," + hash + "," + lat + "," + lon + "," + mark)
			}
		}
	}
}

func (s *Sector) away(l *Lifer, sa, filter int, filters []int) {
	hash := l.hash
	sector := s.name
	for _, lf := range filters {
		for subscriber := range s.subscrs[lf] {
			subscriber.mutex.RLock()
			ssa := subscriber.sa
			sfilter := subscriber.filter
			subscriber.mutex.RUnlock()
			if l != subscriber && chat(sa, filter, ssa, sfilter) {
				subscriber.send <- []byte(codeRemove + "," + hash + "," + sector)
			}
		}
	}
}

func (s *Sector) pack(l *Lifer) {
	pack, err := s.sectorPack(l)
	if err == nil {
		l.send <- []byte(codeSectorPackage + pack)
	}
}

func (s *Sector) glob(l *Lifer, globReqCode string) {
	pack, err := s.sectorPack(l)
	if err == nil {
		l.send <- []byte(codeGlobPackage + "," + globReqCode + pack)
	}
}

func (s *Sector) sectorPack(l *Lifer) (pack string, err error) {
	l.mutex.RLock()
	sa := l.sa
	filter := l.filter
	filters := l.filters
	l.mutex.RUnlock()
	for _, lf := range filters {
		for member := range s.members[lf] {
			member.mutex.RLock()
			msa := member.sa
			mfilter := member.filter
			mlat := member.inboundLat
			mlon := member.inboundLon
			mmark := member.mark
			member.mutex.RUnlock()
			if l != member && chat(sa, filter, msa, mfilter) {
				pack += "," + member.hash + "," + mlat + "," + mlon + "," + mmark
			}
		}
	}
	if pack == "" {
		err = fmt.Errorf("package empty")
	}
	return
}
