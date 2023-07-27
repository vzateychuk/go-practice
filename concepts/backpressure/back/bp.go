package back

import "errors"

type Backpressure struct {
	ch chan struct{}
}

func NewBackpressure(limit int) *Backpressure {
	chanLimited := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		chanLimited <- struct{}{}
	}

	return &Backpressure{
		ch: chanLimited,
	}
}

func (bp *Backpressure) Process(f func()) error {
	select {
	case <-bp.ch:
		f()
		bp.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}
