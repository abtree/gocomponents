package src

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	serverCfgPath = "source/server"
	yyactCfgPath  = "source/yyact"
	pbPath        = "target/"
	end_chat      = "\r\n"
	end_begin     = "\r\n\t"
	pb_array      = "[]"
)

var (
	had_pbs    map[string]bool
	had_yyact  map[string]bool
	wfile      *os.File
	bigmsg     *bytes.Buffer
	bigindex   int
	yyactmsg   *bytes.Buffer
	yyactindex int
	filemsg    *bytes.Buffer
	fileIndex  int
)

func spile(a rune) bool {
	if a == '_' {
		return true
	}
	return false
}

func endChat() {
	wfile.WriteString(end_chat)
}

func writeHead() {
	wfile.WriteString(`syntax = "proto3";`)
	endChat()
	endChat()
	wfile.WriteString(`option optimize_for = LITE_RUNTIME;`)
	endChat()
	endChat()
	// wfile.WriteString(`import "msg_public.proto";`)
	// endChat()
	// wfile.WriteString(`import "msg.proto"; `)
	// endChat()
	wfile.WriteString(`import "msg_config.proto"; `)
	endChat()
	//wfile.WriteString(`import "msg_fight.proto"; `)
	endChat()
	endChat()
	wfile.WriteString(`//---------------说明---------------------------`)
	endChat()
	wfile.WriteString(`//自动生成的pb 配置文件`)
	endChat()
	wfile.WriteString(`//----------------------------------------------`)
	endChat()
	endChat()
	wfile.WriteString(`package pb;`)
	endChat()
	endChat()
	bigmsg.WriteString("message MsgConfigs {")
	bigmsg.WriteString(end_begin)
	bigmsg.WriteString("map<string,bytes> unhandle = 1;")
	bigmsg.WriteString(end_begin)
	//bigmsg.WriteString("repeated MsgStrKeyValueArray dirtyWords = 2;")
	yyactmsg.WriteString("message MsgYYactConfigs {")
	yyactmsg.WriteString(end_begin)
	yyactmsg.WriteString("map<string,bytes> unhandle = 1;")
}

func writeBig(path, pbname, key string) {
	if strings.HasPrefix(path, "YY") {
		writeFile(path, pbname, key)
		return
	}
	bigmsg.WriteString(end_begin)
	if key != "" {
		bigmsg.WriteString("map<" + filterKey(key) + ", " + pbname + ">")
	} else {
		bigmsg.WriteString(pbname)
	}
	bigmsg.WriteString(" ")
	bigmsg.WriteString(path)
	bigmsg.WriteString(" = ")
	bigmsg.WriteString(strconv.Itoa(bigindex))
	bigindex++
	bigmsg.WriteString(";")
}

func writeFile(path, pbname, key string) {
	filemsg.WriteString(end_begin)
	if key != "" {
		filemsg.WriteString("map<" + filterKey(key) + ", " + pbname + ">")
	} else {
		filemsg.WriteString(pbname)
	}
	filemsg.WriteString(" ")
	filemsg.WriteString(path)
	filemsg.WriteString(" = ")
	filemsg.WriteString(strconv.Itoa(fileIndex))
	fileIndex++
	filemsg.WriteString(";")
}

func writeYYact(path, pbname string) {
	yyactmsg.WriteString(end_begin)
	yyactmsg.WriteString("map<uint32, " + pbname + ">")
	yyactmsg.WriteString(" ")
	yyactmsg.WriteString(path)
	yyactmsg.WriteString(" = ")
	yyactmsg.WriteString(strconv.Itoa(yyactindex))
	yyactindex++
	yyactmsg.WriteString(";")
}

func filterKey(key string) string {
	//enum枚举类型不能做key
	if key[0] >= 'A' && key[0] <= 'Z' {
		return "uint32"
	}
	return key
}

func writeFin() {
	bigmsg.WriteString(end_chat)
	bigmsg.WriteString("}")
	wfile.WriteString(bigmsg.String())
	endChat()
	yyactmsg.WriteString(end_chat)
	yyactmsg.WriteString("}")
	wfile.WriteString(yyactmsg.String())
	endChat()
	wfile.WriteString(`
	message MsgAllConfigs{
		MsgConfigs Configs = 1;
		MsgYYactConfigs Yyacts = 2;
	}
	`)
}

func Run() {
	had_pbs = make(map[string]bool)
	had_yyact = make(map[string]bool)
	var err error
	wfile, err = os.OpenFile(pbPath+"msg_cfg_auto.proto", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Open write File %v error with : %v", pbPath, err)
	}
	defer wfile.Close()
	bigmsg = bytes.NewBuffer([]byte{})
	bigindex = 2
	yyactmsg = bytes.NewBuffer([]byte{})
	yyactindex = 2

	writeHead()
	ReloadPath()
	writeFin()
}
