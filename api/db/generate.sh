#!/bin/bash

if [ ! -f "./init.sql" ]; then
  echo "Error: init.sql file not found in ./api/db"
  exit 1
fi

if [ -f "./groceries.db" ]; then
  echo "Removing existing groceries.db..."
  rm "./groceries.db"
fi

echo "Creating new groceries.db from init.sql..."
sqlite3 "./groceries.db" <"./init.sql"

if [ -f "./groceries.db" ]; then
  echo "Success: groceries.db has been created successfully!"
  echo "Tables in the new database:"
  sqlite3 "./groceries.db" ".tables"
else
  echo "Error: Failed to create groceries.db"
  exit 1
fi

exit 0
