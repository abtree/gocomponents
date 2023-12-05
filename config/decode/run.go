package decode

import (
	"config/pb"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

var (
	allcfg *pb.MsgAllConfigs
	// allcfg      *pb.MsgConfigs
	// allcfgValue reflect.Value
	// allcfgType  reflect.Type
)

func Run() {
	allcfg = &pb.MsgAllConfigs{
		Configs: &pb.MsgConfigs{},
		Yyacts:  &pb.MsgYYactConfigs{},
	}
	// allcfgValue = reflect.ValueOf(allcfg).Elem()
	// allcfgType = reflect.TypeOf(*allcfg)
	ReloadPath()
	log.Printf("%+v", allcfg)
}

func fullField(field *reflect.Value, value string, split string) bool {
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
			vals := []string{}
			if split != "" {
				vals = strings.Split(value, split)
			} else {
				vals = append(vals, value)
			}
			for _, vs := range vals {
				v, ok := parseValue(typ, vs)
				if !ok {
					return false
				}
				field.Set(reflect.Append(*field, reflect.ValueOf(v).Convert(field.Type().Elem())))
			}
		}
	case reflect.Ptr:
		{
			vals := []string{}
			if split != "" {
				vals = strings.Split(value, split)
			} else {
				vals = append(vals, value)
			}
			initValue(field)
			// log.Println(field.Elem().NumField())
			for i, vs := range vals {
				fi := field.Elem().Field(i)
				fullField(&fi, vs, "")
			}
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
				panic(err)
				return nil, false
			}
			return int32(f), true
		}
	case reflect.Int,
		reflect.Int64:
		{
			f, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				panic(err)
				return nil, false
			}
			return f, true
		}
	case reflect.Uint32:
		{
			f, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				panic(err)
				return nil, false
			}
			return f, true
		}
	case reflect.Uint,
		reflect.Uint64:
		{
			f, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				panic(err)
				return nil, false
			}
			return f, true
		}
	case reflect.Float32:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				panic(err)
				return nil, false
			}
			return float32(f), true
		}
	case reflect.Float64:
		{
			f, err := strconv.ParseFloat(value, 10)
			if err != nil {
				panic(err)
				return nil, false
			}
			return f, true
		}
	case reflect.Bool:
		{
			f, err := strconv.ParseBool(value)
			if err != nil {
				panic(err)
				return nil, false
			}
			return f, true
		}
	}
	return nil, false
}

func newPB(name string) reflect.Value {
	return reflect.New(proto.MessageType(name).Elem())
}

// func getField(name string) *reflect.Value {
// 	id := getFieldIdByTag(name)
// 	if id < 0 {
// 		log.Panic("cfg file not in pb")
// 	}
// 	target := allcfgValue.Field(id)
// 	return &target
// }

// func getFieldIdByTag(name string) int {
// 	size := allcfgType.NumField()
// 	for i := 0; i < size; i++ {
// 		field := allcfgType.Field(i)
// 		str := string(field.Tag)
// 		if strings.Contains(str, name) {
// 			return i
// 		}
// 	}
// 	return -1
// }
