package main

import (
	"io/ioutil"
	"testing"
	"time"
)

func Test001(t *testing.T) {
	url := "https://habr.com/"
	timeout := 5 * time.Second
	output_path := "sldflsdfsmvsmv"
	contex, _ := HTTPGet(url, timeout)
	err := ioutil.WriteFile(output_path, contex, 0644)
	if err == nil {
		t.Errorf("Error wrong way to file")
	}
}

func Test002(t *testing.T) {
	url := "https://habr.com/"
	timeout := 5 * time.Second
	_, err := HTTPGet(url, timeout)

	if err != nil {
		t.Errorf("The error not have to exist")
	}
}

func Test003(t *testing.T) {
	url := "asdkasdlsak"
	timeout := 5 * time.Second
	_, err := HTTPGet(url, timeout)

	if err == nil {
		t.Errorf("Error ulr have to exist")
	}
}

func Test004(t *testing.T) {
	url := "https://habr.com/"
	timeout := 5 * time.Second
	contex, _ := HTTPGet(url, timeout)

	if contex == nil {
		t.Errorf("Error parse site, need a body page")
	}
}

func Test005(t *testing.T) {
	url := "https://habr.com/"
	timeout := 5 * time.Second
	output_path := "test.txt"
	contex, _ := HTTPGet(url, timeout)
	err := ioutil.WriteFile(output_path, contex, 0644)
	if err != nil {
		t.Errorf("Error write file")
	}
}
