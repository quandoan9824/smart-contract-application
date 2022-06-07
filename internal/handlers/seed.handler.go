package handlers

import (
	"math/rand"
	"net/http"

	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

const (
	SeedRecordCount = 500
)

type SeedHandler struct {
	dataSvc db.DataService
}

func NewSeedHandler(database db.MongoDBDatabase) *SeedHandler {
	ic := &SeedHandler{
		dataSvc: db.NewOrderDataService(database),
	}
	return ic
}

func (sHandler *SeedHandler) SeedDB(c *gin.Context) {
	for i := 0; i < SeedRecordCount; i++ {
		product := []models.Product{
			{
				Name:      faker.Name(),
				Price:     (uint)(rand.Intn(90) + 10),
				Remarks:   faker.Sentence(),
				UpdatedAt: faker.TimeString(),
			},
			{
				Name:      faker.Name(),
				Price:     (uint)(rand.Intn(1000) + 10),
				Remarks:   faker.Sentence(),
				UpdatedAt: faker.TimeString(),
			},
		}

		po := &models.Order{
			Products: product,
		}
		_, err := sHandler.dataSvc.Create(po)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable inserted data",
			})
			panic("Unable to insert data")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully inserted fake data",
		"Count":   SeedRecordCount,
	})
}
