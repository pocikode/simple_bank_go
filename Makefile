DB_URL=postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb pocikode/simple-bank/db/sqlc Store

db_docs:
	dbdocs build ./doc/db.dbml

db_schema:
	dbml2sql ./doc/db.dbml -o ./doc/schema.sql --postgresql

.PHONY: migrateup migratedown migrateup1 migratedown1 test server mock db_docs db_schema