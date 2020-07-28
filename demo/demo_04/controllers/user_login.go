package controllers

import (
	"DayDayUp/demo/demo_04/db"
	"DayDayUp/demo/demo_04/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

// 用户创建函数
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	// phone   字符串转为int64
	phone, _ := strconv.ParseInt(c.PostForm("phone"), 10, 64)
	// 进行密码加密处理
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	userInfo := model.UserRegistration{
		Username: username,
		Password: string(hash),
		Phone:    phone,
		Email:    email,
	}
	db.DB.Create(&userInfo)
	//db.Create1(&userInfo)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user create success!",
	})
}

// 用户登录函数
func LoginUser(c *gin.Context) {
	var usersInfo []model.UserRegistration
	//var todos []todoResponse
	//fmt.Println("fdfdsfsdf")
	username := c.PostForm("username")
	password := c.PostForm("password")
	//db.DB.Find(&usersInfo)
	db.DB.Where("username = ?", username).First(&usersInfo)
	if len(usersInfo) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "用户名不存在！",
		})
		return
	}
	//var _todos []model.UserInfo
	for _, item := range usersInfo {
		//_todos = append(_todos, model.UserInfo{
		err := bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(password)) //验证（对比）
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "密码错误！",
			})
		} else {
			if username == item.Username {
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "恭喜你，用户登录成功！",
				})
			}
		}
		//if username == item.Username  &&  {
		//	c.JSON(http.StatusOK,gin.H{
		//		"status": http.StatusOK,
		//		"message":"恭喜你，用户登录成功！",
		//	})
		//	break
		//}else{
		//	if username != item.Username{
		//		c.JSON(http.StatusNotFound,gin.H{
		//			"status": http.StatusNotFound,
		//			"message": "账号错误！",
		//		})
		//	}
		//	if password !=item.Password{
		//		c.JSON(http.StatusNotFound,gin.H{
		//			"status": http.StatusNotFound,
		//			"message": "密码错误！",
		//		})
		//	}
		//	//c.JSON(http.StatusNotFound,gin.H{
		//	//	"status":http.StatusNotFound,
		//	//	"message":	"Login error!",
		//	//})
		//}
	}
}

// 修改用户信息
func UpdateUser(c *gin.Context) {
	var users model.UserRegistration
	todoID := c.Param("id")
	db.DB.First(&users, todoID)
	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo found!",
		})
		return
	}
	password := c.PostForm("password")
	phone, _ := strconv.ParseInt(c.PostForm("phone"), 10, 64)
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	//db.DB.Model(&users).Update("username", c.PostForm("username"))
	//db.DB.Model(&users).Update("password", string(hash))
	//db.DB.Model(&users).Update("phone", phone)
	//db.DB.Model(&users).Update("email", c.PostForm("email"))

	//db.DB.Model().Omit()
	//  更新多个字段
	//db.DB.Model(&users).Updates(map[string]interface{}{"username": c.PostForm("username") ,"password": string(hash),
	//	"phone": phone,"email":c.PostForm("email")})

	//  修改结构体中的某一个字段，其他字段使用原来的值
	db.DB.Model(&users).Updates(model.UserRegistration{Username: c.PostForm("username"), Password: string(hash),
		Phone: phone, Email: c.PostForm("email")})

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated succefully",
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var users model.UserRegistration
	todoID := c.Param("id")

	db.DB.First(&users, todoID)

	if users.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not todo found!",
		})
		return
	}
	// 逻辑删除
	//db.Delete(&todo)
	// 永久删除
	db.DB.Unscoped().Delete(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo delete succefully",
	})
}
