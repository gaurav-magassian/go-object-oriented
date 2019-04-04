package business

import (
	"../models"
)

// BuildOrder ...
func BuildOrder() models.Customer {
	return models.Customer{
		FName:  "Visharad",
		LName:  "Dhavle",
		Age:    35,
		Gender: "Male",
		Orders: []models.Order{
			models.Order{
				ID: 1,
				Items: []models.Item{
					models.Item{
						Name:  "Nokia 6.2",
						Price: 6200.00,
					},
				},
			},
			models.Order{
				ID: 2,
				Items: []models.Item{
					models.Item{
						Name:  "Nokia 7.2",
						Price: 7200.00,
					},
				},
			},
		},
	}
}
