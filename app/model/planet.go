package model

type Planet struct {
	Name            string
	Terrain         string
	Climate         string
	FilmAppearances int8   `bson:"filmAppearances"`
	Id              string `bson:"_id"`
}
