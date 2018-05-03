package cluster_controller

import (
	"time"
)

// panicTimer panics when it reaches the given duration.
type panicTimer struct {
	d   time.Duration
	msg string
	t   *time.Timer
}

func newPanicTimer(d time.Duration, msg string) *panicTimer {
	return &panicTimer{
		d:   d,
		msg: msg,
	}
}

func (pt *panicTimer) start() {
	pt.t = time.AfterFunc(pt.d, func() {
		panic(pt.msg)
	})
}

// stop stops the timer and resets the elapsed duration.
func (pt *panicTimer) stop() {
	if pt.t != nil {
		pt.t.Stop()
		pt.t = nil
	}
}
