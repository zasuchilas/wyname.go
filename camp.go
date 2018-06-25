package main

import (
	"sync"
)

// Camp for all lifes
type Camp struct {
	sectors map[string]*Sector // ссылки на все секторы
	mutex   *sync.RWMutex      // мьютекс к ресурсу sectors
}

func newcamp() *Camp {
	return &Camp{
		sectors: make(map[string]*Sector, 1000000),
		mutex:   new(sync.RWMutex),
	}
}

// sector возвращает указатель на сектор по его имени
// создает новый сектор, если в sectors отсутствует имя
func (c *Camp) sector(name string) *Sector {
	c.mutex.RLock()
	s, found := c.sectors[name]
	c.mutex.RUnlock()
	if found == false {
		s = newsector()
		c.mutex.Lock()
		s2, found := c.sectors[name]
		if found == false {
			c.sectors[name] = s
			s.run() // run new sector
		} else {
			s = s2
		}
		c.mutex.Unlock()
	}
	return s
}
