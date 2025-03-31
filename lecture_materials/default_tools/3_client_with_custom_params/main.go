package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	v := url.Values{}
	v.Add("id", "1")
	v.Add("name", "Олег")
	queryString := v.Encode()

	body := bytes.NewBufferString("Hello and welcome!")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://google.com/robots.txt"+"?"+queryString, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "mr. Anderson")
	req.Header.Add("Env", "matrix")

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("request:\n%s\n\n", b)

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	b, err = httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response:\n%s", b)

	//fmt.Println("------------------")

	// bodyResp, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Ошибка при чтении ответа:", err)
	// 	return
	// }
	// fmt.Println(string(bodyResp))

	// fmt.Println("------------------")

}
