package server

import (
	"backend/config"
	"backend/database"
	"backend/users/handlers"
	"backend/users/repositories"
	"backend/users/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	App *gin.Engine
	Db  *database.Database
	Cfg *config.Config
}

func NewGinServer(cfg *config.Config, db *database.Database) Server {
	return &ginServer{
		App: gin.Default(),
		Db:  db,
		Cfg: cfg,
	}
}

func (g *ginServer) Start() {
	g.ininitializedUserHttpHandlers()
	g.App.Run(fmt.Sprintf(":%d", g.Cfg.App.Port))
}

func (g *ginServer) ininitializedUserHttpHandlers() {
	database := *g.Db
	repository := repositories.NewPostgresRepository(database.GetDb())
	usecase := usecases.NewUsecaseImpl(repository)
	handler := handlers.NewHandlerImpl(usecase)

	userRoutes := g.App.Group("/users")
	userRoutes.POST("/", handler.InsertUser)
}
