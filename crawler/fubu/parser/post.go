package parser

import (
	"Cookbook/crawler/engine"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

const postRe  = ``

type PostParser struct {
	title string
}
type PostItem struct {
	item engine.Item
}

func (p PostParser) Parse(contents []byte, url string) engine.ParseResult {
	log.Println(p.title)

	return postCssSelector(string(contents))
}


func NewPostParser() *PostParser {
	return &PostParser{}
}



func ParsePostContent(contents []byte,_ string) engine.ParseResult  {
	re := regexp.MustCompile(postRe)
	matches := re.FindAllSubmatch(contents,-1)

	//result := engine.ParseResult{}

	log.Println(matches)

	//for _, m:= range matches{
	//
	//	//result.Requests = append(result.Requests,engine.Request{
	//	//	Url:    "",
	//	//	Parser: nil,
	//	//})
	//}
	return engine.ParseResult{}

}

func postCssSelector(contents string) engine.ParseResult {

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err!=nil{
		log.Fatalln(err)
	}
	result := engine.ParseResult{}
	//log.Println(contents)
	//var j int = 0
	dom.Find("body > center > div:nth-child(6) > form > div").Each(func(i int, selection *goquery.Selection) {

		//#message62422883 > font
		//table > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(2) > td > div:nth-child(1)
		//message62422917
		// body > center > div:nth-child(6) > form > div:nth-child(2) > table > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(2) > td
		selection.Find("table > tbody > tr > td:nth-child(2) > table > tbody > tr:nth-child(2) > td > div").Each(func(j int, d *goquery.Selection){

			t1 := strings.TrimSpace(d.Text())
			if len(t1) > 0{
				t1 =strings.Replace(t1," ","",-1)
				//log.Printf("%v",t1)
				log.Printf("get %d楼 post content",i)
			}
		})



	})
	return result
}
