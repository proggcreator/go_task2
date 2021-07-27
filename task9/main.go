package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HTTPGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel_func := context.WithTimeout(context.Background(), timeout)
	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		cancel_func()
		return nil, fmt.Errorf("INVALID RESPONSE; status: %s", response.Status)
	}

	return ioutil.ReadAll(response.Body)
}

func main() {

	url := flag.String("url", "https://qna.habr.com/", "url")
	timeout := flag.Duration("timeout", 7*time.Second, "string1")
	output_path := flag.String("output", "test.txt", "string2")

	flag.Parse()

	content, err := HTTPGet(*url, *timeout)
	if err != nil {
		log.Fatalln("HTTPGET: ", err)
	}
	err = ioutil.WriteFile(*output_path, content, 0666)
	if err != nil {
		log.Fatalln("WriteFile: ", err)
	}
	return
}
