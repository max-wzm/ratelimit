package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestFixedWindowLimiter(t *testing.T) {
	limiter := NewFixedWindowLimiter(5)
	time.Sleep(900 * time.Millisecond)
	for i := 0; i < 200; i++ {
		time.Sleep(10 * time.Millisecond)
		if limiter.Accessible() {
			fmt.Println(time.Now(), "haha")
		}
	}
}
