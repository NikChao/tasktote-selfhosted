package providers

import (
	"api/models"
	db "api/proxy/sqlite"
)

func CreateHousehold() *models.Household {
	database, _ := db.NewDB()
	defer database.Close()

	household, _ := database.CreateHousehold("")
	return household
}

func JoinHousehold(userId string, householdId string) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.AddUserToHousehold(userId, householdId)
}

func LeaveHousehold(userId string, householdIdToRemove string) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.RemoveUserFromHousehold(userId, householdIdToRemove)
}

func GetOrCreateHousehold(id string) (*models.Household, error) {
	database, _ := db.NewDB()
	defer database.Close()

	household, err := database.GetHousehold(id)

	if err == nil && household != nil {
		return household, nil
	}

	return database.CreateUserHousehold(id)
}
