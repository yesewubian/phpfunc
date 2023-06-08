package function

import (
	"reflect"
)

func ArrayFilter(desk interface{}) {
	deskVal := reflect.ValueOf(desk).Elem()
	direct := reflect.Indirect(deskVal)
	iter := deskVal.MapRange()
	for iter.Next() {
		switch iter.Value().Kind() {
		//todo other kind need to be suported
		case reflect.String:
			if iter.Value().String() == "" {
				direct.SetMapIndex(iter.Key(), reflect.ValueOf(""))
			}
		default:
			if iter.Value().Int() == 0 {
				direct.SetMapIndex(iter.Key(), reflect.ValueOf(0))
			}
		}
	}
}
