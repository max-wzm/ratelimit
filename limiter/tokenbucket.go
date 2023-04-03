package limiter

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	rate       int // max qps per second
	capacity   int
	token      int
	lastRefill int64
	sync.Mutex
}

func NewTokenBucketLimiter(rate int, capacity int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		rate:       rate,
		capacity:   capacity,
		token:      capacity,
		lastRefill: time.Now().UnixNano(),
	}
}

func (limiter *TokenBucketLimiter) Accessible() bool {
	now := time.Now().UnixNano()
	limiter.refill(now)
	if limiter.token > 0 {
		limiter.token--
		limiter.lastRefill = now
		return true
	}
	return false
}

func (limiter *TokenBucketLimiter) refill(now int64) {
	passed := now - limiter.lastRefill
	genTokens := passed * int64(limiter.rate) / time.Second.Nanoseconds()
	limiter.token = min(limiter.capacity, limiter.token+int(genTokens))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
