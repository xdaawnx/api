package middleware

import (
	util "api/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func ConfJWT() *echojwt.Config {
	return &echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_Secret_Key")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		// TokenLookup: "header:Authorization:Bearer",
		// SuccessHandler: func(c echo.Context) {
		// 	user := c.Get("user").(*jwt.Token)
		// 	claims := user.Claims.(*CustomClaims)

		// 	// Custom validation for the "role" claim
		// 	if claims.Role != "admin" {
		// 		c.JSON(http.StatusForbidden, map[string]string{
		// 			"message": "Forbidden: You don't have the required role",
		// 		})
		// 		return
		// 	}
		// },
		ErrorHandler: func(c echo.Context, err error) error {
			// util.Unauthorized.Message = err.Error()
			util.Unauthorized.Message = "Invalid or expired token"
			return util.Unauthorized
		},
	}
}
