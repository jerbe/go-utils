package goutils

import (
	"sync/atomic"
	"testing"
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/6 13:20
  @describe :
*/

var defaultLimiter = NewLimiter(1<<2, time.Nanosecond)

func BenchmarkLimiter_Allow(b *testing.B) {
	b.SetParallelism(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c := atomic.LoadInt64(&defaultLimiter.count)
			allow := defaultLimiter.Allow()
			b.Logf("ts:%v, allow:%v, l.count:%d", time.Now(), allow, c)
		}
	})
}
