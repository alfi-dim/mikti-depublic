package helper

import (
	"github.com/bxcodec/faker/v3"
	"mikti-depublic/model/domain"
)

func FakerUser() domain.User {
	return domain.User{
		UserID:      "USER-001",
		Name:        faker.Name(),
		Username:    faker.Username(),
		Email:       faker.Email(),
		Password:    faker.Password(),
		Phonenumber: faker.Phonenumber(),
	}
}

func FakerAdmin() domain.Admin {
	return domain.Admin{
		ID:       "ADMIN-001",
		Name:     faker.Name(),
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}

func FakerEvent() domain.Event {
	return domain.Event{
		ID:          "EVENT-001",
		AdminId:     "ADMIN-001",
		Name:        faker.Word(),
		Date:        faker.Date(),
		Address:     faker.Sentence(),
		Price:       200000,
		Tickets:     100,
		TicketsSold: 10,
	}

}

func FakerTransaction() domain.Transaction {
	return domain.Transaction{
		ID:         "TRANSACTION-001",
		UserID:     "USER-001",
		EventID:    "EVENT-001",
		Date:       faker.Date(),
		Quantity:   2,
		TotalPrice: 400000,
		Status:     "PENDING",
	}
}
