package router

import (
	"beckend/config"
	"beckend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// ================= ROOT =================
	api := app.Group("/api")

	// ================= PUBLIC ROUTES =================
	api.Post("/login", handler.Login)
	api.Post("/register", handler.CreateUser)

	// ================= PROTECTED ROUTES =================
	// Semua route di bawah ini wajib pakai JWT
	task := api.Group("/task", config.JWTMiddleware())

	// TASK CRUD
	task.Get("/", handler.GetAllTask)       // GET    /api/task
	task.Get("/:id", handler.GetTaskByID)   // GET    /api/task/:id
	task.Post("/", handler.InsertTask)      // POST   /api/task
	task.Put("/:id", handler.UpdateTask)    // PUT    /api/task/:id
	task.Delete("/:id", handler.DeleteTask) // DELETE /api/task/:id

	// Update status saja
	task.Patch("/:id/status", handler.UpdateTaskStatus)
}
