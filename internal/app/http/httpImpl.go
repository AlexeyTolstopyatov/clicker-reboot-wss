package http

import (
	"clicker/internal/def/models"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Register(fctx *fiber.Ctx) error {
	req := models.RegistryUserModel{}
	ctx := (*fctx).UserContext()

	if err := (*fctx).Bind().Body(&req); err != nil {
		return (*fctx).Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.UserApi.Register(ctx, models.RegistryUserModel{
		TelegramId: req.TelegramId,
		TeamName:   req.TeamName,
	}); err != nil {
		return (*fctx).Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "registration failed",
		})
	}

	return (*fctx).Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "done",
	})
}

func (h *Handler) Login(fctx *fiber.Ctx) error {
	req := models.LoginUserModel{}
	ctx := (*fctx).UserContext()

	if err := (*fctx).Bind().Body(&req); err != nil {
		return (*fctx).Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid telegram id",
		})
	}

	auth, err := h.UserApi.Login(ctx, req.TelegramId)
	if err != nil {
		return (*fctx).Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "authorization failed",
		})
	}

	if auth {
		return (*fctx).Status(fiber.StatusOK).JSON(fiber.Map{})
	}

	return (*fctx).Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
}
