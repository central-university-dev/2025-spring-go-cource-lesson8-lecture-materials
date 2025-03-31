package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	// NewRequestWithContext
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://google.com/robots.txt", nil)
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second, // Тайм-аут подключения
		}).DialContext,
		TLSHandshakeTimeout:   3 * time.Second, // Тайм-аут для TLS Handshake
		ResponseHeaderTimeout: 2 * time.Second,
	}

	c := http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // Общий тайм-аут
	}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
