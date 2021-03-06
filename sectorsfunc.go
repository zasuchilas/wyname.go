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
			if hash != subscriber.hash && chat(sa, filter, ssa, sfilter) {
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
			if hash != subscriber.hash && chat(sa, filter, ssa, sfilter) {
				subscriber.send <- []byte(codeRemove + "," + hash + "," + sector)
			}
		}
	}
}

func (s *Sector) pack(l *Lifer, sa, filter int, filters []int) {
	pack, err := s.sectorPack(l, sa, filter, filters)
	if err == nil {
		l.send <- []byte(codeSectorPackage + pack)
	}
}

func (s *Sector) glob(l *Lifer, sa, filter int, filters []int, globReqCode string) {
	pack, err := s.sectorPack(l, sa, filter, filters)
	if err == nil {
		l.send <- []byte(codeGlobPackage + "," + globReqCode + pack)
	}
}

func (s *Sector) sectorPack(l *Lifer, sa, filter int, filters []int) (pack string, err error) {
	hash := l.hash
	for _, lf := range filters {
		for member := range s.members[lf] {
			member.mutex.RLock()
			msa := member.sa
			mfilter := member.filter
			mlat := member.inboundLat
			mlon := member.inboundLon
			mmark := member.mark
			member.mutex.RUnlock()
			if hash != member.hash && chat(sa, filter, msa, mfilter) {
				pack += "," + member.hash + "," + mlat + "," + mlon + "," + mmark
			}
		}
	}
	if pack == "" {
		err = fmt.Errorf("package empty")
	}
	return
}
