package goutils

import (
	"errors"
	"reflect"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 18:49
  @describe :
*/

type CompareFunc func(left, right interface{}) bool

// LT 小于
func LT(left, right interface{}) bool {
	if IsNil(left) || IsNil(right) {
		return false
	}

	if !TypeEqual(left, right) {
		return false
	}

	switch v := left.(type) {
	case int:
		return v < right.(int)
	case int8:
		return v < right.(int8)
	case int16:
		return v < right.(int16)
	case int32:
		return v < right.(int32)
	case int64:
		return v < right.(int64)
	case uint:
		return v < right.(uint)
	case uint8:
		return v < right.(uint8)
	case uint16:
		return v < right.(uint16)
	case uint32:
		return v < right.(uint32)
	case uint64:
		return v < right.(uint64)
	case string:
		return v < right.(string)
	case float32:
		return v < right.(float32)
	case float64:
		return v < right.(float64)
	default:
		panic(errors.New("not compare type value"))
	}
}

// LTE 小于等于
func LTE(left, right interface{}) bool {
	if IsNil(left) || IsNil(right) {
		return false
	}
	if !TypeEqual(left, right) {
		return false
	}

	switch v := left.(type) {
	case int:
		return v <= right.(int)
	case int8:
		return v <= right.(int8)
	case int16:
		return v <= right.(int16)
	case int32:
		return v <= right.(int32)
	case int64:
		return v <= right.(int64)
	case uint:
		return v <= right.(uint)
	case uint8:
		return v <= right.(uint8)
	case uint16:
		return v <= right.(uint16)
	case uint32:
		return v <= right.(uint32)
	case uint64:
		return v <= right.(uint64)
	case string:
		return v <= right.(string)
	case float32:
		return v <= right.(float32)
	case float64:
		return v <= right.(float64)
	default:
		panic(errors.New("not compare type value"))
	}
}

// GT 大于
func GT(left, right interface{}) bool {
	if IsNil(left) || IsNil(right) {
		return false
	}
	if !TypeEqual(left, right) {
		return false
	}

	switch v := left.(type) {
	case int:
		return v > right.(int)
	case int8:
		return v > right.(int8)
	case int16:
		return v > right.(int16)
	case int32:
		return v > right.(int32)
	case int64:
		return v > right.(int64)
	case uint:
		return v > right.(uint)
	case uint8:
		return v > right.(uint8)
	case uint16:
		return v > right.(uint16)
	case uint32:
		return v > right.(uint32)
	case uint64:
		return v > right.(uint64)
	case string:
		return v > right.(string)
	case float32:
		return v > right.(float32)
	case float64:
		return v > right.(float64)
	default:
		panic(errors.New("not compare type value"))
	}
}

// GTE 大于等于
func GTE(left, right interface{}) bool {
	if IsNil(left) || IsNil(right) {
		return false
	}
	if !TypeEqual(left, right) {
		return false
	}

	switch v := left.(type) {
	case int:
		return v >= right.(int)
	case int8:
		return v >= right.(int8)
	case int16:
		return v >= right.(int16)
	case int32:
		return v >= right.(int32)
	case int64:
		return v >= right.(int64)
	case uint:
		return v >= right.(uint)
	case uint8:
		return v >= right.(uint8)
	case uint16:
		return v >= right.(uint16)
	case uint32:
		return v >= right.(uint32)
	case uint64:
		return v >= right.(uint64)
	case string:
		return v >= right.(string)
	case float32:
		return v >= right.(float32)
	case float64:
		return v >= right.(float64)
	default:
		panic(errors.New("not compare type value"))
	}
}

// EQ 等于
func EQ(left, right interface{}) bool {
	if IsNil(left) || IsNil(right) {
		return false
	}

	if !TypeEqual(left, right) {
		return false
	}

	switch v := left.(type) {
	case int:
		return v == right.(int)
	case int8:
		return v == right.(int8)
	case int16:
		return v == right.(int16)
	case int32:
		return v == right.(int32)
	case int64:
		return v == right.(int64)
	case uint:
		return v == right.(uint)
	case uint8:
		return v == right.(uint8)
	case uint16:
		return v == right.(uint16)
	case uint32:
		return v == right.(uint32)
	case uint64:
		return v == right.(uint64)
	case string:
		return v == right.(string)
	case float32:
		return v == right.(float32)
	case float64:
		return v == right.(float64)
	default:
		panic(errors.New("not compare type value"))
	}
}

// IsNil 检测是否是真nil值
func IsNil(v interface{}) bool {
	if v != nil {
		// 如果目标不是xx类型,则返回
		typ := reflect.TypeOf(v)
		if typ.Kind() != reflect.Ptr {
			return false
		}

		value := reflect.ValueOf(v)
		if !value.IsNil() {
			return false
		}
	}
	return true
}

// In 判断obj是否与target或者targets内的某个元素相等
// 如果 obj 等于 target 或者等于 targets 其中一项,则返回true
// 如果没匹配到其中一项,则返回false
func In(obj interface{}, target interface{}, targets ...interface{}) bool {
	if obj == nil {
		if IsNil(target) {
			return true
		}

		for i := 0; i < len(targets); i++ {
			if IsNil(targets[i]) {
				return true
			}
		}
		return false
	}

	if obj == target {
		return true
	}

	for i := 0; i < len(targets); i++ {
		if obj == targets[i] {
			return true
		}
	}
	return false
}

// EqualAll 判断target跟targets内的所有项与obj是否相等,如果全部target都与obj相等
// 如果obj与target跟targets里面的所有数据都相等,则返回true,如果有其中一项不相等,则返回false
func EqualAll(obj interface{}, target interface{}, targets ...interface{}) bool {
	if obj == nil {
		if !IsNil(target) {
			return false
		}

		for i := 0; i < len(targets); i++ {
			if !IsNil(targets[i]) {
				return false
			}
		}
		return true
	}

	if obj != target {
		return false
	}
	for i := 0; i < len(targets); i++ {
		if obj != targets[i] {
			return false
		}
	}
	return true
}
