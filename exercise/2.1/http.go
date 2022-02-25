package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthyz", healthyz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error: %q\n", err)
	}
	//log.Println(http.ListenAndServe(":80", nil))

}

func healthyz(w http.ResponseWriter, r *http.Request) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	for key, value := range r.Header {
		fmt.Printf("Header field key is :%q, value is %q\n", key, value[0])
		w.Header().Set(key, value[0])

	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header yes
	w.Header().Set("Version", os.Getenv("VERSION"))
	io.WriteString(w, "ok")

	//log.Println(w)
	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	//log.Println(r.Header.Get("X-Real-Ip"))
	addr := strings.Split(r.RemoteAddr, `:`)
	log.Println("Client side ip is: ", addr[0])
	log.Println("Client side port is: ", addr[1])
	returnCode := 200
	log.Println("HTTP 返回码:", returnCode)

	//4.当访问 localhost/healthz 时，应返回 200 yes
	w.WriteHeader(returnCode)
}

//Example output:
//Header field key is :"Cache-Control", value is "max-age=0"
//Header field key is :"Upgrade-Insecure-Requests", value is "1"
//Header field key is :"User-Agent", value is "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36 Edg/98.0.1108.51"
//Header field key is :"Accept", value is "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
//Header field key is :"Accept-Encoding", value is "gzip, deflate"
//Header field key is :"Accept-Language", value is "zh,en-US;q=0.9,en;q=0.8,zh-CN;q=0.7"
//Header field key is :"Connection", value is "keep-alive"
//2022/02/26 01:05:59 Client side ip is:  192.168.3.99
//2022/02/26 01:05:59 Client side port is:  49826
//2022/02/26 01:05:59 HTTP 返回码: 200
//2022/02/26 01:05:59 http: superfluous response.WriteHeader call from main.healthyz (http.go:46)
