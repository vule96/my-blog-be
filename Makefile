DB_URL=postgresql://root:P@ssword123@localhost:5432/my_blog?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=P@ssword123 -d postgres:14.2-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root my_blog

dropdb:
	docker exec -it postgres dropdb my_blog

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server