package goutils

import "strconv"

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
