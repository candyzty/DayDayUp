package model

// 定义用户信息
//type User struct {
//	gorm.Model
//	Username  string `json:"username"`
//	Password  string `json:"password"`
//}

type UserRegistration struct {
	//User     User `json:"user"`
	//gorm.Model
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
