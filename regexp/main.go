package main

import (
	"log"
	"regexp"
	"strings"
)

var (
	rFloat *regexp.Regexp
	rInt   *regexp.Regexp
	rInit  *regexp.Regexp
	rSub   *regexp.Regexp
	rArray *regexp.Regexp
)

func regFloat(str string) bool {
	if rFloat.MatchString(str) {
		return true
	}
	return false
}

func regInt(str string) bool {
	if rInt.MatchString(str) {
		return true
	}
	return false
}

func regInit(str string) bool {
	if rInit.MatchString(str) {
		return true
	}
	return false
}

func spile(a rune) bool {
	if a == '_' {
		return true
	}
	return false
}

func chooseArray(str string) int {
	typ := 3
	strs := strings.FieldsFunc(str, spile)
	for _, s := range strs {
		if regInt(s) {
			typ = 1
			continue
		} else if regFloat(s) {
			typ = 2
			continue
		} else {
			typ = 3
			break
		}
	}
	if typ == 1 {
		log.Printf("%v is array int", str)
	} else if typ == 2 {
		log.Printf("%v is array float", str)
	} else {
		log.Printf("%v is array string", str)
	}
	return typ
}

func main() {
	//str := "123_123"
	//str1 := "123.123"
	//str2 := "321"
	str3 := "q321d"
	str4 := "uint32ValuePay"
	//str5 := "Title"
	str6 := "[]floatValuePay"
	str7 := "123_123_123"
	str8 := "1_0.5_0.2"
	str9 := "3_a_5.0"
	rFloat, _ = regexp.Compile(`^[-+]?[0-9]+\.[0-9]+$`)
	rInt, _ = regexp.Compile(`^-?\d+$`)
	rInit, _ = regexp.Compile(`^(\[\])?[a-z]`)
	rSub, _ = regexp.Compile(`^(\[\])?[a-z][a-z0-9_]*[A-Z]`)

	rArray, _ = regexp.Compile(`^(.+_)+`)

	strFind := rSub.FindString(str4)
	if strFind != "" {
		log.Printf("%v is with sub string %v", str4, strFind)
	}

	strFind = rSub.FindString(str6)
	if strFind != "" {
		log.Printf("%v is with sub string %v", str6, strFind)
	}

	if rArray.MatchString(str7) {
		chooseArray(str7)
	}
	if rArray.MatchString(str8) {
		chooseArray(str8)
	}
	if rArray.MatchString(str9) {
		chooseArray(str9)
	}
	if rArray.MatchString(str3) {
		chooseArray(str3)
	}

	//if regFloat(str) {
	//	log.Printf("%v is a float", str)
	//}
	//if regFloat(str1) {
	//	log.Printf("%v is a float", str1)
	//}
	//if regFloat(str2) {
	//	log.Printf("%v is a float", str2)
	//}
	//if regInt(str) {
	//	log.Printf("%v is a int", str)
	//}
	//if regInt(str1) {
	//	log.Printf("%v is a int", str1)
	//}
	//if regInt(str2) {
	//	log.Printf("%v is a int", str2)
	//}
	//if regInt(str3) {
	//	log.Printf("%v is a int", str3)
	//}
	//if regInit(str2) {
	//	log.Printf("%v is init with alpha", str2)
	//}
	//if regInit(str3) {
	//	log.Printf("%v is init with alpha", str3)
	//}
	//if regInit(str4) {
	//	log.Printf("%v is init with alpha", str4)
	//}
	//if regInit(str6) {
	//	log.Printf("%v is init with alpha", str6)
	//}

	//strTitle := strings.Title(str4)
	//log.Printf("string to title: %v", strTitle)
	//log.Printf("string to title: %v", strings.Title(str3))
	//log.Printf("string to title: %v", strings.Title(str5))
	//log.Printf("string to title: %v", strings.Title("cpriceitem"))
}
