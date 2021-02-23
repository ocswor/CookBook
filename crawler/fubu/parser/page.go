package parser

import (
	"Cookbook/crawler/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

const pageRe  = ``
// Ctrl + i
type PageParse struct {

}

func (p PageParse) Parse(contents []byte, url string) engine.ParseResult {

	return cssSelector(string(contents))

}

func re_find(contents []byte)  {

	re := regexp.MustCompile(pageRe)
	matches := re.FindAllSubmatch(contents,-1)
	for _, m:= range matches{
		fmt.Println(m)

	}
}

func cssSelector(contents string) engine.ParseResult {

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err!=nil{
		log.Fatalln(err)
	}
	result := engine.ParseResult{}
	dom.Find("body > center > form > div:nth-child(2) > div").Find("tr > td.f_title").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a").Each(func(i int, a *goquery.Selection){
			if i ==0{
				//log.Println(strings.Repeat("XXXXXXX XXXXXXX", 10))
				a_value,exist := a.Attr("href")
				a_text := a.Text()
				if exist{
					url := "https://bbs.fobshanghai.com/" + a_value
					log.Println(url,a_text)
					result.Requests = append(
						result.Requests, engine.Request{
							Url: url,
							Parser: PostParser{title:a_text},
						})
				}else {
					log.Println("not exist")
				}

			}
		})

	})
	return result
}


