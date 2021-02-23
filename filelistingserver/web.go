package main

import (
	"Cookbook/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type myHandler func(writer http.ResponseWriter,
	request *http.Request)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request :%s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}
func myWrapper(handler myHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
func main() {

	//http.HandleFunc("/",errWrapper(filelisting.HandleFileList))
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//
	//	path:=request.URL.Path[len("/"):]
	//	file,err:= os.Open(path)
	//	if err!= nil {
	//		panic(err)
	//	}
	//	defer file.Close()
	//
	//	all,err:= ioutil.ReadAll(file)
	//	if err!=nil{
	//		panic(err)
	//	}
	//	writer.Write(all)
	//})
	http.HandleFunc("/", errWrapper(filelisting.MyHandleFileList))
	log.Printf("server listen 8888")
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}

}
