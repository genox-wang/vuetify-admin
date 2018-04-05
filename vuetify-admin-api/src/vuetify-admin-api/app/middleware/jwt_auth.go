package middleware

import (
	"net/http"
	"vuetify-admin-api/app/model"

	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	//JWT_SECRET_KEY jwt secret key.
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

// JWTAuthRequired is auth middleware
func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		logrus.Info("[JWTAuthRequired] token : ", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET_KEY), nil
		})

		if err != nil {
			logrus.Error("[JWTAuthRequired] err: ", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := model.User{DisplayName: claims["name"].(string)}
			user.ID = uint(claims["id"].(float64))
			logrus.Info("[JWTAuthRequired] user : ", user)
			c.Set("user", user)
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// GetJWTToken build a jwt token for user
func GetJWTToken(user model.User) (token string) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.DisplayName,
		// "exp":  time.Now().Add(time.Minute * 5).Unix(),
		"exp": 0,
	})

	token, err := t.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		logrus.Error("[GetJWTToken] err: ", err.Error())
	}
	return token
}
