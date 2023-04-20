package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/go-hclog"
	"github.com/shabacha/data"
	"github.com/shabacha/utils"
	"golang.org/x/crypto/bcrypt"
)

type Authentication interface {
	Authenticate(reqUser *data.User, user *data.User) bool
	GenerateAccessToken(user *data.User) (string, error)
	GenerateRefreshToken(user *data.User) (string, error)
	GenerateCustomKey(userID string, password string) string
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, string, error)
}
type AuthService struct {
	logger  hclog.Logger
	configs *utils.Configurations
}
type RefreshTokenCustomClaims struct {
	UserID    string
	CustomKey string
	KeyType   string
	jwt.StandardClaims
}
type AccessTokenCustomClaims struct {
	UserID  string
	KeyType string
	jwt.StandardClaims
}

func NewAuthService(logger hclog.Logger, configs *utils.Configurations) *AuthService {
	return &AuthService{logger, configs}
}
func (auth *AuthService) Authenticate(reqUser *data.User, user *data.User) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		auth.logger.Debug("Password Hashes Aren't The Same")
		return false
	}
	return true
}
