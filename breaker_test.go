package goutils

import (
	"errors"
	"testing"
	"time"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/7 12:16
  @describe :
*/

func TestBreaker_Run(t *testing.T) {

	breaker := NewBreaker(5, 50, time.Second)

	go func() {
		for {
			select {
			case state := <-breaker.State():
				t.Log(state)
			}
		}
	}()

	var sig int

	go func() {
		for {
			select {
			case <-time.NewTimer(time.Second).C:
				sig = 1 - sig
			}
		}
	}()

	for {
		go func() {
			breaker.Run(func() error {
				if sig == 0 {
					return errors.New("调用失败了")
				}
				return nil
			})
		}()
	}

}
