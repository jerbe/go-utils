package goutils

import (
	"errors"
	"strconv"
	"sync/atomic"
	"unsafe"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/8/18 00:35
  @describe :
*/

func IntLen(val int) int {
	return Int64Len(int64(val))
}

func Int8Len(val int8) int {
	return Int64Len(int64(val))
}

func Int16Len(val int16) int {
	return Int64Len(int64(val))
}

func Int32Len(val int32) int {
	return Int64Len(int64(val))
}

func Int64Len(val int64) int {
	return len(strconv.FormatInt(val, 10))
}

//----------------------------------------------------------
//------------------- SORT ---------------------------------
//----------------------------------------------------------

// IntSort 排序整型
// 返回的数据总是 a <= b
func IntSort(a, b int) (int, int) {
	if a > b {
		a, b = b, a
	}
	return a, b
}

// Int8Sort 排序整型
// 返回的数据总是 a <= b
func Int8Sort(a, b int8) (int8, int8) {
	if a > b {
		a, b = b, a
	}
	return a, b
}

// Int16Sort 排序整型
// 返回的数据总是 a <= b
func Int16Sort(a, b int16) (int16, int16) {
	if a > b {
		a, b = b, a
	}
	return a, b
}

// Int32Sort 排序整型
// 返回的数据总是 a <= b
func Int32Sort(a, b int32) (int32, int32) {
	if a > b {
		a, b = b, a
	}
	return a, b
}

// Int64Sort 排序整型
// 返回的数据总是 a <= b
func Int64Sort(a, b int64) (int64, int64) {
	if a > b {
		a, b = b, a
	}
	return a, b
}

// IntBetween dest在指定范围内
// 是一个闭区间 from <= dest <= to
func IntBetween(dest, from, to int) bool {
	if from <= dest && dest <= to {
		return true
	}
	return false
}

// Int8Between dest在指定范围内
// 是一个闭区间 from <= dest <= to
func Int8Between(dest, from, to int8) bool {
	if from <= dest && dest <= to {
		return true
	}
	return false
}

// Int16Between dest在指定范围内
// 是一个闭区间 from <= dest <= to
func Int16Between(dest, from, to int16) bool {
	if from <= dest && dest <= to {
		return true
	}
	return false
}

// Int32Between dest在指定范围内
// 是一个闭区间 from <= dest <= to
func Int32Between(dest, from, to int32) bool {
	if from <= dest && dest <= to {
		return true
	}
	return false
}

// Int64Between dest在指定范围内
// 是一个闭区间 from <= dest <= to
func Int64Between(dest, from, to int64) bool {
	if from <= dest && dest <= to {
		return true
	}
	return false
}

//----------------------------------------------------------
//----------------- ATOMIC ---------------------------------
//----------------------------------------------------------

// IntBitSet 填充设置某一位的数值为1,从右到左方向
func IntBitSet(val interface{}, position int) interface{} {
	switch v := val.(type) {
	case int:
		return v | (1 << position)
	case uint:
		return v | (1 << position)
	case int8:
		return v | (1 << position)
	case uint8:
		return v | (1 << position)
	case int16:
		return v | (1 << position)
	case uint16:
		return v | (1 << position)
	case int32:
		return v | (1 << position)
	case uint32:
		return v | (1 << position)
	case int64:
		return v | (1 << position)
	case uint64:
		return v | (1 << position)

	case *int:
		return *v | (1 << position)
	case *uint:
		return *v | (1 << position)
	case *int8:
		return *v | (1 << position)
	case *uint8:
		return *v | (1 << position)
	case *int16:
		return *v | (1 << position)
	case *uint16:
		return *v | (1 << position)
	case *int32:
		return *v | (1 << position)
	case *uint32:
		return *v | (1 << position)
	case *int64:
		return *v | (1 << position)
	case *uint64:
		return *v | (1 << position)
	default:
		panic("no int")
	}
}

// IntBitClear 清理设置某一位的数值为0,从右到左方向
func IntBitClear(val interface{}, position int) interface{} {
	switch v := val.(type) {
	case int:
		return v &^ (1 << position)
	case uint:
		return v &^ (1 << position)
	case int8:
		return v &^ (1 << position)
	case uint8:
		return v &^ (1 << position)
	case int16:
		return v &^ (1 << position)
	case uint16:
		return v &^ (1 << position)
	case int32:
		return v &^ (1 << position)
	case uint32:
		return v &^ (1 << position)
	case int64:
		return v &^ (1 << position)
	case uint64:
		return v &^ (1 << position)

	case *int:
		return *v &^ (1 << position)
	case *uint:
		return *v &^ (1 << position)
	case *int8:
		return *v &^ (1 << position)
	case *uint8:
		return *v &^ (1 << position)
	case *int16:
		return *v &^ (1 << position)
	case *uint16:
		return *v &^ (1 << position)
	case *int32:
		return *v &^ (1 << position)
	case *uint32:
		return *v &^ (1 << position)
	case *int64:
		return *v &^ (1 << position)
	case *uint64:
		return *v &^ (1 << position)
	default:
		panic("no int")
	}
}

// IntBitCheck 检测某一位的数值是否为1,从右到左方向
func IntBitCheck(val interface{}, position int) bool {
	switch v := val.(type) {
	case int:
		return v&(1<<position) != 0
	case uint:
		return v&(1<<position) != 0
	case int8:
		return v&(1<<position) != 0
	case uint8:
		return v&(1<<position) != 0
	case int16:
		return v&(1<<position) != 0
	case uint16:
		return v&(1<<position) != 0
	case int32:
		return v&(1<<position) != 0
	case uint32:
		return v&(1<<position) != 0
	case int64:
		return v&(1<<position) != 0
	case uint64:
		return v&(1<<position) != 0

	case *int:
		return *v&(1<<position) != 0
	case *uint:
		return *v&(1<<position) != 0
	case *int8:
		return *v&(1<<position) != 0
	case *uint8:
		return *v&(1<<position) != 0
	case *int16:
		return *v&(1<<position) != 0
	case *uint16:
		return *v&(1<<position) != 0
	case *int32:
		return *v&(1<<position) != 0
	case *uint32:
		return *v&(1<<position) != 0
	case *int64:
		return *v&(1<<position) != 0
	case *uint64:
		return *v&(1<<position) != 0
	default:
		panic("no int")
	}
}

//----------------------------------------------------------
//----------------- ATOMIC ---------------------------------
//----------------------------------------------------------

// AtomicAddInt 原子增加
func AtomicAddInt(dst interface{}, delta int64) interface{} {
	switch d := dst.(type) {
	case *int64:
		return atomic.AddInt64(d, int64(delta))
	case *int32:
		return atomic.AddInt32(d, int32(delta))
	case *int16:
		n := (*int32)(unsafe.Pointer(d))
		return int16(atomic.AddInt32(n, int32(delta)))
	case *int8:
		n := (*int32)(unsafe.Pointer(d))
		return int8(atomic.AddInt32(n, int32(delta)))
	case *int:
		n := (*int32)(unsafe.Pointer(d))
		return int(atomic.AddInt32(n, int32(delta)))
	case *uint:
		n := (*uint32)(unsafe.Pointer(d))
		return uint(atomic.AddUint32(n, uint32(delta)))
	case *uint8:
		n := (*uint32)(unsafe.Pointer(d))
		return uint8(atomic.AddUint32(n, uint32(delta)))
	case *uint16:
		n := (*uint32)(unsafe.Pointer(d))
		return uint16(atomic.AddUint32(n, uint32(delta)))
	case *uint32:
		return atomic.AddUint32(d, uint32(delta))
	case *uint64:
		return atomic.AddUint64(d, uint64(delta))
	default:
		panic(errors.New("unsupported type"))
	}
}

// AtomicStoreInt 原子存储
func AtomicStoreInt(dst interface{}, val int64) {

	switch d := dst.(type) {
	case *int64:
		atomic.StoreInt64(d, int64(val))
	case *int32:
		atomic.StoreInt32(d, int32(val))
	case *int16:
		n := (*int32)(unsafe.Pointer(d))
		atomic.StoreInt32(n, int32(val))
	case *int8:
		n := (*int32)(unsafe.Pointer(d))
		atomic.StoreInt32(n, int32(val))
	case *int:
		n := (*int32)(unsafe.Pointer(d))
		atomic.StoreInt32(n, int32(val))
	case *uint:
		n := (*uint32)(unsafe.Pointer(d))
		atomic.StoreUint32(n, uint32(val))
	case *uint8:
		n := (*uint32)(unsafe.Pointer(d))
		atomic.StoreUint32(n, uint32(val))
	case *uint16:
		n := (*uint32)(unsafe.Pointer(d))
		atomic.StoreUint32(n, uint32(val))
	case *uint32:
		atomic.StoreUint32(d, uint32(val))
	case *uint64:
		atomic.StoreUint64(d, uint64(val))
	default:
		panic(errors.New("unsupported type"))
	}
}

// AtomicLoadInt 原子加载
func AtomicLoadInt(dst interface{}) interface{} {
	switch d := dst.(type) {
	case *int64:
		return atomic.LoadInt64(d)
	case *int32:
		return atomic.LoadInt32(d)
	case *int16:
		n := (*int32)(unsafe.Pointer(d))
		return int16(atomic.LoadInt32(n))
	case *int8:
		n := (*int32)(unsafe.Pointer(d))
		return int8(atomic.LoadInt32(n))
	case *int:
		n := (*int32)(unsafe.Pointer(d))
		return int(atomic.LoadInt32(n))
	case *uint:
		n := (*uint32)(unsafe.Pointer(d))
		return uint(atomic.LoadUint32(n))
	case *uint8:
		n := (*uint32)(unsafe.Pointer(d))
		return uint8(atomic.LoadUint32(n))
	case *uint16:
		n := (*uint32)(unsafe.Pointer(d))
		return uint16(atomic.LoadUint32(n))
	case *uint32:
		return atomic.LoadUint32(d)
	case *uint64:
		return atomic.LoadUint64(d)
	default:
		panic(errors.New("unsupported type"))
	}
}

// AtomicCompareAndSwapInt 原子对比并替换
func AtomicCompareAndSwapInt(dst interface{}, oldVal, newVal interface{}) bool {
	switch d := dst.(type) {
	case *int64:
		o, n := oldVal.(int64), newVal.(int64)
		return atomic.CompareAndSwapInt64(d, o, n)
	case *int32:
		o, n := oldVal.(int32), newVal.(int32)
		return atomic.CompareAndSwapInt32(d, o, n)
	case *int16:
		o, n, nd := int32(oldVal.(int16)), int32(newVal.(int16)), (*int32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapInt32(nd, o, n)
	case *int8:
		o, n, nd := int32(oldVal.(int8)), int32(newVal.(int8)), (*int32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapInt32(nd, o, n)
	case *int:
		o, n, nd := int32(oldVal.(int)), int32(newVal.(int)), (*int32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapInt32(nd, o, n)
	case *uint:
		o, n, nd := uint32(oldVal.(uint)), uint32(newVal.(uint)), (*uint32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapUint32(nd, o, n)
	case *uint8:
		o, n, nd := uint32(oldVal.(uint8)), uint32(newVal.(uint8)), (*uint32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapUint32(nd, o, n)
	case *uint16:
		o, n, nd := uint32(oldVal.(uint16)), uint32(newVal.(uint16)), (*uint32)(unsafe.Pointer(d))
		return atomic.CompareAndSwapUint32(nd, o, n)
	case *uint32:
		o, n := oldVal.(uint32), newVal.(uint32)
		return atomic.CompareAndSwapUint32(d, o, n)
	case *uint64:
		o, n := oldVal.(uint64), newVal.(uint64)
		return atomic.CompareAndSwapUint64(d, o, n)
	default:
		panic(errors.New("unsupported type"))
	}
}

// AtomicSwapInt 原子替换
func AtomicSwapInt(dst interface{}, newVal interface{}) interface{} {
	switch d := dst.(type) {
	case *int64:
		n := newVal.(int64)
		return atomic.SwapInt64(d, n)
	case *int32:
		n := newVal.(int32)
		return atomic.SwapInt32(d, n)
	case *int16:
		n, nd := int32(newVal.(int16)), (*int32)(unsafe.Pointer(d))
		return atomic.SwapInt32(nd, n)
	case *int8:
		n, nd := int32(newVal.(int8)), (*int32)(unsafe.Pointer(d))
		return atomic.SwapInt32(nd, n)
	case *int:
		n, nd := int32(newVal.(int)), (*int32)(unsafe.Pointer(d))
		return atomic.SwapInt32(nd, n)
	case *uint:
		n, nd := uint32(newVal.(uint)), (*uint32)(unsafe.Pointer(d))
		return atomic.SwapUint32(nd, n)
	case *uint8:
		n, nd := uint32(newVal.(uint8)), (*uint32)(unsafe.Pointer(d))
		return atomic.SwapUint32(nd, n)
	case *uint16:
		n, nd := uint32(newVal.(uint16)), (*uint32)(unsafe.Pointer(d))
		return atomic.SwapUint32(nd, n)
	case *uint32:
		n := newVal.(uint32)
		return atomic.SwapUint32(d, n)
	case *uint64:
		n := newVal.(uint64)
		return atomic.SwapUint64(d, n)
	default:
		panic(errors.New("unsupported type"))
	}
}
