package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:mysql123@tcp(192.168.20.61:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移数据库结构
	db.AutoMigrate(&todoModel{})
	//defer db.Close()
}

type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// 自定义表明
// Set User's table name to be `profiles`
func (todoModel) TableName() string {
	return "todo"
}

// 返回给前端的数据
type todoResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

func main() {
	router := gin.Default()
	todoV1 := router.Group("/api/v1/todos")
	{
		todoV1.GET("/", fetchAllTodo)
		todoV1.GET("/:id", fetchOneTodo)
		todoV1.POST("/", createTodo)
		todoV1.PUT("/:id", updateTodo)
		todoV1.DELETE("/:id", deleteTodo)
	}
	router.Run(":8000")
	defer db.Close()
}

// 添加数据
func createTodo(c *gin.Context) {
	// 从POST 中获取数据保存到mysql
	title := c.PostForm("title")
	// strconv.Atoi 转string 类型
	completed, _ := strconv.Atoi(c.PostForm("completed"))

	todo := todoModel{
		Title:     title,
		Completed: completed,
	}
	// 保存到mysql 数据库
	db.Create(&todo)
	// 返回json数据
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created successfully",
		"todoID":  todo.ID,
	})
	db.Close()
}

// 更新数据
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo found!",
		})
		return
	}
	db.Model(&todo).Update("title", c.PostForm("title"))

	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated succefully",
	})
}
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo found!",
		})
		return
	}
	// 逻辑删除
	//db.Delete(&todo)
	// 永久删除
	db.Unscoped().Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo delete succefully",
	})

}

// 返回所有数据
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	//var todos []todoResponse
	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo data found!",
		})
		return
	}
	var _todos []todoResponse
	for _, item := range todos {
		_todos = append(_todos, todoResponse{
			ID:        item.ID,
			Title:     item.Title,
			Completed: item.Completed,
		})
	}
	// 查询到了数据
	//for _,title := range todos{
	//
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

func fetchOneTodo(c *gin.Context) {
	var todos todoModel

	todoID := c.Param("id")

	db.First(&todos, todoID)

	if todos.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo found!",
		})
		return
	}
	_todo := todoResponse{
		ID:        todos.ID,
		Title:     todos.Title,
		Completed: todos.Completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todo,
	})
}
