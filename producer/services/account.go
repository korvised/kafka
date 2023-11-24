package services

import (
	"github.com/google/uuid"
	events "github.com/korvised/kafka-events"
	"github.com/korvised/kafka-producer/commands"
	"log"
)

type AccountService interface {
	OpenAccount(command *commands.OpenAccountCommand) (id string, err error)
	DepositFund(command *commands.DepositFundCommand) error
	WithdrawFund(command *commands.WithdrawFundCommand) error
	CloseAccount(command *commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return &accountService{eventProducer}
}

func (s *accountService) OpenAccount(command *commands.OpenAccountCommand) (id string, err error) {
	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	log.Printf("%#v\n", event)

	return event.ID, s.eventProducer.Produce(event)
}

func (s *accountService) DepositFund(command *commands.DepositFundCommand) error {
	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	return s.eventProducer.Produce(event)
}

func (s *accountService) WithdrawFund(command *commands.WithdrawFundCommand) error {
	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	return s.eventProducer.Produce(event)
}

func (s *accountService) CloseAccount(command *commands.CloseAccountCommand) error {
	event := events.CloseAccountEvent{
		ID: command.ID,
	}

	return s.eventProducer.Produce(event)
}
