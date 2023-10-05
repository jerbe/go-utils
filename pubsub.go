package goutils

import "sync"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 18:51
  @describe :
*/

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[chan int]interface{}
	closed      bool
}

func (s *PubSub) Publish(i int) {
	defer func() {
		if obj := recover(); obj != nil {
		}
	}()
	if s.closed {
		panic("pubSub is closed")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	for ch := range s.subscribers {
		go func(in chan int) {

			in <- i
		}(ch)
	}
}

func (s *PubSub) Subscribe() <-chan int {
	if s.closed {
		panic("pubSub is closed")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	output := make(chan int)
	s.subscribers[output] = nil
	return output
}

func (s *PubSub) Close() {
	if s.closed {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for ch := range s.subscribers {
		delete(s.subscribers, ch)
		close(ch)
	}
}

func NewPubSub() *PubSub {
	return &PubSub{subscribers: make(map[chan int]interface{})}
}
