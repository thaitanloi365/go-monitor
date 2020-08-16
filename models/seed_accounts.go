package models

import (
	"fmt"

	"github.com/thaitanloi365/go-monitor/models/enums"
)

// +84346876138
// +84378306893

func seedAccounts() {
	fmt.Println("[================ Seed Admin Account ================]")

	var users = []User{
		{
			Avatar:    "https://i.ibb.co/0B05sKH/sontung.jpg",
			FirstName: "Admin",
			LastName:  "",
			Name:      "Admin",
			Email:     "admin@gomonitor.com",
			Password:  "1234qwer",
			Role:      enums.AdminRole,
		},
		{
			Avatar:    "https://i.ibb.co/dkGX4m8/midu.jpg",
			FirstName: "Midu",
			LastName:  "Tester",
			Name:      "Midu Tester",
			Email:     "guess@gomonitor.com",
			Password:  "1234qwer",
			Role:      enums.GuessRole,
		},
	}

	for _, user := range users {
		var err = dbInstance.Debug().Unscoped().Where("email = ?", user.Email).FirstOrCreate(&user).Error
		if err != nil {
			continue
		}

		if user.Role == "admin" {
			continue
		}

	}

}
