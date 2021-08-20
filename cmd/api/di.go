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
	slot *slots

	httpRouter func()(http.Handler,error)
	insertTicket func()(*ticketRepository,error)
}

func newDIContainer(capacity int) *diContainer {
	diC := &diContainer{}
	diC.slot = newSlots(capacity)
	diC.CommonDIC = Sample.NewCommonDIContainer()
	diC.httpHandlers = newHTTPHandlers(diC)
	diC.httpRouter =  newHTTPRouterDIProvider(diC)
	diC.insertTicket = newTicketRepositoryDIProvider(diC)
	return diC
}

