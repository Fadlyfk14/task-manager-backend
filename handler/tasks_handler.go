package handler

import (
	"beckend/model"
	"beckend/repository"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ================= GET ALL TASK =================
func GetAllTask(c *fiber.Ctx) error {
	data, err := repository.GetAllTask()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data task",
		"data":    data,
	})
}

// ================= INSERT TASK =================
func InsertTask(c *fiber.Ctx) error {
	var task model.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	if task.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Title wajib diisi",
		})
	}

	if err := repository.InsertTask(&task); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menambahkan task",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Task berhasil ditambahkan",
		"data":    task,
	})
}

// ================= GET TASK BY ID =================
func GetTaskByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	data, err := repository.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Task tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil detail task",
		"data":    data,
	})
}

// ================= UPDATE TASK =================
func UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	var input model.Task
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	data, err := repository.UpdateTask(uint(id), input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal update task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task berhasil diupdate",
		"data":    data,
	})
}

// ================= DELETE TASK =================
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	err = repository.DeleteTask(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Task tidak ditemukan",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task berhasil dihapus",
	})
}
func UpdateTaskStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
		})
	}

	type Request struct {
		Status string `json:"status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	// 🔥 NORMALISASI
	req.Status = strings.ToLower(req.Status)

	if req.Status != "todo" && req.Status != "progress" && req.Status != "done" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Status tidak valid",
		})
	}

	data, err := repository.UpdateTaskStatus(uint(id), req.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal update status",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Status berhasil diupdate",
		"data":    data,
	})
}
