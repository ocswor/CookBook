package main

import (
	"Cookbook/crawler/engine"
	"Cookbook/crawler/fubu/parser"
	"Cookbook/crawler/scheduler"
	"log"
)
var fun_2 = engine.FuncParser{
}

//https://www.flysnow.org/2017/05/06/go-in-action-go-log.html
func init(){
	log.SetFlags(log.Ldate|log.Lshortfile)
}
//https://bbs.fobshanghai.com/forum-2-1.html
func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:    "https://bbs.fobshanghai.com/forum-2-1.html",
	//	Parser: parser.PageParse{},
	//
	//})
	
	e:=engine.ConcurrentEngine_One{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:    "https://bbs.fobshanghai.com/forum-2-1.html",
		Parser: parser.PageParse{},
	})

	
}
