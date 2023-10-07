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

// IntBitSet 填充设置某一位的数值为1,从右到左方向
func IntBitSet(val int, position int) int {
	return val | (1 << position)
}

// Int16BitSet 填充设置某一位的数值为1,从右到左方向
func Int16BitSet(val int16, position int) int16 {
	return val | (1 << position)
}

// Int32BitSet 填充设置某一位的数值为1,从右到左方向
func Int32BitSet(val int32, position int) int32 {
	return val | (1 << position)
}

// Int64BitSet 填充设置某一位的数值为1,从右到左方向
func Int64BitSet(val int64, position int) int64 {
	return val | (1 << position)
}

// IntBitClear 清理设置某一位的数值为0,从右到左方向
func IntBitClear(val int, position int) int {
	return val &^ (1 << position)
}

// Int8BitClear 清理设置某一位的数值为0,从右到左方向
func Int8BitClear(val int8, position int) int8 {
	return val &^ (1 << position)
}

// Int16BitClear 清理设置某一位的数值为0,从右到左方向
func Int16BitClear(val int16, position int) int16 {
	return val &^ (1 << position)
}

// Int32BitClear 清理设置某一位的数值为0,从右到左方向
func Int32BitClear(val int32, position int) int32 {
	return val &^ (1 << position)
}

// Int64BitClear 清理设置某一位的数值为0,从右到左方向
func Int64BitClear(val int64, position int) int64 {
	return val &^ (1 << position)
}

// AtomicAdd 原子增加
func AtomicAdd(dst interface{}, delta int) interface{} {
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

// AtomicStore 原子存储
func AtomicStore(dst interface{}, val int) {
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

// AtomicLoad 原子加载
func AtomicLoad(dst interface{}) interface{} {
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
