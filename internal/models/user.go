package models

import (
	"errors"
	"job-board-api/cmd"
	pkg "job-board-api/pkg/crypto"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"username"`
	Password  string `json:"-" gorm:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterUser struct {
	gorm.Model
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}

type LoginUser struct {
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}

func (RegisterUser) TableName() string {
	return "users"
}

func (l *RegisterUser) Signup() (*RegisterUser, error) {

	user, _ := GetUserByUsername(l.Username)
	if user != nil {
		return nil, errors.New("User Already Exists")
	}
	l.Password, _ = cmd.Http.Hash.Create(l.Password)
	cmd.Http.Database.DB.Create(l)
	return l, nil
}

func (l *RegisterUser) ResetPassword() (*RegisterUser, error) {
	l.Password, _ = cmd.Http.Hash.Create(l.Password)
	cmd.Http.Database.DB.Updates(&l)
	return l, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := cmd.Http.Database.Where(&User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id interface{}) (*User, error) {
	var user User
	if err := cmd.Http.Database.Where("id = ? ", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx, userID uint, secret string) (*pkg.Token, error) {
	store := cmd.Http.Session.Get(c)
	store.Set("user_id", userID)
	c.Locals("user_id", userID)
	token, err := (&pkg.Token{}).CreateToken(c, userID, cmd.Http.Jwt.Secret, cmd.Http.Jwt.Expire)
	if err == nil {
		store.Set("user_token", token.Hash)
		store.Set("token_expiry", token.Expire)
	}
	store.Save()
	return token, err
}
