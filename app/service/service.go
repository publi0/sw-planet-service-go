package service

import (
	"context"
	"sw-planet-service-go/client"
	"sw-planet-service-go/model"
)

type API interface {
	CreatePlanet(ctx context.Context, p model.Planet) (model.Planet, error)
}

type Service struct {
	Client client.Client
}

func (s Service) CreatePlanet(ctx context.Context, p model.Planet) (model.Planet, error) {
	appearances, err := s.Client.FindPlanetFilmAppearances(ctx, p.Name)
	p.FilmAppearances = appearances
	return p, err
}
