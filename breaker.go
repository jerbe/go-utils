package goutils

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/6 16:12
  @describe :
*/
var BreakerOpen = errors.New("breaker opened")

// BreakerState 断路器状态
type BreakerState string

const (
	// BreakerStateOpen 打开状态,说明服务不能用
	BreakerStateOpen BreakerState = "BreakerStateOpen"

	// BreakerStateClose 关闭状态,说明服务可用
	BreakerStateClose = "BreakerStateClose"

	// BreakerStateHalfOpen 半开状态,说明已经经过"打开状态"
	BreakerStateHalfOpen = "BreakerStateHalfOpen"
)

func NewBreaker(successThreshold, failureThreshold int, cycleDuration time.Duration) *Breaker {
	return &Breaker{
		rand:             rand.New(rand.NewSource(time.Now().UnixNano())),
		state:            BreakerStateClose,
		successThreshold: successThreshold,
		failureThreshold: failureThreshold,
		failureCount:     0,
		successCount:     0,
		cycleDuration:    cycleDuration,
		cycleStartTime:   time.Time{},
		stateChan:        make(chan BreakerState),
	}
}

// Breaker 断路器
type Breaker struct {
	mu sync.Mutex

	rand *rand.Rand

	// state 状态值
	state BreakerState

	// successThreshold 成功次数阈值
	// 当连续成功次数达到这个阈值的时候,直接闭合上断路器,使链路联通
	successThreshold int

	// failureThreshold 失败参数的阈值
	// 当连续失败次数达到这个阈值的时候,直接打开断路器,使链路进入到打开
	failureThreshold int

	// failureCount 累计失败次数
	failureCount int

	// successCount 累计成功参数
	successCount int

	// cycleDuration 周期时长
	cycleDuration time.Duration

	// cycleStartTime 周期开始时间
	cycleStartTime time.Time

	// stateChan 状态通道,用于通知
	stateChan chan BreakerState

	// stateChanRead
	stateChanRead bool
}

// BreakerRunFunc 断路器执行方法
type BreakerRunFunc func() error

// Run 直接执行方法,根据fn返回的error判断是否为nil判断是成功还是失败
func (b *Breaker) Run(fn BreakerRunFunc) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.beforeCheck()
	// 如果断路器是断开状态,直接返回断路器已经断开
	// 直到下一个周期开始重新判断
	if b.state == BreakerStateOpen {
		return BreakerOpen
	}

	// 如果是半开状态,在当前周期内用约等于一半的概率再次允许执行,用于判断断路器是否可以闭上
	if b.state == BreakerStateHalfOpen {
		y := (1 - b.rand.Int63n(2)) == 1
		if y {
			err := fn()
			b.afterCheck(err == nil)
			return err
		}
		return BreakerOpen
	}

	// 剩下的就是闭合状态了,
	err := fn()
	b.afterCheck(err == nil)
	return err
}

func (b *Breaker) State() <-chan BreakerState {
	b.stateChanRead = true
	return b.stateChan
}

func (b *Breaker) setState(state BreakerState) {
	if b.state == state {
		return
	}

	b.state = state
	if b.stateChanRead {
		b.stateChan <- state
	}
}

func (b *Breaker) beforeCheck() {
	switch b.state {
	case BreakerStateOpen:
		// 如果周期内还是持续打开状态,则设置成半开状态,用于再次探测是否服务已经可用
		if b.cycleStartTime.Add(b.cycleDuration).Before(time.Now()) {
			b.reset()
			b.setState(BreakerStateHalfOpen)
		}
	case BreakerStateClose:
		// 如果是闭合状态,则判断是否已经是新周期,超过周期则重新开始一个新的周期
		if b.cycleStartTime.Add(b.cycleDuration).Before(time.Now()) {
			b.reset()
		}
	}
}

func (b *Breaker) afterCheck(pass bool) {
	if pass {
		b.failureCount = 0
		// 如果是半开状态并且持续成功次数达到成功阈值,则直接闭上,让链路可通
		// 只有半开状态下,成功次数累加才有意义
		if b.state == BreakerStateHalfOpen && AtomicAdd(&b.successCount, 1).(int) >= b.successThreshold {
			b.reset()
			b.setState(BreakerStateClose)
			return
		}
	} else {
		b.successCount = 0
		b.failureCount++
		// 如果当前状态是半开状态或者是闭合状态但是失败次数达到阈值,则直接打开闭合器
		// 并重新一个周期
		//@TODO 半开状态是否需要重新优化
		if b.state == BreakerStateHalfOpen || (b.state == BreakerStateClose && b.failureCount >= b.failureThreshold) {
			b.reset()
			b.setState(BreakerStateOpen)
		}
	}
}

func (b *Breaker) reset() {
	b.failureCount = 0
	b.successCount = 0
	b.cycleStartTime = time.Now()
}
