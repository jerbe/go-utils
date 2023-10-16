package goutils

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/16 09:29
  @describe :
*/

func TestNewBloomFilter(t *testing.T) {

	var start runtime.MemStats
	runtime.ReadMemStats(&start)
	var (
		name1 = "name1"
		name2 = "name2"
		name3 = "name3"
		name4 = "name4"
		name5 = "name5"
		name6 = "name6"
		name7 = "name7"
	)

	b := NewBloomFilter(1<<20, 3)

	//t.Logf("BEFORE: => b.bitmap:%v", b.bitmap)
	b.Set(name1)
	//b.Set(name2)
	//b.Set(name3)
	//b.Set(name4)
	//b.Set(name5)
	//b.Set(name6)
	b.Set(name7)
	//t.Logf("AFTER: => b.bitmap:%v", b.bitmap)

	t.Logf("1:%v", b.Contain(name1))
	t.Logf("2:%v", b.Contain(name2))
	t.Logf("3:%v", b.Contain(name3))
	t.Logf("4:%v", b.Contain(name4))
	t.Logf("5:%v", b.Contain(name5))
	t.Logf("6:%v", b.Contain(name6))
	t.Logf("7:%v", b.Contain(name7))

	var bloomMemStats runtime.MemStats
	runtime.ReadMemStats(&bloomMemStats)
	t.Logf("布隆消耗了:%d MiB内存", (bloomMemStats.TotalAlloc-start.TotalAlloc)>>20)
	cnt := 1000000
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			x := strconv.Itoa(i)
			b.Set(x)
			if !b.Contain(x) {
				t.Logf("x:%v", x)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	var stop runtime.MemStats
	runtime.ReadMemStats(&stop)

	t.Logf("这些代码消耗了:%d MiB内存", (stop.Alloc-start.Alloc)>>20)
	//t.Logf("AFTER2: => len:%v, b.bitmap:%v", len(b.bitmap), b.bitmap)
}
