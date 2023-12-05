package src

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var strbuilder strings.Builder

func parserJson(path, nam, dat string, isyyact bool) {
	pbname := buildJson(nam, dat)
	if _, ok := had_pbs[pbname]; ok {
		log.Fatalln("json pb %s has repeated \n", pbname)
	}
	if isyyact {
		writeFile(path, pbname, "")
	} else {
		writeBig(path, pbname, "")
	}
}

func buildJson(nam, dat string) (pbname string) {
	var config interface{}
	strbuilder = strings.Builder{}
	err := json.Unmarshal([]byte(dat), &config)
	if err != nil {
		log.Panicf(err.Error())
	}
	if reflect.TypeOf(config).Kind() == reflect.Slice {
		m := config.([]interface{})[0]
		parseMap(nam, m.(map[string]interface{}))
		pbname = "repeated J" + nam
	} else {
		m := config.(map[string]interface{})
		parseMap(nam, m)
		pbname = "J" + nam
	}

	wfile.WriteString(strbuilder.String())
	return
}

func parseMap(nam string, m map[string]interface{}) {
	var seq = 1
	strbuilder.WriteString("message J" + nam + "{\n")
	for k, v := range m {
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Slice {
			x := v.([]interface{})[0]
			if reflect.TypeOf(x).Kind() == reflect.Map {
				parseMap(k, x.(map[string]interface{}))
				strbuilder.WriteString("\t repeated J" + k + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			} else if reflect.TypeOf(x).Kind() == reflect.Float64 {
				t := switchTyp(x)
				strbuilder.WriteString("\t repeated " + t + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			} else {
				strbuilder.WriteString("\t repeated " + reflect.TypeOf(x).Name() + " " + k + " = " + strconv.Itoa(seq) + ";\n")
			}
			// for _, x := range v.([]interface{}) {
			// 	log.Println(reflect.TypeOf(x).Name())
			// }
		} else if kind == reflect.Map {
			parseMap(k, v.(map[string]interface{}))
			strbuilder.WriteString("\t J" + k + " " + k + " = " + strconv.Itoa(seq) + ";\n")
		} else if kind == reflect.Float64 {
			t := switchTyp(v)
			strbuilder.WriteString("\t " + t + " " + k + " = " + strconv.Itoa(seq) + ";\n")
		} else {
			strbuilder.WriteString("\t " + reflect.TypeOf(v).Name() + " " + k + " = " + strconv.Itoa(seq) + ";\n")
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
