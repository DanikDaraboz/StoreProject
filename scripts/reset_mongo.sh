#!/bin/bash

echo "Deleting data from MongoDB..."
docker exec -it mongo_local mongosh --eval "db.dropDatabase();"
echo "Database reset complete"
