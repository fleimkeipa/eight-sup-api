package db

import (
	"testing"

	"github.com/fleimkeipa/eight-sup-api/models"
	"github.com/fleimkeipa/eight-sup-api/pkg/utils"
)

func BenchmarkCreateEvent(b *testing.B) {
	client := utils.Connect()
	defer utils.Close(client)
	user := client.Database("eight-sup2").Collection("user")
	event := client.Database("eight-sup2").Collection("event")
	temp := models.Event{
		BuyerUsername:  "user2",
		SellerUsername: "user1",
		Unique:         "gold",
		Items: []models.Items{
			{
				BuyerUsername: "user2",
				Status:        "requested",
				Prop:          "Testing",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		CreateEvent(&temp, event, user)
	}
}
