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
	lifer *Lifer // lifer for unsubscibe in sector
}

func (j *jobUnsubscribe) code() int {
	return 1
}

func createUnsubscribeJob(l *Lifer) *jobUnsubscribe {
	return &jobUnsubscribe{
		lifer: l,
	}
}

//
