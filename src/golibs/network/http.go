package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main() {
	// obtain content by http get
	resp, _ := http.Get("http://www.baidu.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()

	// basic http get, use proxy
	proxy_addr := "http://192.168.1.123:8000"
	url := "http://www.baidu.com"
	url_proxy, _ := url.URL{}.Parse(*proxy_addr)
	transport := &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err = client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		_ = robots
	}

	// start web
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)

	// use tcp connection
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}
