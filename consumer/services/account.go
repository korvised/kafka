package services

import (
	"encoding/json"
	"fmt"
	"gihub.com/korvised/kafka-consumer/repositories"
	events "github.com/korvised/kafka-events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountHandler(accountRepo repositories.AccountRepository) EventHandler {
	return &accountEventHandler{accountRepo}
}

func (s *accountEventHandler) Handle(topic string, eventBytes []byte) {
	fmt.Println("Topic: ", topic)
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := new(events.OpenAccountEvent)
		if err := json.Unmarshal(eventBytes, event); err != nil {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", event)
		bankAccount := &repositories.BackAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}
		if err := s.accountRepo.Save(bankAccount); err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := new(events.DepositFundEvent)
		if err := json.Unmarshal(eventBytes, event); err != nil {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", event)
		bankAccount, err := s.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance += event.Amount
		if err = s.accountRepo.Save(bankAccount); err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := new(events.WithdrawFundEvent)
		if err := json.Unmarshal(eventBytes, event); err != nil {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", event)
		bankAccount, err := s.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance -= event.Amount
		if err = s.accountRepo.Save(bankAccount); err != nil {
			log.Println(err)
			return
		}
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := new(events.CloseAccountEvent)
		if err := json.Unmarshal(eventBytes, event); err != nil {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", event)
		if err := s.accountRepo.Delete(event.ID); err != nil {
			return
		}
	default:
		log.Println("no event handler")
	}
}
