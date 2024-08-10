DB_URL=postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate	create -ext sql -dir db/migration -seq $(name)

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb pocikode/simple-bank/db/sqlc Store

db_docs:
	dbdocs build ./doc/db.dbml

db_schema:
	dbml2sql ./doc/db.dbml -o ./doc/schema.sql --postgresql

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative  \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
		proto/*.proto

	statik -f -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 8014 -r repl

.PHONY: migrateup migratedown migrateup1 migratedown1 new_migration test server mock db_docs db_schema proto evans