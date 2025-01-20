package main

import (
	"context"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/handlers"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/pkg/config"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.AppConfig.DBURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	userService := services.NewUserService(queries)

	authHandler := handlers.NewAuthHandler(userService)

	r := gin.Default()
	store := cookie.NewStore([]byte(config.AppConfig.SessionSecret))
	r.Use(sessions.Sessions("satpl-session", store))

	r.GET("/login", authHandler.Login)
	r.GET("/callback", authHandler.Callback)
	r.GET("/logout", authHandler.Logout)

	r.Run(":8080")
}
