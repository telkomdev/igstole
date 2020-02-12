package httplib

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// HTTPGet will return http response, and error otherwise
func HTTPGet(url string) (*http.Response, error) {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		//IdleConnTimeout:     10 * time.Second,
	}

	client := http.Client{
		Timeout:   time.Second * 5,
		Transport: transport,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	userAgent := randomUserAgent()

	fmt.Println(userAgent)

	request.Header.Set("User-Agent", userAgent)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
