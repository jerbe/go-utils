package goutils

import (
	"errors"
	"reflect"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:09
  @describe :
*/

// SliceUnique 将切片进行唯一归类
func SliceUnique(data interface{}, dst interface{}) error {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	if dataType.Kind() != reflect.Slice {
		return errors.New("`data` is no slice type")
	}

	dstType := reflect.TypeOf(dst)
	if dstType.Kind() != reflect.Ptr {
		return errors.New("`dst` is no ptr")
	}
	dstType = dstType.Elem()
	if dstType.Kind() != reflect.Slice {
		return errors.New("`dst` is no slice ptr")
	}

	if dataType.Kind() != dstType.Kind() {
		return errors.New("`data` is not the same type as `dst`")
	}

	var mapType = reflect.MapOf(dataType.Elem(), reflect.TypeOf(reflect.Interface))

	var mapValue = reflect.MakeMap(mapType)

	var dataValue = reflect.Indirect(reflect.ValueOf(data))

	//v := struct{}{}
	for i := 0; i < dataValue.Len(); i++ {
		mapValue.SetMapIndex(dataValue.Index(i), reflect.ValueOf(reflect.Struct))
	}

	dstValue := reflect.Indirect(reflect.ValueOf(dst))
	newDstValue := reflect.MakeSlice(dstType, mapValue.Len(), mapValue.Len())

	//dstValue = dstValue.
	var i = 0
	itr := mapValue.MapRange()
	for itr.Next() {
		newDstValue.Index(i).Set(itr.Key())
		i++
	}
	dstValue.Set(newDstValue)
	return nil
}
