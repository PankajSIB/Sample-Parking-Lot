package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"sync"
	"time"
)

type ticketRequest struct{
	CarRegistrationNo string `json:"carRegNO"`
}

type postTicketHandler struct{
	getSlot func()(int,error)
	insertTicket func(*TicketInfo)error
}

func newPostTicketHandler(diC *diContainer)(*postTicketHandler,error){
	insertTicket,err :=diC.insertTicket()
	if err != nil {
		return nil,errors.Wrap(err, "InsertTicket")
	}

	return &postTicketHandler{
		insertTicket: insertTicket.insertTicket,
		getSlot: diC.slot.getSlots,
	},nil
}

func newPostTicketHandlerDIProvider(dic *diContainer)func()(http.Handler,error){
	var p *postTicketHandler
	var err error
	var mu sync.Mutex
	return func()(http.Handler,error){
		mu.Lock()
		defer mu.Unlock()
		if p == nil{
			p,err = newPostTicketHandler(dic)
		}
		return p,err
	}
}

func configurePostTicketHTTPRoute(r *mux.Route)*mux.Route{
	return r.Methods(http.MethodPost).Path("/ticket")
}

func(p *postTicketHandler)ServeHTTP(w http.ResponseWriter,req *http.Request) {
	handleHTTP(w,req,"PostTicketHandler",p.handle)
}

func(p *postTicketHandler)handle(w http.ResponseWriter,req *http.Request,) error {
	tR := ticketRequest{}
	err := json.NewDecoder(req.Body).Decode(&tR)
	if err != nil {
		return errors.Wrap(err,"Decoder")
	}

	slot,err := p.getSlot()
	if err != nil {
		return errors.Wrap(err,"get slot")
	}

	ticketInfo := TicketInfo{
		carRegNo: tR.CarRegistrationNo,
		entryTime: time.Now(),
		slot: slot,
	}

	err = p.insertTicket(&ticketInfo)
	if err != nil {
		return errors.Wrap(err,"insert Ticket")
	}

	return nil
}