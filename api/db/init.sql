-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- Create the households table
CREATE TABLE households (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- Create the users table
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    name TEXT
);

-- Create the household_users join table
CREATE TABLE household_users (
    household_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    PRIMARY KEY (household_id, user_id),
    FOREIGN KEY (household_id) REFERENCES households(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the grocery_items table
CREATE TABLE grocery_items (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    kind TEXT NOT NULL,
    category TEXT,
    checked BOOLEAN DEFAULT FALSE,
    household_id TEXT NOT NULL,
    FOREIGN KEY (household_id) REFERENCES households(id) ON DELETE CASCADE
);

-- Create an insert trigger for households to generate UUID
CREATE TRIGGER insert_household_id
AFTER INSERT ON households
WHEN new.id IS NULL
BEGIN
    UPDATE households SET id = (lower(hex(randomblob(16)))) WHERE rowid = new.rowid;
END;

-- Create an insert trigger for users to generate UUID
CREATE TRIGGER insert_user_id
AFTER INSERT ON users
WHEN new.id IS NULL
BEGIN
    UPDATE users SET id = (lower(hex(randomblob(16)))) WHERE rowid = new.rowid;
END;

-- Create an insert trigger for grocery_items to generate UUID
CREATE TRIGGER insert_grocery_item_id
AFTER INSERT ON grocery_items
WHEN new.id IS NULL
BEGIN
    UPDATE grocery_items SET id = (lower(hex(randomblob(16)))) WHERE rowid = new.rowid;
END;

-- Create indexes for better performance
CREATE INDEX idx_household_users_household_id ON household_users(household_id);
CREATE INDEX idx_household_users_user_id ON household_users(user_id);
CREATE INDEX idx_grocery_items_household_id ON grocery_items(household_id);
