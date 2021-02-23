package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
//

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	// check session
	m.r.ServeHTTP(w,r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler  {
	m := middleWareHandler{}
	m.r = r
	return m
}
