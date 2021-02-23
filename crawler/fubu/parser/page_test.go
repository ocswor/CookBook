package parser

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"testing"
)

func TestPageParse_Parse(t *testing.T) {

	contents, err := ioutil.ReadFile("page.html")
	if err != nil {
		panic(err)
	}
	utf8contents, err := GbkToUtf8(contents)
	if err != nil {
		t.Errorf("contents error  %s err %v",
			utf8contents, err)
	}
	result := PageParse{}.Parse(utf8contents, "")
	log.Printf("%d",len(result.Requests))
	const resultSize = 50

	if len(result.Requests) != 50 {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}

}
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
