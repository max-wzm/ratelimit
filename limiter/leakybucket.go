package limiter

import (
	"time"
)

type Event int

type LeakyBucketLimiter struct {
	capacity int
	rate     int // max qps per second
	eventQ   chan Event
	accChan  chan bool
}

func NewLeakyBucketLimiter(capacity int, rate int) *LeakyBucketLimiter {
	l := &LeakyBucketLimiter{
		capacity: capacity,
		rate:     rate,
		eventQ:   make(chan Event, capacity),
		accChan:  make(chan bool),
	}
	go l.leakOut()
	return l
}

func (limiter *LeakyBucketLimiter) leakOut() {
	for {
		time.Sleep(time.Second / time.Duration(limiter.rate))
		b := false
		if len(limiter.eventQ) != 0 {
			b = true
			<-limiter.eventQ
		}
		limiter.accChan <- b
	}
}

func (limiter *LeakyBucketLimiter) Accessible() bool {
	if len(limiter.eventQ) < limiter.capacity {
		limiter.eventQ <- Event(0)
	}
	return <-limiter.accChan
}
