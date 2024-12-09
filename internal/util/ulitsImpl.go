package util

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgconn"
	"time"
)

// ExecuteRepeat
// Executes function every <delay> minutes
func ExecuteRepeat(action func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = action(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return
}

// SendPostgresError
// Sends JSON error message to models
// or saves information on disk
func SendPostgresError(err error) error {
	var pgErr *pgconn.PgError

	if errors.Is(err, pgErr) {
		errors.As(err, &pgErr)
		return fmt.Errorf("db error: %s, internals: %s, at: %s, state: %s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState())
	}

	return err
}

// SendServerError
// Sends JSON error message to models
// or saves information on disk
func SendServerError(context fiber.Ctx, attempts int, err error) error {
	return context.Status(attempts).JSON(fiber.Map{
		"message": err.Error(),
	})
}
