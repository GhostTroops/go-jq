package main

import (
	"bufio"
	"fmt"
	util "github.com/hktalent/go-utils"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"os"
	"strings"
)

const (
	MacLineSize = 10 * 1024 * 1024 // 10M
)

// 如果是文件就读取返回
func ReadFile(s string) string {
	if util.FileExists(s) {
		var err error
		if f1, err := os.Open(s); nil == err {
			if data, err := io.ReadAll(f1); nil == err {
				s1 := string(data)
				return s1
			}
		}
		log.Println("ReadFile", s, err)
		return ""
	}
	return s
}

// 单json处理
func DoOneJson(szJson, szFormat string, query ...string) {
	if !gjson.Valid(szJson) {
		log.Println(szJson, " not is json")
	}
	oJ := gjson.Parse(szJson)
	var a []interface{}
	for _, q1 := range query {
		a = append(a, oJ.Get(q1))
	}
	if 0 < len(a) {
		fmt.Printf(szFormat, a...)
	}
}

// default size 10M
func SetBufSize(scanner *bufio.Scanner) {
	scanner.Buffer(make([]byte, MacLineSize), MacLineSize)
}

// 处理单个文件，或单行
func DoQuery(szFile string, query ...string) {
	s := ReadFile(szFile)

	var szFormat = ""
	if -1 < strings.Index(query[0], "%") {
		szFormat = query[0]
		query = query[1:]
	} else {
		szFormat = strings.Repeat("%v,", len(query))
		szFormat = szFormat[0 : len(szFormat)-1]
	}
	szFormat += "\n"

	// 如果失败，就逐行处理
	if gjson.Valid(s) {
		DoOneJson(s, szFormat, query...)
	} else {
		scanner := bufio.NewScanner(strings.NewReader(s))
		SetBufSize(scanner)
		for scanner.Scan() {
			if value := strings.TrimSpace(scanner.Text()); 0 < len(value) {
				DoOneJson(value, szFormat, query...)
			}
		}
		if nil != scanner.Err() {
			log.Println(scanner.Err())
		}
	}
}

/*
go build -o jq main.go
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "%v:%v" "ip_str" "port"
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "ip_str" "port"
./jq $HOME/MyWork/scan4all/atckData/china_chengdu.json "ip_str" "port"
*/
func main() {
	//os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/china_chengdu.json", "ip_str"}
	a := os.Args[1:]
	if !util.FileExists(a[0]) {
		if nil != os.Stdin {
			scanner := bufio.NewScanner(os.Stdin)
			SetBufSize(scanner)
			for scanner.Scan() {
				value := strings.TrimSpace(scanner.Text())
				DoQuery(value, a...)
			}
		}
	} else {
		DoQuery(a[0], a[1:]...)
	}
}
