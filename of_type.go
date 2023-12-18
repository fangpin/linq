package linq

import "reflect"

func OfType(source []any, expectType reflect.Type) []any {
	target := []any{}
	for _, s := range source {
		reflect.TypeOf(expectType)
		if t, ok := s.(int); ok {
			target = append(target, t)
		}
	}
	return target
}
