package main

import "os"
import (
	"io/ioutil"
	"strings"
	"fmt"
)

var protoFilePath string
var serverSessionFilePath string

func main() {
	args := os.Args
	if (len(args) > 2) {
		protoFilePath = args[1]
		serverSessionFilePath = args[2]
	}
	b, _ := ioutil.ReadFile(serverSessionFilePath)
	serverSessionText := string(b)
	b, _ = ioutil.ReadFile(protoFilePath)
	protoText := string(b)
	//fmt.Println(serverSessionText, protoText)
	enums := GetEnumFromProto(protoText, "REQUEST_TYPE", true)
	code:=GenCode(serverSessionText, enums)
	fmt.Println(code)
}

func GenCode(src string, requests []string) string {
	genCaseBlockFunc := func(s string) string {
		block := fmt.Sprintln(
			fmt.Sprintf("case message.REQUEST_TYPE_%s:\r\n", s),
			fmt.Sprintf("point:= new(message.%sRequest)\r\n", s),
			"proto.Unmarshal(data,point)\r\n",
			fmt.Sprintf("Handle%sFunc(sess,point)\r\n", s),
		)
		return block
	}

	genFuncBodyFunc := func(caseField string) string {
		block := fmt.Sprintln(
			"func HandleMessage(sess *PlayerSession,data []byte){\r\n",
			"sess.IsConnect=true\r\n",
			"x:=new(message.ProtoMessage)\r\n",
			"proto.Unmarshal(data,x)\r\n",
			"switch message.REQUEST_TYPE(x.MessageType){\r\n",
			caseField,
			"default :\r\n",
			"}\r\n",
			"}\r\n",
		)
		return block
	}
	caseCodes := make([]string, len(requests))
	for i, v := range requests {
		caseCodes[i] = genCaseBlockFunc(v)
	}
	funcStart := strings.Index(src, "func HandleMessage")
	beforeFuncBody := src[0:funcStart]
	joinedCaseBody := strings.Join(caseCodes, "\r\n")
	return beforeFuncBody + genFuncBodyFunc(joinedCaseBody)
}

func GetEnumFromProto(protoContent string, enumName string, removeEqual bool) []string {
	tofindEnumHead := "enum " + enumName
	startIndex := strings.Index(protoContent, tofindEnumHead)
	startStr := protoContent[startIndex:]
	startBraceIndex := strings.Index(startStr, "{")
	endBraceIndex := strings.Index(startStr, "}")
	enums := startStr[startBraceIndex+1:endBraceIndex]
	enums = strings.Replace(enums, "\r\n", "", -1)
	enums = strings.Replace(enums, " ", "", -1)
	enums = strings.TrimSuffix(enums, ";")
	enumA := strings.Split(enums, ";")
	fmt.Println(enumA)
	if !removeEqual {
		return enumA
	} else {
		for i, v := range enumA {
			strLen := len(v) - 2
			enumA[i] = v[0:strLen]
		}
	}
	fmt.Println(enumA)
	return enumA
}
