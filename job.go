package main

type job interface {
	code() int
}

// come

type comejob struct {
	lifer *Lifer // lifer for add to sector members
}

func (j *comejob) code() int {
	return 1
}

func newcomejob(l *Lifer) *comejob {
	return &comejob{
		lifer: l,
	}
}

// away

type awayjob struct {
	lifer *Lifer // lifer for remove
}

func (j *awayjob) code() int {
	return 1
}

func newawayjob(l *Lifer) *awayjob {
	return &awayjob{
		lifer: l,
	}
}

// subscribe

type jobSubscribe struct {
	lifer *Lifer // lifer for subsribe in sector
}

func (j *jobSubscribe) code() int {
	return 1
}

func newSubscribeJob(l *Lifer) *jobSubscribe {
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

func newUnsubscribeJob(l *Lifer) *jobUnsubscribe {
	return &jobUnsubscribe{
		lifer: l,
	}
}

//
