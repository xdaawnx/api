package main

import (
	config "api/conf"
	"api/libraries"
	appMidleware "api/middleware"
	"api/routes"
	util "api/utils"
	"log"
	"net/http"
	"os"
	"time"

	_ "api/docs" // swagger generated docs

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API
// @version 1.0
// @description API
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load conf
	config.LoadEnv()
	// Initialize Echo instance
	e := echo.New()
	e.HideBanner = true

	// db := config.ConnectDb()
	dbGrom := config.ConnectDbGorm()

	e.Validator = &libraries.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: appMidleware.LogSkipper,
		Handler: appMidleware.LogRequestResponse,
	}))

	e.HTTPErrorHandler = appMidleware.ErrorHandler

	// Register routes
	routes.NewTimeHandler(e, dbGrom)
	routes.NewPersonHandler(e, dbGrom)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/login", func(c echo.Context) error {
		type user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		u := &user{}
		if err := c.Bind(u); err != nil {
			return util.BadRequest
		}
		// Throws unauthorized error
		if u.Username != "joko" || u.Password != "parminto" {
			return util.Unauthorized
		}

		// Set custom claims
		claims := &appMidleware.CustomClaims{
			UserID: 123,
			Role:   "admin",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Set expiration time
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_Secret_Key")))
		log.Println(t)
		if err != nil {
			util.InternalServerError.Message = err.Error()
			return util.InternalServerError
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	userGroup := e.Group("/", echojwt.WithConfig(*appMidleware.ConfJWT()))
	// Retrieves details of the logged-in user
	// @Summary Get user details
	// @Description Retrieves details of the logged-in user
	// @Tags users
	// @Security BearerAuth
	// @Produce  json
	// @Success 200 {object} map[string]interface{}
	// @Router /users/test [get]
	userGroup.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "mantap",
		})
	})
	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
