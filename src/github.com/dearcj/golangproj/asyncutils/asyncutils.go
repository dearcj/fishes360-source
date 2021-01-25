package asyncutils

import "time"

type AsyncCallbackFunction func()

type AsyncCallback struct {
	f         AsyncCallbackFunction
	delay     time.Duration
	startTime time.Time
}

type AsyncUtil struct {
	callbacks   []AsyncCallback
	time        *time.Time
	paused      bool
	pausedTime  time.Duration
	pausedStart time.Time
}

func (a *AsyncUtil) Pause() {
	if !a.paused {
		a.paused = true
		a.pausedStart = time.Now()
	}
}

func (a *AsyncUtil) Unpause() {
	a.paused = false
	a.pausedTime = a.pausedTime + time.Now().Sub(a.pausedStart)
}

func (a *AsyncUtil) Update() {
	if a.paused {
		return
	}

	for x := len(a.callbacks) - 1; x >= 0; x-- {
		cb := a.callbacks[x]
		t := *a.time
		if t.Sub(cb.startTime) > cb.delay {
			cb.f()
			a.callbacks = append(a.callbacks[:x], a.callbacks[x+1:]...)
		}
	}
}

func (a *AsyncUtil) DelayedCall(f AsyncCallbackFunction, delay time.Duration) {
	a.callbacks = append(a.callbacks, AsyncCallback{f: f, delay: delay, startTime: *a.time})
}

func (util *AsyncUtil) Exec() {
	cblen := len(util.callbacks) - 1
	for x := cblen; x >= 0; x-- {
		cb := util.callbacks[x]
		cb.f()
	}

	util.callbacks = nil
}

func CreateAsycnUtil(time *time.Time) *AsyncUtil {
	return &AsyncUtil{
		callbacks: []AsyncCallback{},
		time:      time,
	}
}
