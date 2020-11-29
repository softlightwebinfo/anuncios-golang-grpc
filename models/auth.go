package models

import (
	"cientosdeanuncios.com/backend/libs"
	"cientosdeanuncios.com/backend/settings"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type GraphModelAuthUsers struct {
	Id       int
	Email    string
	Password string
}
type GraphModelAuthUser struct {
	Users []GraphModelAuthUsers
}

type JwtToken struct {
	Token string `json:"token"`
}
type AuthUser struct {
	User  User   `json:"user"`
	Token string `json:"token"`
	jwt.StandardClaims
}

type AutCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthModel struct {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	ExpirationTime time.Time
	Claims         *AuthUser
	Token          *jwt.Token
}

func (auth *AuthModel) Expired() {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	auth.ExpirationTime = time.Now().Add(settings.LoginExpirationDuration)
}

func (auth *AuthModel) CreateToken(user User) (token string, err error) {
	auth.Expired()
	// Create the JWT claims, which includes the username and expiry time
	auth.Claims = &AuthUser{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: auth.ExpirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	auth.Token = jwt.NewWithClaims(settings.JwtSigningMethod, auth.Claims)
	// Create the JWT string
	token, err = auth.Token.SignedString(settings.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		err = errors.New(`Error en crear el token`)
		return
	}
	return
}

//func AuthDecodeUser(c *gin.Context) (user *AuthUser, isLogin bool) {
//	u, exist := c.Get("user")
//	if !exist {
//		isLogin = false
//		return
//	}
//	user = u.(*AuthUser)
//	isLogin = true
//	return
//}
func ComparePasswordAndGenerateToken(user GraphModelAuthUsers, password string) (success bool) {
	if !libs.ComparePasswords(user.Password, libs.GetPwd(password)) {
		return false
	}
	return true
}
