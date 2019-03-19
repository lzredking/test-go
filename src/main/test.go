package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testgo"
)

var i int

func main() {
	//语法与Java有区别，同C
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
	if i == 0 {
		i = 2
	}
	fmt.Println("Hello, World!", i)

	kv := map[string]string{"a": "av", "b": "bv"}

	for k, v := range kv {
		fmt.Println(k, v)
	}

	h := testgo.Hello()
	fmt.Println(h)
	///////////////////////////////
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	////////////////////////////
	//返回的body是byte数组
	//byte转为string
	fmt.Println(string(body[:]))
	//string 转为[]byte
	var str string = "test"
	var data []byte = []byte(str)
	fmt.Println(data)
	////////////////////////

	byte, err := ioutil.ReadFile("C://logs/spring.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byte[:]))

	//	data []byte= []byte("这一句话放在哪里？")
	err = WriteFile("C://logs/spring.log", data, 0666)
	fmt.Println(err)
}

//执行顺序在main()之前
func init() {
	str := "test" //函数内变量
	fmt.Println("this is init...", str)

	a := 0
	b := 10
	for a < b {
		a++
		fmt.Println(a)
	}

	fmt.Println(more("T-a", "T-b"))
}

//返回多个值  a,b是入参  后一个括号是返回类型
func more(a, b string) (string, string) {
	return a, b
}

/**
*在文件后面追加，改造自ioutil.WriteFile.参照网络
 */
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
