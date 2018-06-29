package main

type job interface {
	code() int
}

// come

type jobCome struct {
	lifer   *Lifer // lifer for add to sector members
	lat     string // snapshot inboundLat
	lon     string // snapshot inbounLon
	sa      int    // snapshot sa
	filter  int    // snapshot filter
	mark    string // snapshot mark
	filters []int  // snapshot filters
}

func (j *jobCome) code() int {
	return 1
}

func createComeJob(l *Lifer) *jobCome {
	l.mutex.RLock()
	lat := l.inboundLat
	lon := l.inboundLon
	sa := l.sa
	filter := l.filter
	mark := l.mark
	filters := make([]int, len(l.filters))
	copy(filters, l.filters)
	l.mutex.RUnlock()
	return &jobCome{
		lifer:   l,
		lat:     lat,
		lon:     lon,
		sa:      sa,
		filter:  filter,
		mark:    mark,
		filters: filters,
	}
}

// move

type jobMove struct {
	lifer   *Lifer // lifer for add to sector members
	lat     string // snapshot inboundLat
	lon     string // snapshot inbounLon
	sa      int    // snapshot sa
	filter  int    // snapshot filter
	mark    string // snapshot mark
	filters []int  // snapshot filters
}

func (j *jobMove) code() int {
	return 1
}

func createMoveJob(l *Lifer) *jobMove {
	l.mutex.RLock()
	lat := l.inboundLat
	lon := l.inboundLon
	sa := l.sa
	filter := l.filter
	mark := l.mark
	filters := make([]int, len(l.filters))
	copy(filters, l.filters)
	l.mutex.RUnlock()
	return &jobMove{
		lifer:   l,
		lat:     lat,
		lon:     lon,
		sa:      sa,
		filter:  filter,
		mark:    mark,
		filters: filters,
	}
}

// away

type jobAway struct {
	lifer   *Lifer // lifer for remove
	sa      int    // snapshot of lifers sa
	filter  int    // snapshot of lifers filter
	filters []int  // copy lifer.filters
}

func (j *jobAway) code() int {
	return 1
}

func createAwayJob(l *Lifer) *jobAway {
	l.mutex.RLock()
	sa := l.sa
	filter := l.filter
	filters := make([]int, len(l.filters))
	copy(filters, l.filters)
	l.mutex.RUnlock()
	return &jobAway{
		lifer:   l,
		sa:      sa,
		filter:  filter,
		filters: filters,
	}
}

// subscribe

type jobSubscribe struct {
	lifer   *Lifer // lifer for subsribe in sector
	sa      int    // snapshot of lifers sa
	filter  int    // snapshot of lifers filter
	filters []int  // copy lifer.filters
}

func (j *jobSubscribe) code() int {
	return 1
}

func createSubscribeJob(l *Lifer) *jobSubscribe {
	l.mutex.RLock()
	sa := l.sa
	filter := l.filter
	filters := make([]int, len(l.filters))
	copy(filters, l.filters)
	l.mutex.RUnlock()
	return &jobSubscribe{
		lifer:   l,
		sa:      sa,
		filter:  filter,
		filters: filters,
	}
}

// unsubscribe

type jobUnsubscribe struct {
	lifer  *Lifer // lifer for unsubscibe in sector
	sa     int    // snapshot of lifers sa
	notify bool   // notify client about unpack
}

func (j *jobUnsubscribe) code() int {
	return 1
}

func createUnsubscribeJob(l *Lifer, notify bool) *jobUnsubscribe {
	l.mutex.RLock()
	sa := l.sa
	l.mutex.RUnlock()
	return &jobUnsubscribe{
		lifer:  l,
		sa:     sa,
		notify: notify,
	}
}

// glob

type jobGlob struct {
	lifer       *Lifer // the lifer who made the request
	sa          int    // snapshot of lifers sa
	filter      int    // snapshot of lifers filter
	filters     []int  // copy lifer.filters
	globReqCode string // glob code of request
}

func (j *jobGlob) code() int {
	return 1
}

func createGlobJob(l *Lifer, globReqCode string) *jobGlob {
	l.mutex.RLock()
	sa := l.sa
	filter := l.filter
	filters := make([]int, len(l.filters))
	copy(filters, l.filters)
	l.mutex.RUnlock()
	return &jobGlob{
		lifer:       l,
		sa:          sa,
		filter:      filter,
		filters:     filters,
		globReqCode: globReqCode,
	}
}

//
