package providers

import (
	"api/models"
	db "api/proxy/sqlite"
	"fmt"
)

var groceriesTableName = "Groceries"

func GetGroceryItems(householdId string) ([]models.GroceryItem, error) {
	database, _ := db.NewDB()
	defer database.Close()

	_, err := GetOrCreateHousehold(householdId)
	if err != nil {
		return nil, fmt.Errorf("Could not get or create household %w: %w", householdId, err)
	}

	return database.ListGroceryItemsByHousehold(householdId)
}

func GetSchedule(taskIds []string) ([]models.TaskScheduleItem, error) {
	database, _ := db.NewDB()
	defer database.Close()

	return database.GetTaskSchedule(taskIds)
}

func CreateTaskSchedule(request models.ScheduleTaskRequest) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.CreateTaskSchedule(request.TaskId, request.Dates)
}

func CreateGroceryItem(groceryItem models.GroceryItem) error {
	database, _ := db.NewDB()
	defer database.Close()

	household, _ := database.GetHousehold(groceryItem.HouseholdId)
	if household == nil {
		database.CreateUserHousehold(groceryItem.HouseholdId)
	}

	_, err := database.CreateGroceryItem(groceryItem.Name, groceryItem.Kind, groceryItem.Category, groceryItem.HouseholdId)

	return err
}

func UpdateGroceryItem(groceryItem models.GroceryItem) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.UpdateGroceryItemStatus(groceryItem.Id, groceryItem.Checked)
}

func DeleteGroceryItem(householdId string, groceryItemId string) error {
	database, _ := db.NewDB()
	defer database.Close()

	return database.DeleteGroceryItems([]string{groceryItemId})
}

func BatchDeleteGroceryItems(groceryItems []models.GroceryItem) error {
	database, _ := db.NewDB()
	defer database.Close()

	ids := make([]string, len(groceryItems))
	for i, item := range groceryItems {
		ids[i] = item.Id
	}

	return database.DeleteGroceryItems(ids)
}
