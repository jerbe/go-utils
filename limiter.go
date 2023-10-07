package goutils

import (
	"sync/atomic"
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/6 11:17
  @describe :
*/

// Limiter 数量限定器
type Limiter struct {
	// last 最后一次时间
	last time.Time

	// rate 速率
	rate time.Duration

	// count 剩余多少数量
	count int64

	// size 可容纳数量
	size int64
}

// getFillCount 获取需要填充数量
func (l *Limiter) getFillCount() int64 {
	c := atomic.LoadInt64(&l.count)
	if c >= l.size {
		return 0
	}

	if !l.last.IsZero() {
		// 获取得出需要填充的数量
		cnt := int64(time.Now().Sub(l.last) / l.rate)

		// 限定填充数量不能超过限定总大小
		if l.size-c >= cnt {
			return cnt
		} else {
			return l.size - l.count
		}
	}
	return l.size
}

// fill 填充数量
func (l *Limiter) fill() {
	atomic.AddInt64(&l.count, l.getFillCount())
}

// Allow 判断是否允许通过
func (l *Limiter) Allow() bool {
	l.fill()

	if atomic.LoadInt64(&l.count) > 0 {
		atomic.AddInt64(&l.count, -1)
		l.last = time.Now()
		return true
	}
	return false
}

// SetSize 设置大小
func (l *Limiter) SetSize(val int64) {
	l.size = val
}

// NewLimiter 返回新的限制器
func NewLimiter(size int64, rate time.Duration) *Limiter {
	return &Limiter{size: size, rate: rate, count: size}
}
