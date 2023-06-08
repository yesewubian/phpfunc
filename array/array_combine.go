package array

import (
	"fmt"
	"reflect"
)

/**
 * 将两个slice合并成一个map
 *
 */
func ArrayCombine(desk, sliceK, sliceV interface{}) (err error) {

	deskType := reflect.TypeOf(desk)
	deskValue := reflect.ValueOf(desk)
	direct := reflect.Indirect(deskValue)
	mapReflect := reflect.MakeMap(deskType.Elem())
	sliceKValue := reflect.ValueOf(sliceK)
	sliceKType := reflect.TypeOf(sliceK)
	sliceVValue := reflect.ValueOf(sliceV)
	sliceVType := reflect.TypeOf(sliceV)

	if deskType.Kind() != reflect.Ptr {
		return fmt.Errorf("desk'type expect Prt, got %s", deskType.Kind())
	}
	if deskType.Elem().Kind() != reflect.Map {
		return fmt.Errorf("desk'value expect map, got %s", deskType.Elem().Kind())
	}
	if sliceKValue.Kind() != reflect.Slice {
		return fmt.Errorf("sliceK'type expect Slice, got %s", sliceKValue.Kind())
	}
	if sliceVValue.Kind() != reflect.Slice {
		return fmt.Errorf("sliceV'type expect Slice, got %s", sliceVValue.Kind())
	}

	if sliceKValue.Len() != sliceVValue.Len() {
		return fmt.Errorf("sliceK's len expect equal to sliceV's")
	}

	if deskType.Elem().Key().Kind() != sliceKType.Elem().Kind() {
		return fmt.Errorf("desk's key type expect %s, got %s", deskType.Elem().Key().Kind(), sliceKType.Elem().Kind())
	}

	if deskType.Elem().Elem().Kind() != sliceVType.Elem().Kind() {
		return fmt.Errorf("desk's element type expect %s, got %s", deskType.Elem().Elem().Kind(), sliceVType.Elem().Kind())
	}

	// fmt.Println(sliceKValue.Len())
	for i := 0; i < sliceKValue.Len(); i++ {
		// fmt.Println(sliceKValue.Index(i).Int())
		mapReflect.SetMapIndex(sliceKValue.Index(i), sliceVValue.Index(i))
		direct.Set(mapReflect)
	}

	return
}
