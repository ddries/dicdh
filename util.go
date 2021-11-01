package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var availableApis = [...]string{
	"https://api.ipify.org/",
}

func GetIp() (string, error) {
	var resp *http.Response = nil
	
	for _, api := range availableApis {
		_resp, err := http.Get(api)

		if err != nil {
			continue
		}

		resp = _resp
		break
	}

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		return string(body), nil
	}

	return "", errors.New("could not fetch public ip data")
}

func Quit(text string) {
	fmt.Println(text)
	os.Exit(1)
}