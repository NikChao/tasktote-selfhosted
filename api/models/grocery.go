package models

import "github.com/google/uuid"

type BatchDeleteGroceryItemsRequest struct {
	ItemsToDelete []GroceryItem `json:"itemsToDelete"`
}

type GroceryItem struct {
	HouseholdId   string          `json:"householdId"`
	Id            string          `json:"id"`
	Name          string          `json:"name"`
	StoreOverride StorePreference `json:"storeOverride"`
	Category      string          `json:"category"`
	Checked       bool            `json:"checked"`
}

type LayoutBlockType string

const (
	Text          LayoutBlockType = "Text"
	GroceryItemId LayoutBlockType = "GroceryItemId"
)

type LayoutBlock struct {
	Value string          `json:"value"`
	Type  LayoutBlockType `json:"type"`
}

type GroceryList struct {
	Name   string        `json:"name"`
	Items  []GroceryItem `json:"items"`
	Layout []LayoutBlock `json:"layout"`
}

// Function to generate UUID for ID field
func (item *GroceryItem) GenerateID() {
	item.Id = uuid.NewString()
}

func (item *GroceryItem) GetOrGenerateID() string {
	if item.Id == "" {
		item.GenerateID()
	}

	return item.Id
}
