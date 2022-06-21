package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type PriceItem struct {
	LV    int32
	Award string
}

type Data struct {
	Id    int32
	Items []*PriceItem
	Desc  []string
}

func splitRule(c rune) bool {
	if c == '\t' || c == ' ' {
		return true
	} else {
		return false
	}
}

func main() {
	lines := []string{
		"1 1 aaa 2 bbb 3 ccc ddd eee",
		"2 1 aaa 2 bbb 3 ccc ddd eee",
		"3 1 aaa 2 bbb 3 ccc ddd eee",
	}

	data := &Data{
		Items: make([]*PriceItem, 3),
		Desc:  make([]string, 2),
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		strs := strings.FieldsFunc(line, splitRule)
		pos := 0
		parseStruct(reflect.ValueOf(data).Elem(), strs, &pos)
	}

	fmt.Println(data)
}

func parseStruct(target reflect.Value, strs []string, pos *int) {
	size := target.NumField()
	for i := 0; i < size; i++ {
		field := target.Field(i)
		if field.Kind() == reflect.Ptr {
			parsePtr(field, strs, pos)
		} else if field.Kind() == reflect.Struct {
			parseStruct(field, strs, pos)
		} else if field.Kind() == reflect.Slice {
			parseSlice(field, strs, pos)
		} else {
			fullField(&field, strs[*pos])
			*pos++
		}
	}
}

func parsePtr(ptr reflect.Value, strs []string, pos *int) {
	if ptr.IsNil() {
		ptr.Set(reflect.New(ptr.Type().Elem()))
	}
	field := ptr.Elem()
	if field.Kind() == reflect.Ptr {
		parsePtr(field, strs, pos)
	} else if field.Kind() == reflect.Struct {
		parseStruct(field, strs, pos)
	} else if field.Kind() == reflect.Slice {

	} else {
		fullField(&field, strs[*pos])
		*pos++
	}
}

func parseSlice(array reflect.Value, strs []string, pos *int) {
	if array.IsNil() {
		return
	}
	for i := 0; i < array.Len(); i++ {
		field := array.Index(i)
		if field.Kind() == reflect.Ptr {
			parsePtr(field, strs, pos)
		} else if field.Kind() == reflect.Struct {
			parseStruct(field, strs, pos)
		} else if field.Kind() == reflect.Slice {
			parseSlice(field, strs, pos)
		} else {
			fullField(&field, strs[*pos])
			*pos++
		}
	}
}

func fullField(field *reflect.Value, value string) bool {
	switch field.Type().Kind() {
	case reflect.String,
		reflect.Int32,
		reflect.Int,
		reflect.Int64,
		reflect.Uint32,
		reflect.Uint,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool:
		{
			v, ok := parseValue(field.Type().Kind(), value)
			if !ok {
				return false
			}
			field.Set(reflect.ValueOf(v).Convert(field.Type()))
		}
	case reflect.Slice:
		{
			typ := field.Type().Elem().Kind()
			v, ok := parseValue(typ, value)
			if !ok {
				return false
			}
			field.Set(reflect.Append(*field, reflect.ValueOf(v).Convert(field.Type().Elem())))
		}
	}
	return true
}

func parseValue(typ reflect.Kind, value string) (interface{}, bool) {
	switch typ {
	case reflect.String:
		{
			return value, true
		}
	case reflect.Int32:
		{
			f, err := strconv.Atoi(value)
			if err != nil {
				return nil, false
			}
			return int32(f), true
		}
	case reflect.Int,
		reflect.Int64:
		{
			f, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, false
			}
			return f, true
		}
	case reflect.Uint32:
		{
			f, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return nil, false
			}
			return f, true
		}
	case reflect.Uint,
		reflect.Uint64:
		{
			f, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return nil, false
			}
			return f, true
		}
	case reflect.Float32:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				return nil, false
			}
			return float32(f), true
		}
	case reflect.Float64:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				return nil, false
			}
			return f, true
		}
	case reflect.Bool:
		{
			f, err := strconv.ParseBool(value)
			if err != nil {
				return nil, false
			}
			return f, true
		}
	}
	return nil, false
}
