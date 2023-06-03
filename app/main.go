package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"sw-planet-service-go/client"
	"sw-planet-service-go/config"
	"sw-planet-service-go/controller"
	"sw-planet-service-go/repository"
	"sw-planet-service-go/service"
)

const (
	ServiceName = "star-wars-service"
)

func main() {
	ctx := context.Background()
	ctx = zerolog.New(os.Stdout).With().
		Str("service", ServiceName).
		Logger().WithContext(ctx)

	clientHttp := http.Client{}
	dataBase := config.GetDataBase(ctx)
	rp := repository.Repository{Db: dataBase}
	c := client.Client{Client: &clientHttp}
	s := service.Service{Client: &c, Repository: &rp}
	cl := controller.Controller{Service: &s, Ctx: ctx}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(config.StructuredLogger(ctx))
	r.Use(config.HeaderValidator(ctx))
	r.Use(gin.Recovery())

	routerGroup := r.Group("/v1/starwars/planets")
	routerGroup.POST("/", cl.CreatePlanet)
	routerGroup.GET("/", cl.FindPlanets)
	routerGroup.GET("/:id", cl.FindPlanetById)
	routerGroup.DELETE("/:id", cl.DeletePlanetById)

	r.Run()
}
