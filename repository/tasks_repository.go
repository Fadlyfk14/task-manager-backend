package repository

import (
	"beckend/config"
	"beckend/model"

	"gorm.io/gorm"
)

// ================= GET ALL =================
func GetAllTask() ([]model.Task, error) {
	var tasks []model.Task
	err := config.DB.Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// ================= INSERT =================
func InsertTask(task *model.Task) error {
	return config.DB.Create(task).Error
}

// ================= GET BY ID =================
func GetTaskByID(id uint) (model.Task, error) {
	var task model.Task
	err := config.DB.First(&task, id).Error
	return task, err
}

// ================= UPDATE FULL =================
func UpdateTask(id uint, input model.Task) (model.Task, error) {
	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		return task, err
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Deadline = input.Deadline
	task.Status = input.Status

	err := config.DB.Save(&task).Error
	return task, err
}

// ================= DELETE =================
func DeleteTask(id uint) error {
	result := config.DB.Delete(&model.Task{}, id)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
func UpdateTaskStatus(id uint, status string) (model.Task, error) {
	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		return task, err
	}

	task.Status = status

	err := config.DB.Save(&task).Error
	return task, err
}

// ================= UPDATE TASK (EDIT DATA SAJA - DTO) =================
func UpdateTaskDTO(id uint, input struct {
	Title       string
	Description string
	Deadline    string
}) (model.Task, error) {

	var task model.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		return task, err
	}

	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}
	if input.Deadline != "" {
		task.Deadline = input.Deadline
	}

	err := config.DB.Save(&task).Error
	return task, err
}
