package settings

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	LoginExpirationDuration = time.Duration(24) * time.Hour
)

var JwtKey = []byte("softlightweb_code_secret_@12a3dfscx#1789")
var JwtSigningMethod = jwt.SigningMethodHS256