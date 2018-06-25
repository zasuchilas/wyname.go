package main

type job interface {
	code() int
}

// come

type comejob struct {
	lifer *Lifer // lifer for add to sector members
}

func (j *comejob) code() int {
	return codeCome
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
	return codeAway
}

func newawayjob(l *Lifer) *awayjob {
	return &awayjob{
		lifer: l,
	}
}

//
