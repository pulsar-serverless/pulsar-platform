package main

import (
	"log"
	"os"

	"pulsar/internal/adapters/secondary/postgres"
	"pulsar/internal/core/domain/billing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConn := postgres.SetupDB(os.Getenv("POSTGRES_CONNECTION"))

	plans := []*billing.PricingPlan{
		{
			ID:              uuid.New(),
			Name:            "Free plan",
			Description:     "Ideal for experimenting, learning, and building small-scale serverless applications.",
			PlanTeir:        billing.Free,
			NotifyThreshold: 0,
			Price:           0.00,
			PlanResources: billing.PlanResources{
				ID:        uuid.New(),
				Memory:    100,
				Bandwidth: 100,
				Requests:  100,
			},
		},
		{
			ID:              uuid.New(),
			Name:            "Pro plan",
			Description:     "Paid plans offer more resources, allowing you to handle higher traffic and more complex applications.",
			PlanTeir:        billing.Pro,
			NotifyThreshold: 0,
			Price:           9.99,
			PlanResources: billing.PlanResources{
				ID:        uuid.New(),
				Memory:    100,
				Bandwidth: 100,
				Requests:  100,
			},
		},
	}
	result := dbConn.Create(plans)
	if result.Error != nil {
		log.Fatal("Unable to seed pricing plan")
	}
}
