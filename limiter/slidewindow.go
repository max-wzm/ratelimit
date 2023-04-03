package limiter

import (
	"sync"
	"time"
)

type windowsQ struct {
	windows     []int
	front, rear int
	windowCnt   int
}

func NewWindowQ(windowCnt int) *windowsQ {
	q := &windowsQ{
		windows:   make([]int, windowCnt),
		front:     0,
		rear:      windowCnt - 1,
		windowCnt: windowCnt,
	}
	return q
}

func (q *windowsQ) DeQ() int {
	out := q.windows[q.front]
	q.front = (q.front + 1) % q.windowCnt
	return out
}

func (q *windowsQ) EnQ(w int) {
	q.rear = (q.rear + 1) % q.windowCnt
	q.windows[q.rear] = w
}

type SlideWindowLimiter struct {
	maxQPS       int
	interval     int
	windowsQ     *windowsQ
	totalCounter int
	sync.Mutex
}

func NewSlideWindowLimiter(maxQPS int, windowCnt int) *SlideWindowLimiter {
	limiter := &SlideWindowLimiter{
		maxQPS:       maxQPS,
		interval:     int(time.Second),
		windowsQ:     NewWindowQ(windowCnt),
		totalCounter: 0,
	}
	go limiter.Slide()
	return limiter
}

func (limiter *SlideWindowLimiter) Slide() {
	for {
		time.Sleep(time.Second / time.Duration(limiter.windowsQ.windowCnt))
		limiter.Lock()
		w := limiter.windowsQ.DeQ()
		limiter.totalCounter -= w
		limiter.windowsQ.EnQ(0)
		limiter.Unlock()
	}
}

func (limiter *SlideWindowLimiter) Accessible() bool {
	limiter.Lock()
	defer limiter.Unlock()
	if limiter.totalCounter < limiter.maxQPS {
		limiter.windowsQ.windows[limiter.windowsQ.rear]++
		limiter.totalCounter++
		return true
	} else {
		return false
	}
}
