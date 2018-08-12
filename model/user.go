package model

import (
	"fmt"
	"go_projects/api_server/pkg/auth"
	"go_projects/api_server/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 下午7:22'
*/

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

// crud operations of User
func (u *UserModel) CreateUser() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) UpdateUser() error {
	return DB.Self.Save(&u).Error
}

func DeleteUser(id uint64) error {
	u := &UserModel{}
	u.BaseModel.Id = id
	return DB.Self.Delete(&u).Error
}

func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Encrypt the user password
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// validate the fields
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
