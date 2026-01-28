package handler

import (
	"beckend/model"
	"beckend/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Format data salah"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "username dan password wajib diisi"})
	}
	// cek username sudah ada atau belum
	if _, err := repository.FindUserByUsername(req.Username); err == nil {
		return c.Status(409).JSON(fiber.Map{"message": "username sudah digunakan"})
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal hash password",
			"error": err.Error()})
	}
	newUser := model.User{
		Username: req.Username,
		Password: string(hash),
	}
	created, err := repository.CreateUser(newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat user",
			"error": err.Error()})
	}
	// jangan balikin password
	return c.Status(201).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data": fiber.Map{
			"id":       created.ID,
			"username": created.Username,
			"role":     created.Role,
		},
	})
}
