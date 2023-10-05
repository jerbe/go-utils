package goutils

import "reflect"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 20:18
  @describe :
*/

// TypeEqual 判断两个对象的类型是否相等
func TypeEqual(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
