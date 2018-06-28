package main

type job interface {
	code() int
}

// come

type jobCome struct {
	lifer *Lifer // lifer for add to sector members
}

func (j *jobCome) code() int {
	return 1
}

func createComeJob(l *Lifer) *jobCome {
	return &jobCome{
		lifer: l,
	}
}

// move

type jobMove struct {
	lifer *Lifer // lifer for add to sector members
}

func (j *jobMove) code() int {
	return 1
}

func createMoveJob(l *Lifer) *jobMove {
	return &jobMove{
		lifer: l,
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

func createAwayJob(l *Lifer, sa int, filter int, f []int) *jobAway {
	return &jobAway{
		lifer:   l,
		sa:      sa,
		filter:  filter,
		filters: f,
	}
}

// subscribe

type jobSubscribe struct {
	lifer *Lifer // lifer for subsribe in sector
}

func (j *jobSubscribe) code() int {
	return 1
}

func createSubscribeJob(l *Lifer) *jobSubscribe {
	return &jobSubscribe{
		lifer: l,
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

func createUnsubscribeJob(l *Lifer, sa int, notify bool) *jobUnsubscribe {
	return &jobUnsubscribe{
		lifer:  l,
		sa:     sa,
		notify: notify,
	}
}

// glob

type jobGlob struct {
	lifer       *Lifer // the lifer who made the request
	globReqCode string // glob code of request
}

func (j *jobGlob) code() int {
	return 1
}

func createGlobJob(l *Lifer, globReqCode string) *jobGlob {
	return &jobGlob{
		lifer:       l,
		globReqCode: globReqCode,
	}
}

//
