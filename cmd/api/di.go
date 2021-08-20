package main

import (
	"github.com/Sample-Parking-Lot"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type diContainer struct {
	CommonDIC interface {
		MongoDBClient() (*mongo.Client, error)
	}
	httpHandlers *httpHandlers

	httpRouter func()(http.Handler,error)
}

func newDIContainer() *diContainer {
	diC := &diContainer{}
	diC.CommonDIC = Sample.NewCommonDIContainer()
	diC.httpHandlers = newHTTPHandlers()
	diC.httpRouter =  newHTTPRouterDIProvider(diC)
	return diC
}

