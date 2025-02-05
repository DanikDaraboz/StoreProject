
#!/bin/bash

echo "Starting MongoDB (if not running)..."
docker ps | grep mongo_local > /dev/null
if [ $? -ne 0 ]; then
  docker run --name mongo_local -d -p 27017:27017 mongo
  echo "MongoDB started"
else
  echo "MongoDB is already running"
fi

echo "Seeding test data..."
docker exec -it mongo_local mongosh /docker-entrypoint-initdb.d/seed_mongo.js

echo "Running Go MongoDB test..."
go run test/mongo_local_con.go

echo "Test completed"
