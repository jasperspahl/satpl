package main

import (
	"context"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/handlers"
	"github.com/jasperspahl/satpl/internal/renderer"
	"github.com/jasperspahl/satpl/internal/services"
	"github.com/jasperspahl/satpl/pkg/config"
)

func main() {
	m, err := migrate.New("file://db/migrations", fmt.Sprintf("%s?sslmode=disable", config.AppConfig.DBURL))
	if err != nil {
		panic(err)
	}
	defer m.Close()
	m.Up()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.AppConfig.DBURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	userService := services.NewUserService(queries)
	playlistService := services.NewPlaylistService(queries)
	searchService := services.NewSearchService(queries)

	authHandler := handlers.NewAuthHandler(userService)
	homepageHandler := handlers.NewHomepageHandler(userService)
	playlistHandler := handlers.NewPlaylistHandler(playlistService)
	searchHandler := handlers.NewSearchHandler(searchService)

	r := gin.Default()
	store := cookie.NewStore([]byte(config.AppConfig.SessionSecret))
	r.HTMLRender = renderer.Default
	r.Use(sessions.Sessions("satpl-session", store))

	r.GET("/", homepageHandler.Home)
	r.GET("/config", homepageHandler.Config)

	r.GET("/playlists", playlistHandler.GetPlaylists)
	r.POST("/playlists", playlistHandler.Create)

	r.GET("/login", authHandler.Login)
	r.GET("/callback", authHandler.Callback)
	r.GET("/logout", authHandler.Logout)

	r.GET("/search", searchHandler.Search)

	r.Run(":8080")
}
