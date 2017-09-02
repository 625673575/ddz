package main

import "os"
import (
	"io/ioutil"
	"strings"
	"fmt"
)

var protoFilePath string
var serverSessionFilePath string
var csNetFilePath string

func main() {
	args := os.Args
	if (len(args) > 2) {
		protoFilePath = args[1]
		serverSessionFilePath = args[2]
		csNetFilePath=args[3]
	}
	b, _ := ioutil.ReadFile(serverSessionFilePath)
	serverSessionText := string(b)
	b, _ = ioutil.ReadFile(protoFilePath)
	protoText := string(b)
	fmt.Println(serverSessionText, protoText)
	GenClientCode(protoText)
	GenServerCode(protoText,serverSessionText)
}
func GenClientCode(protoText string){
	enums := GetEnumFromProto(protoText, "REPLY_TYPE", 2)
	code:=GenClientGetReplyCode(enums)
	WriteFile(csNetFilePath,code)
}
func GenServerCode(protoText string, serverSessionText string) {
	enums := GetEnumFromProto(protoText, "REQUEST_TYPE", 1)
	code := GenGoSessionCode(serverSessionText, enums)
	fmt.Println(code)
	WriteFile(serverSessionFilePath,code)
}
func WriteFile(path string, content string) {
	file6, error := os.Create(path)
	if error != nil {
		fmt.Println(error)
	}
	file6.WriteString(content)
	file6.Close()
}
func GenGoSessionCode(src string, requests []string) string {
	genCaseBlockFunc := func(s string) string {
		block := fmt.Sprintln(
			fmt.Sprintf("case message.REQUEST_TYPE_%s:\r\n", s),
			fmt.Sprintf("point:= new(message.%sRequest)\r\n", s),
			"proto.Unmarshal(x.Data,point)\r\n",
			fmt.Sprintf("Handle%sFunc(sess,point)", s),
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
		if strings.HasPrefix(v, "NO_") || strings.HasPrefix(v, "No_") || strings.HasPrefix(v, "no_") {
			continue
		}

		caseCodes[i] = genCaseBlockFunc(v)
	}
	funcStart := strings.Index(src, "func HandleMessage")
	beforeFuncBody := src[0:funcStart]
	joinedCaseBody := strings.Join(caseCodes, "\r\n")
	return beforeFuncBody + genFuncBodyFunc(joinedCaseBody)
}

func GetEnumFromProto(protoContent string, enumName string, removeRule int) []string {
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
	if removeRule == 1 {
		for i, v := range enumA {
			startIndex := strings.Index(v,"=")
			enumA[i] = v[0:startIndex]
		}
	}
	if removeRule == 2 {
		for i, v := range enumA {
			strStart := strings.Index(v, "_")
			enumA[i] = v[0:strStart]
		}
	}
	return enumA
}

func GenClientGetReplyCode(replys []string)string {
	csNetmanager := `using System;
using UnityEngine;
public static class NetManager
{
   %s
    public static void OnGetData(byte[] obj)
    {
        ProtoMessage message = ProtoMessage.Parser.ParseFrom(obj);
        var messageType = (REPLY_TYPE)message.MessageType;
        switch (messageType)
        {
            %s
        }
    }
    static void SafeInvoke<T>(Action<T> callback, T obj0)
    {
        if (callback != null)
            callback.Invoke(obj0);
    }
}`
	GetCallBack := func(name string) string {
		return fmt.Sprintf("public static Action<%sReply> On%sReply;", name, name)
	}
	GetCaseBody := func(name string) string {
		resultVarName:=strings.ToLower( name)+"_reply"
		return fmt.Sprintf(`case REPLY_TYPE.%sResult:
			var %s = %sReply.Parser.ParseFrom(message.Data);
			SafeInvoke(On%sReply, %s);
			Debug.Log("收到消息  "+messageType + %s);
			break;`, name,resultVarName,name, name,resultVarName,resultVarName)
	}
	callbackTable := make([]string, len(replys))
	caseTable := make([]string, len(replys))
	for i, v := range replys {
		if v=="NO"|| v== "No"{
			continue
		}
		callbackTable[i] = GetCallBack(v)
		caseTable[i] = GetCaseBody(v)
	}
	joinedCallback := strings.Join(callbackTable, "\r\n")
	joinedCase := strings.Join(caseTable, "\r\n")
	code := fmt.Sprintf(csNetmanager, joinedCallback, joinedCase)
	return code
}
