package goutils

import (
	"sync"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 18:51
  @describe :
*/

// PubSub 发布/订阅器
type PubSub struct {
	subscribers map[*SubSignal]struct{}
	rwMute      sync.RWMutex
	closed      bool
}

// SubSignal 订阅信号
type SubSignal struct {
	ch     chan interface{}
	C      <-chan interface{}
	close  chan struct{}
	signal *PubSub
}

// Close 订阅者信号
func (sb *SubSignal) Close() error {
	sb.signal.Unsubscribe(sb)
	return nil
}

// Publish 推送
func (s *PubSub) Publish(i interface{}) {
	if s.closed {
		panic("signal is closed")
	}
	s.rwMute.RLock()
	defer s.rwMute.RUnlock()
	for sb := range s.subscribers {
		go func(in *SubSignal) {
			select {
			case <-in.close:
				close(in.ch)
				return
			case in.ch <- i:
			}
		}(sb)
	}
}

// Subscribe 订阅
func (s *PubSub) Subscribe() *SubSignal {
	if s.closed {
		panic("signal is closed")
	}
	ch := make(chan interface{})
	sb := &SubSignal{
		ch:     ch,
		C:      ch,
		close:  make(chan struct{}),
		signal: s,
	}
	s.rwMute.Lock()
	defer s.rwMute.Unlock()
	s.subscribers[sb] = struct{}{}
	return sb
}

// Unsubscribe 取消订阅
func (s *PubSub) Unsubscribe(sb *SubSignal) {
	s.rwMute.Lock()
	defer s.rwMute.Unlock()
	if _, ok := s.subscribers[sb]; ok {
		delete(s.subscribers, sb)
		close(sb.close)
	}
}

// Close 关闭该发布/订阅器
func (s *PubSub) Close() {
	if s.closed {
		return
	}
	s.rwMute.Lock()
	defer s.rwMute.Unlock()
	s.closed = true
	for sb := range s.subscribers {
		delete(s.subscribers, sb)
	}
}

// NewPubSub 新生成一个发布/订阅器
func NewPubSub() *PubSub {
	return &PubSub{subscribers: make(map[*SubSignal]struct{})}
}
