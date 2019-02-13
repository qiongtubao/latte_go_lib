package sqlite3

import "reflect"

type Object struct {
	modelClass reflect.Type
	data       reflect.Value
}

func (object Object) Get(key string) interface{} {
	value := object.data.FieldByName(key)
	return value.Interface()
}

func (object Object) Set(key string, value interface{}) error {
	v := object.data.FieldByName(key)
	v.Set(reflect.ValueOf(value))
	return nil
}
func (object Object) Interface() interface{} {
	return object.data.Interface()
}
func CreateObject(model reflect.Type) *Object {
	data := reflect.New(model).Elem()
	return &Object{
		modelClass: model,
		data:       data,
	}
}
