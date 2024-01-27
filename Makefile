migrateup:
	migrate -path db/migration -database "postgresql://default:secret@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://default:secret@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/pocikode/simple_bank_go/db/sqlc Store

.PHONY: migrateup migratedown sqlc mock