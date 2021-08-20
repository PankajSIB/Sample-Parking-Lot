package main

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)
const DbName = "parkinglot"
const collection = "ticket"

type TicketInfo struct {
	carRegNo 	string		`json:"carRegNo"`
	entryTime 	time.Time	`json:"entryTime"`
	exitTime	time.Time	`json:"exitTime"`
	slot 		int			`json:"slot"`
	charges		float64		`json:"charges"`
}

type ticketRepository struct{
	mongoColl *mongo.Client
}

func newTicketRepository(diC *diContainer) (*ticketRepository,error) {
	dbClient,err := diC.CommonDIC.MongoDBClient()
	if err != nil {
		return nil,errors.Wrap(err,"Mongo DB Client")
	}
	return &ticketRepository{
		mongoColl: dbClient,
	},nil
}

func newTicketRepositoryDIProvider(dic *diContainer) func()(*ticketRepository,error) {
	var t *ticketRepository
	var err error
	var mu sync.Mutex
	return func()(*ticketRepository,error) {
		mu.Lock()
		defer mu.Unlock()
		if t == nil{
			t,err=newTicketRepository(dic)
		}
		return t,err
	}
}

func (t *ticketRepository)insertTicket(ticketInfo *TicketInfo)(err error){
	_,err = t.mongoColl.Database(DbName).Collection(collection).
		InsertOne(
			context.Background(),
			ticketInfo,
		)
	return err
}