/*
@Time : 2020/5/3 下午10:22
@Author : xiukang
@File : fetch
@Software: GoLand
*/
package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func TestFetch() {
	url := os.Args
	if len(url) <= 1 {
		fmt.Printf("请输入地址")
		return
	}

	host := url[1]
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	fmt.Printf(host)
	res, err := http.Get(host)

	if err != nil {
		fmt.Print(err)
	}

	if res.StatusCode == 200 {
		b, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)

		}
		fmt.Printf("%s", b)
	}

}
