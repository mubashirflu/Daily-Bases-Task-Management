package controllers

import (
	"net/http"
	"strconv"
	"task-management/database"
	"task-management/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Can not created the task",
		})
		return
	}
	if err := database.DB.Create(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sorry Internal Server error",
		})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}
func GetDataTask(ctx *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can not Get the Tasks",
		})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
func UpdateTask(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	var updateData models.Task

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	task.Title = updateData.Title
	task.Description = updateData.Description
	task.Complete = updateData.Complete
	task.ReminderEnabled = updateData.ReminderEnabled
	task.ReminderAt = updateData.ReminderAt

	if err := database.DB.Save(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update task",
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
func GetScheduledTasks(ctx *gin.Context) {
	var tasks []models.Task

	if err := database.DB.
		Where("reminder_enabled = ?", true).
		Order("reminder_at ASC").
		Find(&tasks).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not get scheduled tasks",
		})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
func DeleteTask(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not delete task",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
