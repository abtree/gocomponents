package src

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
遍历指定目录的所有文件
*/

func ReloadPath() {
	walk_files(serverCfgPath, "", false)
	walk_yyact(yyactCfgPath, "YY")
}

func walk_yyact(path, ex string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		pos := strings.LastIndexByte(file.Name(), '_')
		act := file.Name()[:pos]
		sub := file.Name()[pos+1:]
		pos = strings.Index(act, "_")
		fi := strings.Title(act[pos+1:])
		name := act[:pos] + fi
		nex := ex + name
		if _, ok := had_yyact[nex]; ok {
			continue //已经解析过的活动
		}
		had_yyact[nex] = true
		npath := filepath.Join(path, file.Name())
		writeYYact(nex, "Msg"+name)
		filemsg = bytes.NewBuffer([]byte{})
		filemsg.WriteString("message Msg" + name + "{")
		fileIndex = 1
		if file.IsDir() {
			walk_files(npath, "", true)
		} else {
			pos = strings.Index(sub, ".")
			ext := sub[pos:]
			dat := readBytes(npath)
			if ext == ".ini.txt" {
				parserIni(fi, dat, true)
			} else if ext == ".json.txt" {
				parserJson(fi, fi, dat, true)
			} else if ext == ".txt" {
				parserTxt(fi, dat, true)
			}
		}
		filemsg.WriteString(end_chat)
		filemsg.WriteString("}")
		wfile.WriteString(filemsg.String())
		endChat()
	}
}

func walk_files(path, ex string, isyyact bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln("Read Configs error")
	}
	for _, file := range files {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			nex := ex + strings.Title(file.Name())
			walk_files(npath, nex, isyyact)
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
			if ext == ".ini.txt" {
				parserIni(nex, dat, isyyact)
			} else if ext == ".json.txt" {
				parserJson(nex, fi, dat, isyyact)
			} else if ext == ".txt" {
				parserTxt(nex, dat, isyyact)
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
