package providers

import (
	"api/models"
	db "api/proxy/sqlite"
)

var usersTableName = "Users"

func CreateUser() *models.User {
	database, _ := db.NewDB()
	defer database.Close()

	user, _ := database.CreateUser("")
	user.HouseholdIds = []string{}
	return user
}

func UpdateUser(user models.User) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.UpdateUser(user.Id, user.Name)
}

func GetOrCreateUser(id string) (*models.User, error) {
	database, _ := db.NewDB()
	defer database.Close()

	user, err := database.GetUser(id)

	if user != nil && err == nil {
		households, _ := database.GetUserHouseholds(user.Id)
		var householdIds []string
		for _, household := range households {
			householdIds = append(householdIds, household.Id)
		}
		user.HouseholdIds = householdIds
		return user, nil
	}

	return CreateUser(), nil
}
