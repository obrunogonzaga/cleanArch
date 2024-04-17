createmigration:
	migrate create -ext=sql -dir=internal/infra/database/sql/migrations -seq init

migrateup:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

start:
	cd cmd/ordersystem/ && go run main.go wire.go
.PHONY: migrate