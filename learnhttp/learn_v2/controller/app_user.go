package controller

import (
	"io"
	"log"
	"net/http"
)

type App_User struct {
}

func (h App_User) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	//q := strings.TrimSpace(req.FormValue("q"))
	//log.Printf("%v", req)
	http.Error(w, "",
		http.StatusBadRequest)
	if ok, err := io.WriteString(w, "ok"); err != nil {
		log.Printf("error:%s", err)
	}else {
		log.Print(ok)
	}


}
