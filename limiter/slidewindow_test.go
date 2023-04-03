package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestSlideWindowLimiter(t *testing.T) {
	limiter := NewSlideWindowLimiter(5, 5)
	time.Sleep(400 * time.Millisecond)
	for i := 0; i < 500; i++ {
		time.Sleep(10 * time.Millisecond)
		if limiter.Accessible() {
			fmt.Println(time.Now(), "haha")
		}
	}
}
