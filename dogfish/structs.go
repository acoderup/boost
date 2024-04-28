package dogfish

import (
	"fmt"
	"reflect"
)

type mapConverter struct {
	Tag string
}

func newMapConverter(tag string) *mapConverter {
	return &mapConverter{
		Tag: tag,
	}
}

func (mc *mapConverter) joinPrefix(k, prefix string) string {
	if prefix == "" {
		return k
	}
	return fmt.Sprintf("%s.%s", prefix, k)
}

func (mc *mapConverter) fieldKey(typ reflect.Type, i int) string {
	if mc.Tag == "" {
		return typ.Field(i).Name
	} else {
		return typ.Field(i).Tag.Get(mc.Tag)
	}
}

func (mc *mapConverter) recursiveToMap(mapValue map[string]interface{}, structValue interface{}, prefix string) error {
	value := reflect.ValueOf(structValue).Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		k := mc.fieldKey(typ, i)

		// filter unexported field
		if typ.Field(i).PkgPath != "" {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			fieldValue = fieldValue.Elem()
		}
		if fieldValue.Kind() == reflect.Struct {
			err := mc.recursiveToMap(mapValue, fieldValue.Addr().Interface(), mc.joinPrefix(k, prefix))
			if err != nil {
				return err
			}
			continue
		}
		mapValue[mc.joinPrefix(k, prefix)] = fieldValue.Interface()
	}
	return nil
}

func (mc *mapConverter) toMap(structValue interface{}) (map[string]interface{}, error) {
	value := reflect.ValueOf(structValue)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, fmt.Errorf("struct value cannot be nil")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct value kind is not Struct or Ptr of Struct")
	}
	mapValue := make(map[string]interface{})
	err := mc.recursiveToMap(mapValue, value.Addr().Interface(), "")
	if err != nil {
		return nil, err
	}
	return mapValue, nil
}

func (mc *mapConverter) recursiveFromMap(mapValue map[string]interface{}, structValue interface{}, prefix string) error {
	value := reflect.ValueOf(structValue).Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		k := mc.fieldKey(typ, i)

		// filter unexported field
		if typ.Field(i).PkgPath != "" {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			fieldValue = fieldValue.Elem()
		}
		if fieldValue.Kind() == reflect.Struct {
			err := mc.recursiveFromMap(mapValue, fieldValue.Addr().Interface(), mc.joinPrefix(k, prefix))
			if err != nil {
				return err
			}
			continue
		}
		if v, ok := mapValue[mc.joinPrefix(k, prefix)]; ok {
			fieldValue.Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func (mc *mapConverter) fromMap(mapValue map[string]interface{}, structValue interface{}) error {
	value := reflect.ValueOf(structValue)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return fmt.Errorf("struct value cannot be nil")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("struct value kind is not Struct or Ptr of Struct")
	}
	return mc.recursiveFromMap(mapValue, value.Addr().Interface(), "")
}

type stringMapConverter struct {
	Tag string
}

func newStringMapConverter(tag string) *stringMapConverter {
	return &stringMapConverter{
		Tag: tag,
	}
}

func (smc *stringMapConverter) joinPrefix(k, prefix string) string {
	if prefix == "" {
		return k
	}
	return fmt.Sprintf("%s.%s", prefix, k)
}

func (smc *stringMapConverter) fieldKey(typ reflect.Type, i int) string {
	if smc.Tag == "" {
		return typ.Field(i).Name
	} else {
		return typ.Field(i).Tag.Get(smc.Tag)
	}
}

func (smc *stringMapConverter) recursiveToMap(mapValue map[string]string, structValue interface{}, prefix string) error {
	value := reflect.ValueOf(structValue).Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		k := smc.fieldKey(typ, i)

		// filter unexported field
		if typ.Field(i).PkgPath != "" {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			fieldValue = fieldValue.Elem()
		}
		if fieldValue.Kind() == reflect.Struct {
			err := smc.recursiveToMap(mapValue, fieldValue.Addr().Interface(), smc.joinPrefix(k, prefix))
			if err != nil {
				return err
			}
			continue
		}
		var err error
		mapValue[smc.joinPrefix(k, prefix)], err = ToString(fieldValue.Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func (smc *stringMapConverter) toMap(structValue interface{}) (map[string]string, error) {
	value := reflect.ValueOf(structValue)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, fmt.Errorf("struct value cannot be nil")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct value kind is not Struct or Ptr of Struct")
	}
	mapValue := make(map[string]string)
	err := smc.recursiveToMap(mapValue, value.Addr().Interface(), "")
	if err != nil {
		return nil, err
	}
	return mapValue, nil
}

func (smc *stringMapConverter) recursiveFromMap(mapValue map[string]string, structValue interface{}, prefix string) error {
	value := reflect.ValueOf(structValue).Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		k := smc.fieldKey(typ, i)

		// filter unexported field
		if typ.Field(i).PkgPath != "" {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			fieldValue = fieldValue.Elem()
		}
		if fieldValue.Kind() == reflect.Struct {
			err := smc.recursiveFromMap(mapValue, fieldValue.Addr().Interface(), smc.joinPrefix(k, prefix))
			if err != nil {
				return err
			}
			continue
		}
		if s, ok := mapValue[smc.joinPrefix(k, prefix)]; ok {
			err := FromString(s, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (smc *stringMapConverter) fromMap(mapValue map[string]string, structValue interface{}) error {
	value := reflect.ValueOf(structValue)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return fmt.Errorf("struct value cannot be nil")
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("struct value kind is not Struct or Ptr of Struct")
	}
	return smc.recursiveFromMap(mapValue, value.Addr().Interface(), "")
}

func ToMap(structValue interface{}) map[string]interface{} {
	m, err := ToMapE(structValue)
	if err != nil {
		panic(err)
	}
	return m
}

func ToMapE(structValue interface{}) (map[string]interface{}, error) {
	return newMapConverter("").toMap(structValue)
}

func ToJSONMap(structValue interface{}) map[string]interface{} {
	m, err := ToJSONMapE(structValue)
	if err != nil {
		panic(err)
	}
	return m
}

func ToJSONMapE(structValue interface{}) (map[string]interface{}, error) {
	return newMapConverter("json").toMap(structValue)
}

func FromMap(mapValue map[string]interface{}, structValue interface{}) {
	err := FromMapE(mapValue, structValue)
	if err != nil {
		panic(err)
	}
}

func FromMapE(mapValue map[string]interface{}, structValue interface{}) error {
	return newMapConverter("").fromMap(mapValue, structValue)
}

func FromJSONMap(mapValue map[string]interface{}, structValue interface{}) {
	err := FromJSONMapE(mapValue, structValue)
	if err != nil {
		panic(err)
	}
}

func FromJSONMapE(mapValue map[string]interface{}, structValue interface{}) error {
	return newMapConverter("json").fromMap(mapValue, structValue)
}

func ToStringMap(structValue interface{}) map[string]string {
	m, err := ToStringMapE(structValue)
	if err != nil {
		panic(err)
	}
	return m
}

func ToStringMapE(structValue interface{}) (map[string]string, error) {
	return newStringMapConverter("").toMap(structValue)
}

func ToStringJSONMap(structValue interface{}) map[string]string {
	m, err := ToStringJSONMapE(structValue)
	if err != nil {
		panic(err)
	}
	return m
}

func ToStringJSONMapE(structValue interface{}) (map[string]string, error) {
	return newStringMapConverter("json").toMap(structValue)
}

func FromStringMap(mapValue map[string]string, structValue interface{}) {
	err := FromStringMapE(mapValue, structValue)
	if err != nil {
		panic(err)
	}
}

func FromStringMapE(mapValue map[string]string, structValue interface{}) error {
	return newStringMapConverter("").fromMap(mapValue, structValue)
}

func FromStringJSONMap(mapValue map[string]string, structValue interface{}) {
	err := FromStringJSONMapE(mapValue, structValue)
	if err != nil {
		panic(err)
	}
}

func FromStringJSONMapE(mapValue map[string]string, structValue interface{}) error {
	return newStringMapConverter("json").fromMap(mapValue, structValue)
}
