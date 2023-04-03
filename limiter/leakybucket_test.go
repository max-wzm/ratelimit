package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLeakyBucketLimiter(t *testing.T) {
	limiter := NewLeakyBucketLimiter(5, 5)
	time.Sleep(900 * time.Millisecond)
	for i := 0; i < 50; i++ {
		time.Sleep(900 * time.Millisecond)
		for j := 0; j < 10; j++ {
			if limiter.Accessible() {
				fmt.Println(time.Now(), "haharr")
			}
		}
	}
}
