package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/fleimkeipa/eight-sup-api/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(user models.UserStruct) string {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "userToken"
	claims["value"] = user.Username
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Generate encoded token and send it as response.
	//declare your secret in .env
	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return err.Error()
	}
	return t
}

func CheckToken(c echo.Context) error {
	user := c.Get("userToken").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["value"].(string)
	return c.JSON(http.StatusOK, "Welcome "+name+"!")
}
