package model

import (
	"github.com/apiserver/pkg/auth"
	validator "github.com/go-playground/validator/v10"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

//验证字段
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

//加密用户密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

//新增用户
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}
