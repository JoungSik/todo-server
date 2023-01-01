package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;auto_increment"`
	Email     string    `gorm:"unique"`
	NickName  string    `gorm:"unique"`
	Password  string    `gorm:"unique"`
	Token     string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type JWTClaims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (user User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (u *User) GenerateToken() error {
	claims := &JWTClaims{
		u.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	u.Token = result
	return nil
}
