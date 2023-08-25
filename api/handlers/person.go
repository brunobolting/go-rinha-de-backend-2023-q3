package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	entity "github.com/brunobolting/go-rinha-backend/domain"
	person "github.com/brunobolting/go-rinha-backend/usecase"
	"github.com/gofiber/fiber/v2"
)

func MakePersonHandlers(app *fiber.App, s *person.Service) {
    app.Get("/pessoas/:id", func(c *fiber.Ctx) error {
		return getPerson(c, s)
	})

	app.Post("/pessoas", func(c *fiber.Ctx) error {
		return createPerson(c, s)
	})

	app.Get("/pessoas", func(c *fiber.Ctx) error {
		return findPerson(c, s)
	})

	app.Get("/contagem-pessoas", func(c *fiber.Ctx) error {
		return countPerson(c, s)
	})
}

func getPerson(c *fiber.Ctx, s *person.Service) error {
	p, err := s.GetPerson(c.Params("id"))

	if errors.Is(err, entity.ErrEntityNotFound) {
		return c.Status(404).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	return c.Status(200).Send(payload)
}

type Payload struct {
	Apelido string `json:"apelido"`
	Nome string `json:"nome"`
	Nascimento string `json:"nascimento"`
	Stack []string `json:"stack"`
}

func createPerson(c *fiber.Ctx, s *person.Service) error {
	var payload *Payload

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	p, err := s.CreatePerson(payload.Apelido, payload.Nome, payload.Nascimento, payload.Stack)

	if errors.Is(err, entity.ErrInvalidNickname) || errors.Is(err, entity.ErrInvalidName) {
		return c.Status(422).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	if errors.Is(err, entity.ErrInvalidBirthdate) || errors.Is(err, entity.ErrInvalidStack) {
		return c.Status(422).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	c.Set("Location", fmt.Sprintf("/pessoas/%s", p.ID.String()))
	return c.Status(201).SendString("{\"status\": \"created\"}")
}

func findPerson(c *fiber.Ctx, s *person.Service) error {
	param := c.Query("t")
	if param == "" {
		return c.Status(400).SendString("{\"error\": \"t query is not provided\"}")
	}

	p, err := s.FindPerson(param)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	return c.Status(200).Send(payload)
}

func countPerson(c *fiber.Ctx, s *person.Service) error {
	v, err := s.CountPerson()
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("{\"error\": %s}", strconv.Quote(err.Error())))
	}

	return c.Status(200).SendString(fmt.Sprintf("{\"total\": %d}", v))
}
