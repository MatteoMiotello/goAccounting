package security

import (
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type UserJwtClaims struct {
	jwt.RegisteredClaims
	UserId   uint
	Email    string
	Name     string
	Username string
}

func BuildJwtFromUser(user models.User) (string, error) {
	expirationMin := viper.GetInt64("security.jwt.expiration")
	expiration := jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMin) * time.Minute))

	userClaims := UserJwtClaims{
		UserId:   user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiration,
			Issuer:    viper.GetString("security.jwt.issuer"),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedString, err := t.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ParseUserJwt(token string) (*UserJwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return t.Claims.(*UserJwtClaims), nil
}
