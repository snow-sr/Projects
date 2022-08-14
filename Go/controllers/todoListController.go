package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snow-sr/learningGo/database"
	"github.com/snow-sr/learningGo/models"
)

func GetAll(c *gin.Context) {
	var todos []models.TodoItem
	db := database.GetDatabase()

	err := db.Find(&todos).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting todos",
		})
		return
	}

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func GetNonCompleted(c *gin.Context) {
	var todos []models.TodoItem
	db := database.GetDatabase()

	err := db.Where("completed = ?", false).Find(&todos).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting todos",
		})
		return
	}

	c.JSON(200, gin.H{
		"todos": todos,
	})

}

func GetById(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id is not a number",
		})
		return
	}

	db := database.GetDatabase()
	var todo models.TodoItem

	err = db.First(&todo, newId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"message": "todo item not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "get todo item",
		"todo":    todo,
	})

}

func SolveTodo(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id is not a number",
		})
		return
	}

	db := database.GetDatabase()
	var todo models.TodoItem

	err = db.Find(&todo, newId).Update("Completed", true).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error updating todo item",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "todo item solved",
		"todo":    todo,
	})

}

func CreateTodo(c *gin.Context) {
	var todo models.TodoItem
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Create(&todo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "create todo item",
		"todo":    todo,
	})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id is not a number",
		})
		return
	}

	db := database.GetDatabase()
	var todo models.TodoItem

	err = db.First(&todo, newId).Delete(&todo).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting todo item",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "todo item deleted",
		"todo":    todo,
	})
}
