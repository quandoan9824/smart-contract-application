package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceInfo struct {
	Name        string
	UpTime      time.Time
	Environment string
	Version     string
}

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"order_id"`
	LastUpdatedAt string             `bson:"last_updated_at,omitempty"`
	Products      []Product          `bson:"products,omitempty"`
}

type Product struct {
	Name      string `bson:"name,omitempty"`
	UpdatedAt string `bson:"updated_at,omitempty"`
	Price     uint   `bson:"price,omitempty"`
	Status    string `bson:"status,omitempty"`
	Remarks   string `bson:"remarks,omitempty"`
}
