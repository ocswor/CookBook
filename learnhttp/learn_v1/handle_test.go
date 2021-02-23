package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandle(t *testing.T)  {


	url := "http://127.0.0.1:89/search/"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	s,_:=ioutil.ReadAll(resp.Body)

	fmt.Printf(string(s))
	//if resp.StatusCode != http.StatusOK {
	//	t.Errorf("Response code is %v", resp.StatusCode)
	//}
	
}

func BenchmarkCreaterUser(b *testing.B) {

}