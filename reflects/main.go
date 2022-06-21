package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type PriceItem struct {
	Lv    int32
	Award string
}

type Data struct {
	Id     int32
	Awards []*PriceItem
}

type cfgcolumn struct {
	name  string
	vtype string
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
		"1001 1 aaa",
		"1001 10 bbb",
		"1001 30 ccc",
	}
	column := make(map[int]cfgcolumn)
	column[0] = cfgcolumn{
		vtype: "Id",
	}
	column[1] = cfgcolumn{
		vtype: "Lv",
	}
	column[2] = cfgcolumn{
		vtype: "Award",
	}
	data := &Data{}

	for j, l := range lines {
		l = strings.TrimSpace(l)
		strs := strings.FieldsFunc(l, splitRule)

		for i, value := range strs {
			col, ok := column[i]
			if !ok {
				continue
			}
			tv := reflect.ValueOf(data)
			if i == 0 {
				field := tv.Elem().FieldByName(col.vtype)
				fullField(&field, value)
			} else {
				t := tv.Elem().FieldByName("Awards")
				if t.IsNil() {
					slic := reflect.MakeSlice(t.Type(), 0, 0)
					t.Set(slic)
				}
				for t.Len() <= j {
					fmt.Printf("%v %v \n", j, t.Type().Elem().Elem())
					vvv := reflect.New(t.Type().Elem().Elem())
					t.Set(reflect.Append(t, vvv))
				}
				t = t.Index(j)
				vv := t.Elem().FieldByName(col.vtype)
				fullField(&vv, value)
			}
		}
	}
	fmt.Printf("%v \n", data)
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
