package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 判断爬取是否成功
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	// gopm get -g -v golang.org/x/text   编码转换
	// gopm get -g -v golang.org/x/net/html  获取页面编码

	//e := determineEncoding(resp.Body)
	// 页面编码转换成utf8
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	// 获取爬取内容
	//all, err := ioutil.ReadAll(utf8Reader)

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

// 获取页面编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	// ^> 非右括号的字符
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}
	fmt.Println(len(matches))
}
