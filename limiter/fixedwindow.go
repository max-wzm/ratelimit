package limiter

import (
	"sync/atomic"
	"time"
)

// Based on timestamp
// also can be based on timer
type FixedWindowLimiter struct {
	interval int64
	window   int64
	maxQPS   uint64
	counter  uint64
}

func NewFixedWindowLimiter(maxQPS uint64) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		interval: time.Second.Nanoseconds(),
		window:   time.Now().UnixNano(),
		maxQPS:   maxQPS,
		counter:  0,
	}
}

func (limiter *FixedWindowLimiter) Accessible() bool {
	now := time.Now().UnixNano()
	if now > limiter.window+limiter.interval {
		limiter.window = now
		atomic.StoreUint64(&limiter.counter, 0)
		return true
	} else {
		atomic.AddUint64(&limiter.counter, 1)
		return limiter.counter <= limiter.maxQPS
	}
}
