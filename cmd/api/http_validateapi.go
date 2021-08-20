package main

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type validateApiHandler struct{}

func newValidateApi()*validateApiHandler{
	return &validateApiHandler{}
}

func newValidateApiDIProvider() func()(http.Handler,error){
	var v *validateApiHandler
	var mu sync.Mutex
	return func()(http.Handler,error){
		mu.Lock()
		defer mu.Unlock()
		if v == nil {
			v = newValidateApi()
		}
		return v,nil
	}
}

func configureValidateAPIHTTPRoute(r *mux.Route) *mux.Route {
	return r.Methods(http.MethodGet).Path( "/validate-api")
}

func (v *validateApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	handleHTTP(w, req, "validateapi", v.handle)
}

func (v *validateApiHandler) handle(w http.ResponseWriter, req *http.Request)error{
	w.WriteHeader(http.StatusOK)
	_,err := w.Write([]byte("Validation Passed"))
	if err != nil {
		return errors.Wrap(err, "write response")
	}
	return nil
}