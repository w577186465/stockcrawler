package crawler

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	// "bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/bitly/go-simplejson"
	"io"
	"net/url"
	"regexp"
	"strings"
)

type RequestOption struct {
	Method      string
	Url         string
	Head        http.Header
	ConnTimeout time.Duration
	Timeout     time.Duration
	Times       int
	Delay       time.Duration
	Char        string
	Data        url.Values
	Retrytime   int
}

type Response struct {
	Response *http.Response
}

type Reader struct {
	Reader io.Reader
}

// simple downloader
func Download(url string) *Reader {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil) // 发起请求

	host := gethost(url) // get host

	req.Header.Add("Host", host)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")

	// 请求失败重新请求
	var reqerr error
	var resp *http.Response
	retry := 3 // 重新请求次数
	for i := 0; i < retry; i++ {
		resp, reqerr = client.Do(req)
		if reqerr != nil {
			continue
		}
	}
	if reqerr != nil {
		fmt.Println("网页打开失败")
		fmt.Println(reqerr)
		return nil
	}

	response := Reader{
		Reader: resp.Body,
	}

	return &response
}

// 获取url的host
func gethost(url string) string {
	a1 := strings.Split(url, "//")[1]
	return strings.Split(a1, "/")[0]
}

func (reader *Reader) Json() (*simplejson.Json, error) {
	// defer reader.Reader.Close()
	return simplejson.NewFromReader(reader.Reader)
}

/*
	jsonp 转 json
	正则提取出jsonp的json
*/
func (reader *Reader) Jsonp() (*simplejson.Json, error) {
	// defer reader.Reader.Close()
	body, _ := ioutil.ReadAll(reader.Reader)

	reg := regexp.MustCompile(`^[^\[{]*([\[{][\s\S]*?[\]}])[^\]}]*$`) // 提取json正则表达式
	match := reg.FindSubmatch(body)                                   // 提取json
	if len(match) < 2 {
		return nil, errors.New("jsonp提取json失败，正则无法匹配")
	}

	return simplejson.NewJson(match[1])
}

func (reader *Reader) Html() (*goquery.Document, error) {
	return nil, errors.New("response is nil")
	if reader.Reader == nil {
		return nil, errors.New("response is nil")
	}
	fmt.Println(reader)
	return goquery.NewDocumentFromReader(reader.Reader)
}

func (option *RequestOption) Download() *Reader {
	result := &Reader{}
	// 默认设置
	switch {
	case option.Method == "":
		option.Method = "GET"
	case option.ConnTimeout == 0:
		option.ConnTimeout = 10 // 连接超时
	case option.Timeout == 0:
		option.Timeout = 15 // 传输超时
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*option.ConnTimeout) // 设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(option.Timeout * time.Second)) // 设置发送接收数据超时
				return c, nil
			},
		},
	}

	// form
	formio := option.Data.Encode()
	form := strings.NewReader(formio)

	req, err := http.NewRequest(option.Method, option.Url, form) // 发起请求
	if err != nil {
		fmt.Println(err)
		result.Reader = nil
		return &Reader{Reader: nil}
	}

	// header
	// req.Header = req.Head
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 请求失败重新请求
	var reqerr error
	var resp *http.Response
	retry := 3 // 重新请求次数
	for i := 0; i < retry; i++ {
		fmt.Printf("请求失败重试：%d\r\n", i+1)
		resp, reqerr = client.Do(req)
		if reqerr != nil {
			continue
		}
	}

	if reqerr != nil {
		fmt.Println("网页打开失败")
		fmt.Println(reqerr)
		return &Reader{Reader: nil}
	}

	var rd io.Reader
	if option.Char != "" && option.Char != "utf-8" {
		dec := mahonia.NewDecoder(option.Char)

		rd = dec.NewReader(resp.Body)
	} else {
		rd = resp.Body
	}

	return &Reader{rd}
}

func Request(url string) *RequestOption {
	return &RequestOption{Url: url}
}

func (req *RequestOption) Transcoding(char string) *RequestOption {
	req.Char = char
	return req
}

func (req *RequestOption) Retry(n int) *RequestOption {
	req.Retrytime = n
	return req
}
