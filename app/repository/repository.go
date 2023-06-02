package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sw-planet-service-go/model"
)

type API interface {
	savePlanet(ctx context.Context, p model.Planet) (model.Planet, error)
	deletePlanetById(ctx context.Context, id string) (model.Planet, error)
	findPlanetById(ctx context.Context, id string) (model.Planet, error)
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) savePlanet(ctx context.Context, p model.Planet) (model.Planet, error) {
	p.Id = uuid.New().String()
	r.db.Save(&p)
	return p, nil
}

func (r *Repository) deletePlanetById(ctx context.Context, id string) (model.Planet, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) findPlanetById(ctx context.Context, id string) (model.Planet, error) {
	//TODO implement me
	panic("implement me")
}
