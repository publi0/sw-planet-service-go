package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sw-planet-service-go/dto"
	"sw-planet-service-go/model"
	"sw-planet-service-go/service"
	"sw-planet-service-go/util"
)

type API interface {
	CreatePlanet(c *gin.Context)
}

type Controller struct {
	Service service.Service
}

func (cl Controller) CreatePlanet(c *gin.Context) {
	var planet dto.CreatePlanet

	err := c.ShouldBindJSON(&planet)
	if err != nil {
		fields := util.ParseValidationFields(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fields})
		return
	}

	createPlanet, err := cl.Service.CreatePlanet(c, model.Planet{
		Name:    planet.Name,
		Terrain: planet.Terrain,
		Climate: planet.Climate,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, createPlanet)
}
