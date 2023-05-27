package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type httpstruct struct {
	user string
	pass string
}

var (
	username = "nsroot"
	password = ""
)

func main() {
	m := httpstruct{user: "nsroot", pass: "6ni$$3<0Rp"}
	call("https://10.170.0.45/nitro/v1/config/lbvserver", "GET", &m)
}

func call(url, method string, val *httpstruct) error {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	client := &http.Client{
		Timeout: time.Second * 10, Transport: transCfg,
	}

	req, err := http.NewRequest(method, url, nil)
	//fmt.Println(req)
	// if err != nil {
	// 	//fmt.Println("Error is Nil")
	// 	//return fmt.Errorf("Got error %s", err.Error())
	// 	fmt.Println(err)
	// }
	req.SetBasicAuth(val.user, val.pass)
	response, err := client.Do(req)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	//decoder := json.NewDecoder()
	//fmt.Println("Response:", response)
	//fmt.Println("Decoded Response:", decoder)
	fmt.Println("Body:", string(body))
	fmt.Println("Error:", err)
	// if err != nil {
	// 	return fmt.Errorf("Got error %s", err.Error())
	// }

	defer response.Body.Close()
	return nil
}
