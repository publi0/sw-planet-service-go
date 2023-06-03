package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"sw-planet-service-go/dto"
	"sw-planet-service-go/model"
	"sw-planet-service-go/repository"
	"sw-planet-service-go/service"
	"sw-planet-service-go/util"
)

type API interface {
	CreatePlanet(c *gin.Context)
	FindPlanets(c *gin.Context)
	FindPlanetById(c *gin.Context)
	DeletePlanetById(c *gin.Context)
}

type Controller struct {
	Service *service.Service
	Ctx     context.Context
}

func (cl Controller) CreatePlanet(c *gin.Context) {
	var planet dto.CreatePlanet
	ctx := log.Logger.With().
		Str("correlationId", c.GetHeader("correlation-id")).
		Logger().WithContext(cl.Ctx)

	log.Ctx(ctx).Info().Msg("request for create planet")

	err := c.ShouldBindJSON(&planet)
	if err != nil {
		fields := util.ParseValidationFields(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": fields})
		return
	}

	createPlanet, err := cl.Service.CreatePlanet(ctx, model.Planet{
		Name:    planet.Name,
		Terrain: planet.Terrain,
		Climate: planet.Climate,
	})

	if err != nil {
		var pd *service.PlanetWithDuplicatedNameError
		if errors.As(err, &pd) {
			c.JSON(http.StatusConflict, gin.H{"errors": "planet with this name already exists"})
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusCreated, createPlanet)
}

func (cl Controller) FindPlanets(c *gin.Context) {
	ctx := log.Logger.With().
		Str("correlationId", c.GetHeader("correlation-id")).
		Logger().WithContext(cl.Ctx)

	nameQuery := c.Query("name")

	var planets []model.Planet
	var err error

	if nameQuery == "" {
		log.Ctx(ctx).Info().Msg("request for find all planets")
		planets, err = cl.Service.FindPlanets(ctx)
	} else {
		log.Ctx(ctx).Info().Msgf("request for find all planets with name [%s]", nameQuery)
		planets, err = cl.Service.FindPlanetByName(ctx, nameQuery)
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, planets)
}

func (cl Controller) FindPlanetById(c *gin.Context) {
	ctx := log.Logger.With().
		Str("correlationId", c.GetHeader("correlation-id")).
		Logger().WithContext(cl.Ctx)

	id := c.Param("id")
	if id == "" {
		log.Ctx(ctx).Info().Msg("missing id")
		c.JSON(http.StatusBadRequest, gin.H{"errors": "missing id"})
	}

	log.Ctx(ctx).Info().Msgf("request for find planet with id [%s]", id)

	planet, err := cl.Service.FindPlanetById(ctx, id)
	if err != nil {
		var nf *repository.PlanetNotFoundError
		if errors.As(err, &nf) {
			c.JSON(http.StatusNotFound, gin.H{"errors": "planet not found"})
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, planet)
}

func (cl Controller) DeletePlanetById(c *gin.Context) {
	ctx := log.Logger.With().
		Str("correlationId", c.GetHeader("correlation-id")).
		Logger().WithContext(cl.Ctx)

	id := c.Param("id")
	if id == "" {
		log.Ctx(ctx).Info().Msg("missing id")
		c.JSON(http.StatusBadRequest, gin.H{"errors": "missing id"})
	}

	log.Ctx(ctx).Info().Msgf("request for delete planet with id [%s]", id)

	err := cl.Service.DeletePlanetById(ctx, id)
	if err != nil {
		var nf *repository.PlanetNotFoundError
		if errors.As(err, &nf) {
			c.JSON(http.StatusNotFound, gin.H{"errors": "planet not found"})
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
}
