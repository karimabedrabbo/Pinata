package tests

import (
	"github.com/brianvoe/gofakeit/v4"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func GenerateUniversities(count int) []*dbmodel.University {
	univs := make([]*dbmodel.University, count)
	for i := 0; i < count; i++ {
		univs[i] = GenerateUniversity()
	}
	return univs
}

func GenerateUniversity() *dbmodel.University {
	return &dbmodel.University{
		BaseModel:    GenerateBaseModel(),
		Name:         gofakeit.BeerName(),
		Alias:        gofakeit.BeerStyle(),
		Domain:       IncrementString() + ".edu",
		Lat:          float32(gofakeit.Latitude()),
		Lng:          float32(gofakeit.Longitude()),
		Size:         gofakeit.Number(1, 55000),
	}
}

