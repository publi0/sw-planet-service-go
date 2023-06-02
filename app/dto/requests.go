package dto

type CreatePlanet struct {
	Name    string `binding:"required"`
	Terrain string `binding:"required"`
	Climate string `binding:"required"`
}
