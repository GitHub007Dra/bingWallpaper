package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func Post(url string, data interface{}) (int, error) {

	// 超时时间：2秒
	client := &http.Client{Timeout: 2 * time.Second}
	jsonStr, err1 := json.Marshal(data)
	if err1 != nil {


	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	//
	//result, _ := ioutil.ReadAll(resp.Body)
	//return string(result)
	return resp.StatusCode, err
}

func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}
