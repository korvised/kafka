package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/korvised/kafka-producer/commands"
	"github.com/korvised/kafka-producer/response"
	"github.com/korvised/kafka-producer/services"
	"log"
	"net/http"
)

type AccountHandler interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) AccountHandler {
	return &accountHandler{accountService}
}

func (h *accountHandler) OpenAccount(c *fiber.Ctx) error {
	command := new(commands.OpenAccountCommand)

	if err := c.BodyParser(command); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	accountID, err := h.accountService.OpenAccount(command)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, fiber.Map{
		"message": "open account success",
		"data":    accountID,
	})
}

func (h *accountHandler) DepositFund(c *fiber.Ctx) error {
	command := new(commands.DepositFundCommand)

	customerID := c.Params("cus_id")
	command.ID = customerID

	if err := c.BodyParser(command); err != nil {
		log.Println(err)
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.accountService.DepositFund(command); err != nil {
		log.Println(err)
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fiber.Map{
		"message": "deposit fund success",
	})
}

func (h *accountHandler) WithdrawFund(c *fiber.Ctx) error {
	command := new(commands.WithdrawFundCommand)

	customerID := c.Params("cus_id")
	command.ID = customerID

	if err := c.BodyParser(command); err != nil {
		log.Println(err)
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.accountService.WithdrawFund(command); err != nil {
		log.Println(err)
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fiber.Map{
		"message": "withdraw fund success",
	})
}

func (h *accountHandler) CloseAccount(c *fiber.Ctx) error {
	command := new(commands.CloseAccountCommand)

	customerID := c.Params("cus_id")
	command.ID = customerID

	if err := h.accountService.CloseAccount(command); err != nil {
		log.Println(err)
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fiber.Map{
		"message": "close account success",
	})
}
