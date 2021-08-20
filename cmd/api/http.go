package main

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"sync"
)

type httpHandlers struct {
	validateapi         func() (http.Handler, error)
}

func newHTTPHandlers()*httpHandlers {
	return &httpHandlers{
		validateapi:newValidateApiDIProvider(),
	}
}

func registerHTTPHandlers(r *mux.Router,hs *httpHandlers)(error){
	for _,v := range []struct {
		name      string
		configure func(*mux.Route) *mux.Route
		handler   func() (http.Handler, error)
	}{
		{
			name: "validateapi",
			configure:configureValidateAPIHTTPRoute,
			handler: hs.validateapi,
		},
	}{
		h,err := v.handler()
		if err != nil {
			return errors.Wrap(err,v.name)
		}
		v.configure(r.NewRoute()).Handler(h)
	}
	return nil
}

func newHTTPRouter(dic *diContainer)(http.Handler,error){
	r := mux.NewRouter()
	err := registerHTTPHandlers(r,dic.httpHandlers)
	if err != nil {
		return nil,errors.Wrap(err,"register handlers")
	}
	return r,nil
}

func newHTTPRouterDIProvider(dic *diContainer) func() (http.Handler, error) {
	var h http.Handler
	var mu sync.Mutex
	return func() (http.Handler, error) {
		mu.Lock()
		defer mu.Unlock()
		var err error
		if h == nil {
			h, err = newHTTPRouter(dic)
		}
		return h, err
	}
}

func handleHTTP(w http.ResponseWriter, req *http.Request, name string, f func(http.ResponseWriter, *http.Request) error) {
	err := f(w, req)
	handleHTTPError(w, req, err, name)
}

func handleHTTPError(w http.ResponseWriter, req *http.Request, err error, name string) {
	if err == nil {
		return
	}
	err = errors.Wrap(err, name)
	err = errors.Wrap(err, "handle HTTP")
}

func runHTTPServer(dic *diContainer, addr string) error {
	r, err := dic.httpRouter()
	if err != nil {
		return errors.Wrap(err, "get router")
	}
	log.Printf("Start HTTP server on %s", addr)
	err = http.ListenAndServe(":"+addr, r)
	if err != nil {
		return errors.Wrap(err, "listen and serve")
	}
	log.Println("Stopped HTTP server")
	return nil
}