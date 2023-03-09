package main

import (
	"bufio"
	"github.com/hktalent/go-jq/pkg"
	util "github.com/hktalent/go-utils"
	"os"
	"strings"
)

/*
go build -o jq main.go
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "%v:%v" "ip_str" "port"
cat $HOME/MyWork/scan4all/atckData/china_chengdu.json|./jq "ip_str" "port"
./jq $HOME/MyWork/scan4all/atckData/china_chengdu.json "ip_str" "port"
./jq $HOME/MyWork/scan4all/atckData/china_chengdu.json "%v:%v" "ip_str" "port"
*/
func main() {
	//os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/china_chengdu.json", "ip_str"}
	//os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/x01/x01.json"}
	// os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/x01/x01.json", "host"}
	//os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/x01/a0988c54b5a57d258a43a0a95f54e5975aaec96e.xml", "%v:%v", "nmaprun.host.#.address.addr", "nmaprun.host.#.ports.port.#.portid"} //
	//os.Args = []string{"", "/Users/51pwn/MyWork/scan4all/atckData/x01/domains_httpx.json", "%v", "url"} //

	a := os.Args[1:]
	if !util.FileExists(a[0]) {
		if nil != os.Stdin {
			scanner := bufio.NewScanner(os.Stdin)
			pkg.SetBufSize(scanner)
			for scanner.Scan() {
				value := strings.TrimSpace(scanner.Text())
				pkg.DoQuery(value, a...)
			}
		}
	} else {
		pkg.DoQuery(a[0], a[1:]...)
	}
}
