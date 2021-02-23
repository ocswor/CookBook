package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestPostParse_Parse(t *testing.T) {

	contents, err := ioutil.ReadFile("post.html")
	if err != nil {
		panic(err)
	}
	utf8contents, err := GbkToUtf8(contents)
	if err != nil {
		t.Errorf("contents error  %s err %v",
			utf8contents, err)
	}
	result := PostParser{title:"出"}.Parse(utf8contents, "")
	log.Printf("%d",len(result.Requests))


}
