package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"sw-planet-service-go/client"
	"sw-planet-service-go/controller"
	"sw-planet-service-go/service"
)

func init() {
	godotenv.Load()
}

func main() {
	//db := config.DatabaseConnect()
	clientHttp := http.Client{}
	c := client.Client{Client: &clientHttp}
	s := service.Service{Client: c}
	cl := controller.Controller{Service: s}

	r := gin.Default()
	routerGroup := r.Group("/v1/starwars/planets")
	routerGroup.POST("/", cl.CreatePlanet)
	r.Run()
}
