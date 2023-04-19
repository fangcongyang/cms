package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID          	int64
	Account     	string
	NickName    	string
	RoleId          string
	Permission    	string
	BufferTime  	int64
	jwt.StandardClaims
}
