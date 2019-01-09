package globals

import "sync"

type GlobalStates struct {
	RunningQueries    uint64
	TotalQueries      uint64
	FailedQueries     uint64
	SuccessfulQueries uint64

	RQMutex sync.Mutex `xml:"-" json:"-"`
	TQMutex sync.Mutex `xml:"-" json:"-"`
	FQMutex sync.Mutex `xml:"-" json:"-"`
	SQMutex sync.Mutex `xml:"-" json:"-"`
}

func (a *GlobalStates) IncrementRQ() {
	a.RQMutex.Lock()
	a.RunningQueries++
	a.RQMutex.Unlock()
}

func (a *GlobalStates) DecrementRQ() {
	a.RQMutex.Lock()
	a.RunningQueries--
	a.RQMutex.Unlock()
}

func (a *GlobalStates) IncrementTQ() {
	a.TQMutex.Lock()
	a.TotalQueries++
	a.TQMutex.Unlock()
}

func (a *GlobalStates) DecrementTQ() {
	a.TQMutex.Lock()
	a.TotalQueries--
	a.TQMutex.Unlock()
}

func (a *GlobalStates) IncrementFQ() {
	a.FQMutex.Lock()
	a.FailedQueries++
	a.FQMutex.Unlock()
}

func (a *GlobalStates) DecrementFQ() {
	a.FQMutex.Lock()
	a.FailedQueries--
	a.FQMutex.Unlock()
}

func (a *GlobalStates) IncrementSQ() {
	a.SQMutex.Lock()
	a.SuccessfulQueries++
	a.SQMutex.Unlock()
}

func (a *GlobalStates) DecrementSQ() {
	a.SQMutex.Lock()
	a.SuccessfulQueries--
	a.SQMutex.Unlock()
}
