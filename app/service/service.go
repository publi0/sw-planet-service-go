package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"strings"
	"sw-planet-service-go/client"
	"sw-planet-service-go/model"
	"sw-planet-service-go/repository"
)

type PlanetWithDuplicatedNameError struct {
}

func (m *PlanetWithDuplicatedNameError) Error() string {
	return "planet with this name already exists"
}

type filmAppearancesChan struct {
	appearances int8
	err         error
}

type API interface {
	CreatePlanet(ctx context.Context, p model.Planet) (model.Planet, error)
	FindPlanets(ctx context.Context) ([]model.Planet, error)
	FindPlanetById(ctx context.Context, id string) (model.Planet, error)
	DeletePlanetById(ctx context.Context, id string) error
	FindPlanetByName(ctx context.Context, name string) ([]model.Planet, error)
}

type Service struct {
	Client     *client.Client
	Repository *repository.Repository
}

func (s Service) CreatePlanet(ctx context.Context, p model.Planet) (model.Planet, error) {
	pcChan := make(chan int)
	go func(pc chan int) {
		planets, _ := s.FindPlanetByName(ctx, p.Name)
		pc <- len(planets)
		close(pc)
	}(pcChan)

	faChan := make(chan filmAppearancesChan)
	go func(f chan filmAppearancesChan) {
		appearances, err := s.Client.FindPlanetFilmAppearances(ctx, p.Name)
		f <- filmAppearancesChan{
			appearances,
			err,
		}
		close(f)
	}(faChan)

	pcR := <-pcChan

	if pcR != 0 {
		log.Ctx(ctx).Info().Msg("planet with this name already exists")
		return model.Planet{}, &PlanetWithDuplicatedNameError{}
	}

	faR := <-faChan
	if faR.err != nil {
		log.Ctx(ctx).Error().Msg("error requesting starwars api")
		return model.Planet{}, faR.err
	}

	p.FilmAppearances = faR.appearances
	p.Id = uuid.New().String()
	p.Name = strings.ToLower(p.Name)
	err := s.Repository.SavePlanet(ctx, p)
	if err != nil {
		return model.Planet{}, err
	}
	return p, err
}

func (s Service) FindPlanets(ctx context.Context) ([]model.Planet, error) {
	planets, err := s.Repository.FindAllPlanets(ctx)
	if err != nil {
		return nil, err
	}

	if len(planets) == 0 {
		return []model.Planet{}, nil
	}
	return planets, nil
}

func (s Service) FindPlanetById(ctx context.Context, id string) (model.Planet, error) {
	planet, err := s.Repository.FindPlanetById(ctx, id)
	if err != nil {
		return model.Planet{}, err
	}
	return planet, nil
}

func (s Service) DeletePlanetById(ctx context.Context, id string) error {
	err := s.Repository.DeletePlanetById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) FindPlanetByName(ctx context.Context, name string) ([]model.Planet, error) {
	planets, err := s.Repository.FindPlanetsByName(ctx, strings.ToLower(name))
	if err != nil {
		return planets, err
	}
	if len(planets) == 0 {
		return []model.Planet{}, nil
	}
	return planets, nil
}
