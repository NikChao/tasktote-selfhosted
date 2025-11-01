package routes

import (
	"api/models"
	"api/parsing"
	"api/providers"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/gin-gonic/gin"
)

func GroceryMagic(c *gin.Context) {
	var request models.GroceryMagicRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var groceryItems []models.GroceryItem
	layoutBlockMap := make(map[models.StorePreference][]models.LayoutBlock)

	var wg sync.WaitGroup

	for _, item := range request.GroceryList.Items {
		recipeUrl, isRecipeUrl := parseUrl(item.Name)

		if isRecipeUrl {
			wg.Add(1)
			providers.DeleteGroceryItem(item.HouseholdId, item.Id)
			go func() {
				defer wg.Done()
				recipeGroceryItems, extractedLayoutBlockMap := extractAndCreateGroceryItemsFromRecipeUrl(recipeUrl, request.HouseholdId, groceryItems, request.PreferredStores)

				groceryItems = append(groceryItems, recipeGroceryItems...)

				for key, value := range extractedLayoutBlockMap {
					layoutBlockMap[key] = append(layoutBlockMap[key], value...)
				}
			}()
			continue
		}

		storePreference := getStorePreferenceForItem(item)

		layoutBlockMap[storePreference] = append(layoutBlockMap[storePreference], models.LayoutBlock{
			Value: item.Id,
			Type:  models.GroceryItemId,
		})

		groceryItems = append(groceryItems, models.GroceryItem{
			Id:          item.Id,
			Name:        item.Name,
			Kind:        item.Kind,
			HouseholdId: item.HouseholdId,
			Checked:     item.Checked,
		})
	}

	wg.Wait()
	var layout []models.LayoutBlock
	for key, blocks := range layoutBlockMap {
		layout = append(layout, models.LayoutBlock{Value: string(key), Type: models.Text})
		layout = append(layout, blocks...)
	}

	groceryList := models.GroceryList{
		Items:  groceryItems,
		Layout: layout,
	}

	response := models.GroceryMagicResponse{
		GroceryList: groceryList,
	}

	c.JSON(http.StatusOK, response)
}

func getStorePreferenceForItem(item models.GroceryItem) models.StorePreference {
	return models.Unknown
}

func getStorePreferenceForItemName(itemName string) models.StorePreference {
	return models.Unknown
}

func removeEmojis(s string) string {
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, s)
}

func parseItemName(itemName string) string {
	parsedItemName := removeEmojis(itemName)
	parsedItemName = strings.TrimSpace(parsedItemName)
	return strings.ToLower(parsedItemName)
}

func parseUrl(itemName string) (string, bool) {
	u, err := url.ParseRequestURI(itemName)

	if err != nil {
		return "", false
	}

	return u.String(), true
}

func extractAndCreateGroceryItemsFromRecipeUrl(recipeUrl string, householdId string, existingGroceryItems []models.GroceryItem, preferredStores []models.StorePreference) ([]models.GroceryItem, map[models.StorePreference][]models.LayoutBlock) {
	recipe, _ := parsing.NewFromURL(recipeUrl)
	ingredients := recipe.IngredientList().Ingredients

	var groceryItems []models.GroceryItem = make([]models.GroceryItem, len(ingredients))
	layoutBlockMap := make(map[models.StorePreference][]models.LayoutBlock)

	for i, ingredient := range ingredients {
		if isIngredientAlreadyInGroceryList(ingredient, existingGroceryItems) {
			continue
		}

		storePreference := getCheapestStoreForItemOrStorePreference(ingredient.Name, preferredStores)

		groceryItem := models.GroceryItem{
			HouseholdId:   householdId,
			Name:          ingredient.Name,
			Checked:       false,
			StoreOverride: "",
		}

		groceryItem.GenerateID()

		providers.CreateGroceryItem(groceryItem)
		groceryItems[i] = groceryItem

		layoutBlockMap[storePreference] = append(layoutBlockMap[storePreference], models.LayoutBlock{
			Value: groceryItem.Id,
			Type:  models.GroceryItemId,
		})
	}

	return groceryItems, layoutBlockMap
}

func isIngredientAlreadyInGroceryList(ingredient parsing.Ingredient, groceryItems []models.GroceryItem) bool {
	for _, existingGroceryItem := range groceryItems {
		if ingredient.Name == existingGroceryItem.Name {
			return true
		}
	}

	return false
}

func getCheapestStoreForItemOrStorePreference(itemName string, preferredStores []models.StorePreference) models.StorePreference {
	return models.Unknown
}

func extractNumber(s string) (float64, error) {
	// Regular expression to match the first numeric value in the string
	re := regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]+`)
	match := re.FindString(s)
	if match == "" {
		return 0, fmt.Errorf("no numeric value found")
	}
	return strconv.ParseFloat(match, 64)
}
