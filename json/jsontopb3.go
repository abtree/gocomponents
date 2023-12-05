package main

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var strbuilder strings.Builder

//dat := `{"a":1,"b":[1,2,3], "c":{"d":4.2,"e":5.4}, "f":[{"g":"str1","h":"str2"},{"g":"str3","h":"str4"}]}`
func parser(k string, dat string) {
	var config interface{}
	strbuilder = strings.Builder{}
	err := json.Unmarshal([]byte(dat), &config)
	if err != nil {
		log.Panicf(err.Error())
	}
	m := config.(map[string]interface{})
	parseMap("Cfg", m)
	log.Println(strbuilder.String())
}

func parseMap(nam string, m map[string]interface{}) {
	var seq = 1
	strbuilder.WriteString("message Cfg" + nam + "{\n")
	for k, v := range m {
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Slice {
			x := v.([]interface{})[0]
			if reflect.TypeOf(x).Kind() == reflect.Map {
				parseMap(k, x.(map[string]interface{}))
				strbuilder.WriteString("\t repeated Cfg" + k + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
			} else if kind == reflect.Float64 {
				t := switchTyp(x)
				strbuilder.WriteString("\t repeated " + t + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
			} else {
				strbuilder.WriteString("\t repeated " + reflect.TypeOf(x).Name() + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
			}
			// for _, x := range v.([]interface{}) {
			// 	log.Println(reflect.TypeOf(x).Name())
			// }
		} else if kind == reflect.Map {
			parseMap(k, v.(map[string]interface{}))
			strbuilder.WriteString("\t Cfg" + k + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
		} else if kind == reflect.Float64 {
			t := switchTyp(v)
			strbuilder.WriteString("\t " + t + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
		} else {
			strbuilder.WriteString("\t " + reflect.TypeOf(v).Name() + " " + strings.Title(k) + " = " + strconv.Itoa(seq) + ";\n")
		}
		seq++
	}
	strbuilder.WriteString("} \n")
}

func switchTyp(v interface{}) string {
	s := strconv.FormatFloat(v.(float64), 'f', -1, 64)
	if _, err := strconv.Atoi(s); err == nil {
		return "sint32"
	} else {
		return "double"
	}
}
