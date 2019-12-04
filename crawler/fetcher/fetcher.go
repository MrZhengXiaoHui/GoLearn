package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 判断爬取是否成功
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	// gopm get -g -v golang.org/x/text   编码转换
	// gopm get -g -v golang.org/x/net/html  获取页面编码

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	// 页面编码转换成utf8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 获取爬取内容
	return ioutil.ReadAll(utf8Reader)

}

// 获取页面编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Panicf("获取编码 error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
