migrateup:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb pocikode/simple-bank/db/sqlc Store

.PHONY: migrateup migratedown migrateup1 migratedown1 test server mock