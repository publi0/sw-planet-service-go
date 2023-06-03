package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sw-planet-service-go/model"
)

const (
	COLLECTION_NAME = "PLANETS"
)

type PlanetNotFoundError struct {
}

func (m *PlanetNotFoundError) Error() string {
	return "planet not found"
}

type API interface {
	SavePlanet(ctx context.Context, p model.Planet) error
	DeletePlanetById(ctx context.Context, id string) error
	FindPlanetById(ctx context.Context, id string) (model.Planet, error)
	FindPlanetsByName(ctx context.Context, name string) ([]model.Planet, error)
	FindAllPlanets(ctx context.Context) ([]model.Planet, error)
}

type Repository struct {
	Db *mongo.Database
}

func (r *Repository) SavePlanet(ctx context.Context, p model.Planet) error {
	result, err := r.Db.Collection(COLLECTION_NAME).InsertOne(ctx, p)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error while persisting planet")
		return err
	}
	log.Ctx(ctx).Info().Msgf("planet persisted with id [%s]", result.InsertedID)
	return nil
}

func (r *Repository) DeletePlanetById(ctx context.Context, id string) error {
	filter := bson.D{{"_id", id}}
	var planet model.Planet
	err := r.Db.Collection(COLLECTION_NAME).FindOneAndDelete(ctx, filter).Decode(&planet)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("planet not found")
		return &PlanetNotFoundError{}
	}
	log.Ctx(ctx).Info().Msg("deleted document")
	return nil
}

func (r *Repository) FindPlanetById(ctx context.Context, id string) (model.Planet, error) {
	filter := bson.D{{"_id", id}}
	var planet model.Planet
	err := r.Db.Collection(COLLECTION_NAME).FindOne(ctx, filter).Decode(&planet)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("planet not found")
		return planet, &PlanetNotFoundError{}
	}

	log.Ctx(ctx).Info().Msg("found planet")
	return planet, nil
}

func (r *Repository) FindPlanetsByName(ctx context.Context, name string) ([]model.Planet, error) {
	filter := bson.D{{"name", name}}
	var planets []model.Planet
	cursor, err := r.Db.Collection(COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error find planets by name")
		return nil, err
	}
	err = cursor.All(ctx, &planets)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error converting planets")
		return nil, err
	}
	log.Ctx(ctx).Info().Msgf("found [%s] planets", len(planets))
	return planets, nil
}

func (r *Repository) FindAllPlanets(ctx context.Context) ([]model.Planet, error) {
	var planets []model.Planet
	cursor, err := r.Db.Collection(COLLECTION_NAME).Find(ctx, bson.D{})
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error find planets")
		return nil, err
	}
	err = cursor.All(ctx, &planets)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error converting planets")
		return nil, err
	}
	log.Ctx(ctx).Info().Msgf("found [%n] planets", len(planets))
	return planets, nil
}
