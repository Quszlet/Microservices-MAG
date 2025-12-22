migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/doctors_db?sslmode=disable" up

migrate -path ./doctors_service/migrations \  -database "postgres://postgres:postgres@localhost:5432/doctors_db?sslmode=disable" \
  force 1