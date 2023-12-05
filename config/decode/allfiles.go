package decode

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

/*
遍历指定目录的所有文件
*/

const serverCfgPath = "source/server"
const yyactCfgPath = "source/yyact"

func ReloadPath() {
	walk_files(serverCfgPath, "", reflect.ValueOf(allcfg.Configs))
	walk_yyact(yyactCfgPath, "YY", reflect.ValueOf(allcfg.Yyacts))
}

func walk_yyact(path, ex string, cfg reflect.Value) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		pos := strings.LastIndexByte(file.Name(), '_')
		act := file.Name()[:pos]
		sub := file.Name()[pos+1:]
		pos = strings.Index(act, "_")
		fi := strings.Title(act[pos+1:])
		name := act[:pos] + fi
		nex := ex + name
		pos = strings.Index(sub, ".")
		var idx uint32
		if pos == -1 {
			x, _ := strconv.ParseUint(sub, 10, 32)
			idx = uint32(x)
		} else {
			y := sub[:pos]
			x, _ := strconv.ParseUint(y, 10, 32)
			idx = uint32(x)
		}
		field := cfg.Elem().FieldByName(nex)
		if field.IsNil() {
			field.Set(reflect.MakeMap(field.Type()))
		}
		fkey := reflect.ValueOf(idx)
		row := newPB("pb.Msg" + name)
		field.SetMapIndex(fkey, row)
		if file.IsDir() {
			walk_files(npath, "", row)
		} else {
			ext := sub[pos:]
			dat := readBytes(npath)
			//获取结构对象
			curfield := row.Elem().FieldByName(fi)
			initValue(&curfield)
			if ext == ".ini.txt" {
				decodeIni(dat, curfield.Elem())
			} else if ext == ".json.txt" {
				inf := curfield.Interface()
				decodeJson(dat, &inf)
			} else if ext == ".txt" {
				decodeTxt(dat, &curfield)
			}
		}
	}
}

func walk_files(path, ex string, cfg reflect.Value) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			nex := ex + strings.Title(file.Name())
			walk_files(npath, nex, cfg)
		} else {
			ext := filepath.Ext(file.Name())
			fi := strings.TrimSuffix(file.Name(), ext)

			if ext == ".txt" && strings.HasSuffix(fi, ".ini") {
				ext = ".ini.txt"
				fi = strings.TrimSuffix(fi, ".ini")
			}
			if ext == ".txt" && strings.HasSuffix(fi, ".json") {
				ext = ".json.txt"
				fi = strings.TrimSuffix(fi, ".json")
			}

			nex := ex + strings.Title(fi)
			dat := readBytes(npath)
			//获取结构对象
			field := cfg.Elem().FieldByName(nex)
			initValue(&field)
			if ext == ".ini.txt" {
				decodeIni(dat, field.Elem())
			} else if ext == ".json.txt" {
				inf := field.Interface()
				decodeJson(dat, &inf)
			} else if ext == ".txt" {
				decodeTxt(dat, &field)
			}
		}
	}
}

func readBytes(npath string) string {
	file, err := os.Open(npath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	str := string(byteValue)
	if strings.HasPrefix(str, "\uFEFF") {
		str = strings.TrimPrefix(str, "\uFEFF") //去除utf8 with bom中的bom
	}
	return str
}
