package main

import (
	"fmt"
	"os"
	"net/http"
	"regexp"
)

func main(){
	if(len(os.Args) > 1){
		method := "GET"
		args := os.Args[1:]
		url := string(args[0])
		if len(args) > 1 && args[1] == "-I" {
			method = "HEAD"
		}
		if ok, _ := regexp.MatchString(`^(http|https)://127.0.0.1:\d+/(\S+|)$`, url); ok {
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				fail(fmt.Sprintf("http.NewRequest: %s", err))
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fail(fmt.Sprintf("http.NewRequest: %s", err))
			}
			if resp.StatusCode != 200 {
				fail(fmt.Sprintf("http.NewRequest status code != 200: %s", resp.StatusCode))
			}
			os.Exit(0)
		}else{
			fail("localhealth: url invalid!")
		}
	}
	os.Exit(1)
}

func fail(msg string){
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}