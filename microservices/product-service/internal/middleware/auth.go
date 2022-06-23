package middleware

import (
	"fmt"
	"os"
	"strings"

	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		// todo add to context

		// var id uint
		// ID := token.Claims.(jwt.MapClaims)["id"]
		// if ID != nil {
		// 	id = ID.(uint)
		// } else {
		// 	if err != nil {
		// 		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		// 	}
		// }

		return next(c)
	}
}
