package model

import "github.com/jinzhu/gorm"

// 定义用户信息
//type User struct {
//	gorm.Model
//	Username  string `json:"username"`
//	Password  string `json:"password"`
//}

// 定义邮箱手机号信息
type UserRegistration struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
