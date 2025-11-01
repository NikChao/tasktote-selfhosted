package sqlite

import (
	"api/models"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	dbPath := filepath.Join("db", "groceries.db")

	// Open existing database
	sqliteDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Check if database can be accessed
	if err := sqliteDB.Ping(); err != nil {
		sqliteDB.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{sqliteDB}, nil
}

// Household Methods
func (db *DB) CreateHousehold(name string) (*models.Household, error) {
	uuidv7, _ := uuid.NewV7()
	id := uuidv7.String()

	_, err := db.Exec("INSERT INTO households (id, name) VALUES (?, ?)", id, name)
	if err != nil {
		return nil, fmt.Errorf("failed to create household: %w", err)
	}

	return &models.Household{Id: id, Name: name}, nil
}

func (db *DB) CreateUserHousehold(id string) (*models.Household, error) {
	_, err := db.Exec("INSERT INTO households (id, name) VALUES (?, ?)", id, id)
	if err != nil {
		return nil, fmt.Errorf("failed to create household: %w", err)
	}

	err = db.QueryRow("SELECT id FROM households WHERE rowid = last_insert_rowid()").Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to get household Id: %w", err)
	}

	return &models.Household{Id: id, Name: id}, nil
}

func (db *DB) GetHousehold(id string) (*models.Household, error) {
	var household models.Household
	err := db.QueryRow("SELECT id, name FROM households WHERE id = ?", id).Scan(&household.Id, &household.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("household not found")
		}
		return nil, fmt.Errorf("failed to get household: %w", err)
	}

	return &household, nil
}

func (db *DB) UpdateHousehold(id, name string) error {
	result, err := db.Exec("UPDATE households SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return fmt.Errorf("failed to update household: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("household not found")
	}

	return nil
}

func (db *DB) DeleteHousehold(id string) error {
	result, err := db.Exec("DELETE FROM households WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete household: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("household not found")
	}

	return nil
}

func (db *DB) ListHouseholds() ([]models.Household, error) {
	rows, err := db.Query("SELECT id, name FROM households ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("failed to list households: %w", err)
	}
	defer rows.Close()

	var households []models.Household
	for rows.Next() {
		var h models.Household
		if err := rows.Scan(&h.Id, &h.Name); err != nil {
			return nil, fmt.Errorf("failed to scan household row: %w", err)
		}
		households = append(households, h)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %w", err)
	}

	return households, nil
}

func (db *DB) CreateUser(name string) (*models.User, error) {
	_, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	var id string
	err = db.QueryRow("SELECT id FROM users WHERE rowid = last_insert_rowid()").Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user Id: %w", err)
	}

	return &models.User{Id: id, Name: name}, nil
}

func (db *DB) GetUser(id string) (*models.User, error) {
	var user models.User
	var name sql.NullString
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.Id, &name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if name.Valid {
		user.Name = name.String
	}

	return &user, nil
}

func (db *DB) UpdateUser(id string, name string) error {
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (db *DB) DeleteUser(id string) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// ListUsers returns all users
func (db *DB) ListUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT id, name FROM users ORDER BY name")
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		var name sql.NullString
		if err := rows.Scan(&u.Id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		if name.Valid {
			u.Name = name.String
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %w", err)
	}

	return users, nil
}

// Household-User Methods
func (db *DB) AddUserToHousehold(userId, householdId string) error {
	// First check if the user and household exist
	if _, err := db.GetUser(userId); err != nil {
		return err
	}
	if _, err := db.GetHousehold(householdId); err != nil {
		return err
	}

	_, err := db.Exec("INSERT INTO household_users (household_id, user_id) VALUES (?, ?)",
		householdId, userId)
	if err != nil {
		return fmt.Errorf("failed to add user to household: %w", err)
	}
	return nil
}

// RemoveUserFromHousehold dissociates a user from a household
func (db *DB) RemoveUserFromHousehold(userId, householdId string) error {
	result, err := db.Exec("DELETE FROM household_users WHERE household_id = ? AND user_id = ?",
		householdId, userId)
	if err != nil {
		return fmt.Errorf("failed to remove user from household: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("user not found in household")
	}

	return nil
}

// GetHouseholdUsers returns all users in a household
func (db *DB) GetHouseholdUsers(householdId string) ([]models.User, error) {
	// First check if the household exists
	if _, err := db.GetHousehold(householdId); err != nil {
		return nil, err
	}

	query := `
		SELECT u.id, u.name 
		FROM users u
		JOIN household_users hu ON u.id = hu.user_id
		WHERE hu.household_id = ?
		ORDER BY u.name
	`
	rows, err := db.Query(query, householdId)
	if err != nil {
		return nil, fmt.Errorf("failed to get household users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		var name sql.NullString
		if err := rows.Scan(&u.Id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		if name.Valid {
			u.Name = name.String
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %w", err)
	}

	return users, nil
}

// GetUserHouseholds returns all households a user belongs to
func (db *DB) GetUserHouseholds(userId string) ([]models.Household, error) {
	// First check if the user exists
	if _, err := db.GetUser(userId); err != nil {
		return nil, err
	}

	query := `
		SELECT h.id, h.name 
		FROM households h
		JOIN household_users hu ON h.id = hu.household_id
		WHERE hu.user_id = ?
		ORDER BY h.name
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user households: %w", err)
	}
	defer rows.Close()

	var households []models.Household
	for rows.Next() {
		var h models.Household
		if err := rows.Scan(&h.Id, &h.Name); err != nil {
			return nil, fmt.Errorf("failed to scan household row: %w", err)
		}
		households = append(households, h)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %w", err)
	}

	return households, nil
}

// Grocery Item Methods

// CreateGroceryItem adds a new grocery item
func (db *DB) CreateGroceryItem(name string, kind models.GroceryItemKind, category string, householdId string) (*models.GroceryItem, error) {
	// First check if the household exists
	if _, err := db.GetHousehold(householdId); err != nil {
		return nil, err
	}

	uuidv7, _ := uuid.NewV7()
	id := uuidv7.String()

	_, err := db.Exec("INSERT INTO grocery_items (id, name, kind, category, household_id) VALUES (?, ?, ?, ?, ?)",
		id, name, kind, category, householdId)
	if err != nil {
		return nil, fmt.Errorf("failed to create grocery item: %w", err)
	}

	return &models.GroceryItem{Id: id, Name: name, Kind: kind, Category: category, HouseholdId: householdId}, nil
}

// GetGroceryItem retrieves a grocery item by Id
func (db *DB) GetGroceryItem(id string) (*models.GroceryItem, error) {
	var item models.GroceryItem
	var category sql.NullString
	err := db.QueryRow("SELECT id, name, kind, category, household_id FROM grocery_items WHERE id = ?", id).
		Scan(&item.Id, &item.Name, &item.Kind, &category, &item.HouseholdId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("grocery item not found")
		}
		return nil, fmt.Errorf("failed to get grocery item: %w", err)
	}

	if category.Valid {
		item.Category = category.String
	}

	return &item, nil
}

// UpdateGroceryItem updates a grocery item
func (db *DB) UpdateGroceryItemStatus(id string, checked bool) error {
	result, err := db.Exec("UPDATE grocery_items SET checked = ? WHERE id = ?",
		checked, id)
	if err != nil {
		return fmt.Errorf("failed to update grocery item: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("grocery item not found")
	}

	return nil
}

func (db *DB) DeleteGroceryItems(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("no Ids provided")
	}

	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1] // Remove trailing comma

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	query := fmt.Sprintf("DELETE FROM grocery_items WHERE id IN (%s)", placeholders)
	result, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete grocery items: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("no grocery items were deleted")
	}

	return nil
}

func (db *DB) ListGroceryItemsByHousehold(householdId string) ([]models.GroceryItem, error) {
	if _, err := db.GetHousehold(householdId); err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, name, kind, category, household_id, checked FROM grocery_items WHERE household_id = ? ORDER BY name", householdId)
	if err != nil {
		return nil, fmt.Errorf("failed to list grocery items: %w", err)
	}
	defer rows.Close()

	items := make([]models.GroceryItem, 0)
	for rows.Next() {
		var i models.GroceryItem
		var category sql.NullString
		if err := rows.Scan(&i.Id, &i.Name, &i.Kind, &category, &i.HouseholdId, &i.Checked); err != nil {
			return nil, fmt.Errorf("failed to scan grocery item row: %w", err)
		}
		if category.Valid {
			i.Category = category.String
		}
		items = append(items, i)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %w", err)
	}

	return items, nil
}

func (db *DB) GetTaskSchedule(taskIds []string) ([]models.TaskScheduleItem, error) {
	placeholders := strings.Repeat("?,", len(taskIds))
	placeholders = placeholders[:len(placeholders)-1] // Remove trailing comma

	args := make([]interface{}, len(taskIds))
	for i, id := range taskIds {
		args[i] = id
	}

	query := fmt.Sprintf("SELECT task_id, date from scheduled_items WHERE task_id in (%s)", placeholders)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule for task: %w", err)
	}
	defer rows.Close()

	items := make([]models.TaskScheduleItem, 0)
	for rows.Next() {
		var item models.TaskScheduleItem
		if err := rows.Scan(&item.TaskId, &item.Date); err != nil {
			return nil, fmt.Errorf("failed to scan scheduled item: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in row iteration: %w", err)
	}

	return items, nil
}

func (db *DB) CreateTaskSchedule(taskId string, dates []string) error {
	_, err := db.Exec("DELETE FROM scheduled_items WHERE task_id = ?", taskId)
	if err != nil {
		return fmt.Errorf("failed to delete scheduled items: %w", err)
	}

	for _, date := range dates {
		_, err := db.Exec("INSERT INTO scheduled_items (task_id, date) VALUES (?, ?)", taskId, date)
		if err != nil {
			return fmt.Errorf("failed to scheduled item: %w", err)
		}

	}

	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}
