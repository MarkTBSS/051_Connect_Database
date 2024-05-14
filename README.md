```
docker run --name kawaii_db_test -e POSTGRES_USER=kawaii -e POSTGRES_PASSWORD=123456 -p 5433:5432 -d postgres:alpine

docker exec -it kawaii_db_test bash
psql -U kawaii
CREATE DATABASE kawaii_db_test;
\l

# Migrate up
migrate -database 'postgres://kawaii:123456@0.0.0.0:5433/kawaii_db_test?sslmode=disable' -source file://C:/Rayato159/E-Commerce/035_Migrate_Database/pkg/databases/migrations -verbose up

# Migrate down
migrate -database 'postgres://kawaii:123456@0.0.0.0:5433/kawaii_db_test?sslmode=disable' -source file://C:/Rayato159/E-Commerce/035_Migrate_Database/pkg/databases/migrations -verbose down

```