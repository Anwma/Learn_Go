package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, this is home page")
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	//流的设计理念,只读一次
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s \n", string(body)) //1

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	//2
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	body, _ := r.GetBody()
	io.ReadAll(body)

	body, _ = r.GetBody()
	io.ReadAll(body)
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()                 //type Values map[string][]string
	fmt.Fprintf(w, "query is %v\n", values) //query is map[]
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	//json的序列化
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
	//{"Scheme":"","Opaque":"","User":null,"Host":"",
	//"Path":"/wholeUrl","RawPath":"","OmitHost":false,
	//"ForceQuery":false,"RawQuery":"","Fragment":"","RawFragment":""}
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)
	//https-header头部信息
	//header is map[
	//Accept:[text/html,application/xhtml+xml,application/xml;
	//q=0.9,image/avif,image/webp,image/apng,*/*; q=0.8,application/signed-exchange; v=b3; q=0.9]
	//Accept-Encoding:[gzip, deflate, br]
	//Accept-Language:[zh-CN,zh;q=0.9]
	//Connection:[keep-alive]
	//Cookie:[Goland-ad29cbe6=1773aec1-569c-4a43-a3f8-45596eb514dd]
	//Sec-Ch-Ua:["Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"]
	//Sec-Ch-Ua-Mobile:[?0]
	//Sec-Ch-Ua-Platform:["Windows"]
	//Sec-Fetch-Dest:[document]
	//Sec-Fetch-Mode:[navigate]
	//Sec-Fetch-Site:[none]
	//Sec-Fetch-User:[?1]
	//Upgrade-Insecure-Requests:[1]
	//User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36]
	//]
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	//调用之前要先使用一下ParseForm()
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
	//before parse form map[]
	//before parse form map[]
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
