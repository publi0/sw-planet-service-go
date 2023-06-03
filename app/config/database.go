package config

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	DatabaseName = "STAR_WARS"
)

func GetDataBase(ctx context.Context) *mongo.Database {
	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("error while connecting to database")
	}
	return mdb.Database(DatabaseName)
}
