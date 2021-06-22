package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/54m/echo-routing/output"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/star-integrations/project-boilerplate/back/docs"
	"github.com/star-integrations/project-boilerplate/back/pkg/config"
	"github.com/star-integrations/project-boilerplate/back/server"
	"github.com/star-integrations/project-boilerplate/back/server/props"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Title ProjectAPI
// @Version 0.0.1
// @Description Backend Server

// @Host localhost:1234
// @BasePath /

var (
	port = func() string {
		p := os.Getenv("PORT")
		if p == "" {
			p = "1234"
		}
		return p
	}()
	isMock = false
)

func main() {
	ctx := context.Background()
	e := echo.New()

	if !isMock {
		cfg, err := config.ReadConfig(ctx)
		if err != nil {
			log.Fatalf("failed to read config: %+v", err)
		}

		cp := &props.ControllerProps{
			Config: cfg,
		}
		ctrl := server.NewControllers(cp, e)

		// CORS settings
		ctrl.AddMiddleware("/api/", middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: cfg.CORSAllowOrigins,
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderXRequestedWith,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderCookie,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
				http.MethodPatch,
			},
			AllowCredentials: true,
		}))

		// TODO: turn on when authentication is required
		/*
			ctrl.AddMiddleware("/api/", echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
				ContextKey:     "jwt",
				SuccessHandler: nil,
				SigningKey:     []byte("Super_Secure_P4ssw0rd!"),
				SigningMethod:  jwt.SigningMethodHS512.Name,
				Claims:         new(jwt.StandardClaims),
				TokenLookup:    "cookie:ApiGenSession",
			}))
		*/

		e.GET("/swagger/*", echoSwagger.WrapHandler)
	} else {
		server.NewControllers(new(props.ControllerProps), e)
	}

	serverAddr := ":" + port

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	if !isMock {
		output.Do(e.Routes())

		fmt.Println("Swagger UI")
		fmt.Printf("http://localhost:%s/swagger/index.html", port)
	}

	e.Logger.Fatal(e.Start(serverAddr))
}
