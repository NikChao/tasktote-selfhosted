package models

type User struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	HouseholdIds []string `json:"householdIds"`
}
