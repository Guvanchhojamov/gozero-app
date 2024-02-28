# Docker
app-dependencies-up:
	docker-compose -f ./deploy/app-dependencies/docker-compose.yaml up -d

# Database
create-migration:
	migrate create -ext sql -dir db/migrations $(name)
migrate:
	migrate -database 'postgres://db_admin:abc12345@localhost:5464/testdb?sslmode=disable' -path db/migrations up
migrate-rollback:
	migrate -database 'postgres://db_admin:abc12345@localhost:5464/testdb?sslmode=disable' -path db/migrations down

# Control
generate-api:
	goctl api go -api gateway/api/contract/gateway.api -dir ./gateway/api/ -style goZero
format-api:
	goctl api format -dir gateway/api/contract/gateway.api