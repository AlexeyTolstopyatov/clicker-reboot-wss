package util

import "github.com/gofiber/fiber/v3"

type ErrorHandlable interface {
	SendPostgresError(err error) error
	SendServerError(context fiber.Ctx, attempts int, err error) error
}
