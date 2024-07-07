migrateup:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb pocikode/simple-bank/db/sqlc Store

.PHONY: migrateup migratedown test server mock