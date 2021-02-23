package engine

// 每一个 Request 对象 中包含一个 url链接 和 Parser解析器
// Parser解析器解析对应的url ,所以可以需要被实现,所以是一个interface


type Item struct {
	Url     string
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	//Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParserFunc func(
	contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

type NilParser struct{}

func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}