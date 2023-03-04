package pkg

import (
	"bufio"
	"fmt"
	util "github.com/hktalent/go-utils"
	xj "github.com/hktalent/goxml2json"
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

func GetEq(a, b string) string {
	a1 := []byte(a)
	a2 := []byte(b)
	a3 := []byte{}
	for i, x := range a1 {
		if i > len(a2) {
			break
		}
		if x == a2[i] {
			a3 = append(a3, x)
		}
	}
	return string(a3)
}

func DoPrint(a []interface{}, szFormat string) {
	var nC = 1
	// 维度长度
	for _, x := range a {
		if x1 := x.(gjson.Result).Array(); 0 < len(x1) {
			nC *= len(x1)
		}
	}
	var a1 = make([][]interface{}, nC)
	for _, x := range a {
		if x1 := x.(gjson.Result).Array(); 0 < len(x1) {
			for i := 0; i < nC; i++ {
				a1[i] = append(a1[i], x1[i%len(x1)])
			}
		}
	}
	nC1 := len(a)
	for i := 0; i < nC; i++ {
		if len(a1[i]) == nC1 {
			fmt.Printf(szFormat, a1[i]...)
		}
	}
}

// 单json处理
func DoOneJson(szJson, szFormat string, query ...string) {
	var bHvQuery = 0 < len(query)
	if !gjson.Valid(szJson) {
		log.Println(szJson, " not is json")
	}
	oJ := gjson.Parse(szJson)
	var fnDoOne = func(oJ1 gjson.Result) {
		var a []interface{}
		for _, q1 := range query {
			a = append(a, oJ1.Get(q1))
		}
		if 0 < len(a) {
			DoPrint(a, szFormat)
			//fmt.Printf(szFormat, a...)
		}
	}
	if bHvQuery {
		szlstEq := query[0]
		for i := 1; i < len(query); i++ {
			szlstEq = GetEq(szlstEq, query[i])
		}
		if 0 < len(szlstEq) {
			if strings.HasSuffix(szlstEq, ".") {
				szlstEq = szlstEq[0 : len(szlstEq)-1]
			}
			for i, _ := range query {
				query[i] = query[i][len(szlstEq)+1:]
			}

			if strings.HasSuffix(szlstEq, ".#") {
				szlstEq = szlstEq[0 : len(szlstEq)-2]
				oJ = oJ.Get(szlstEq)
				if oA1 := oJ.Array(); 0 < len(oA1) {
					for _, x := range oA1 {
						fnDoOne(x)
					}
				}
				return
			}
		}
		fnDoOne(oJ)
	} else {
		var m1 = map[string]interface{}{}
		util.Json.Unmarshal([]byte(szJson), &m1)
		//if o := util.GetJson4Query(&m1, ".nmaprun.host.[].ports.extraports.extrareasons"); nil != o {
		//	if oM11, ok := o.(map[string]interface{}); ok {
		//		delete(oM11, "ports")
		//	}
		//}
		// nmaprun.host[].ports.extraports.extrareasons.ports
		if o, err := FormatJson(m1); nil == err {
			fmt.Println(o)
		} else {
			log.Println("FormatJson", err)
		}
	}
}

// default size 10M
func SetBufSize(scanner *bufio.Scanner) {
	scanner.Buffer(make([]byte, MacLineSize), MacLineSize)
}

func FormatJson(data interface{}) (string, error) {
	val, err := util.Json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// 处理单个文件，或单行
func DoQuery(szFile string, query ...string) {
	var s string
	if strings.HasSuffix(szFile, ".xml") {
		xj.AttrPrefix = ""
		xj.ContentPrefix = ""
		if fio, err := os.OpenFile(szFile, os.O_RDONLY, os.ModePerm); nil == err {
			if j1, err := xj.Convert(fio); nil == err {
				s = string(j1.Bytes())
			}
		}
	} else {
		s = ReadFile(szFile)
	}

	var szFormat = ""
	var bHvQuery = 0 < len(query)
	if bHvQuery {
		if -1 < strings.Index(query[0], "%") {
			szFormat = query[0]
			query = query[1:]
		} else {
			szFormat = strings.Repeat("%v,", len(query))
			szFormat = szFormat[0 : len(szFormat)-1]
		}
		szFormat += "\n"
	}

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
