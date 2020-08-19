package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/models/enums"
	"golang.org/x/crypto/bcrypt"
)

// User user
type User struct {
	Model

	Avatar      string     `json:"avatar"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	TokenIssuer string     `gorm:"default:''" json:"-"`
	LastLogin   *int64     `gorm:"default:null" json:"last_login"`
	LoggedOutAt *int64     `gorm:"default:null" json:"logged_out_at"`
	Password    string     `gorm:"not null" json:"-"`
	Role        enums.Role `gorm:"default:'guess'" json:"-"`
	Timezone    string     `gorm:"default:'Asia/Ho_Chi_Minh'" json:"timezone"`
}

// LoginResponse response after user login successfully
type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// BeforeSave gorm hook
func (user *User) BeforeSave(scope *gorm.Scope) error {
	var name = user.GetFullName()
	if name != "" {
		scope.SetColumn("name", name)
	}
	return nil
}

// BeforeCreate gorm hook
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return user.HashPassword()
}

// HashPassword substitutes User.Password with its bcrypt hash
func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

// ComparePassword compares User.Password hash with raw password
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// GetFullName get full name
func (user *User) GetFullName() string {
	if user.FirstName == "" && user.LastName != "" {
		return user.LastName
	} else if user.LastName == "" && user.FirstName != "" {
		return user.FirstName
	} else if user.LastName != "" && user.FirstName != "" {
		return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	}

	return user.Name
}

// GenerateToken generate jwt
func (user *User) GenerateToken(duration ...time.Duration) (string, error) {
	claims := &JwtClaims{}
	claims.ID = user.ID
	claims.ExpiresAt = 0
	claims.Audience = string(user.Role)

	if user.TokenIssuer == "" {
		user.TokenIssuer = xid.New().String()
	}

	claims.Issuer = user.TokenIssuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.GetInstance().JWTSecret)
	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return t, nil
}
