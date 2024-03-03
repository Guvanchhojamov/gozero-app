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

# RPC
generate-authorization-proto:
	goctl rpc protoc  gateway/services/authorization/protos/v1/auth.proto --go_out=./gateway/services/authorization/rpc --go-grpc_out=./gateway/services/authorization/rpc/. --zrpc_out=./gateway/services/authorization/rpc/.
generate-products-proto:
	goctl rpc protoc  gateway/services/products/protos/v1/products.proto --go_out=./gateway/services/products/rpc --go-grpc_out=./gateway/services/products/rpc/. --zrpc_out=./gateway/services/products/rpc/.
generate-orders-proto:
	goctl rpc protoc  gateway/services/orders/protos/v1/orders.proto --go_out=./gateway/services/orders/rpc --go-grpc_out=./gateway/services/orders/rpc/. --zrpc_out=./gateway/services/orders/rpc/.

# Run
run-gateway:
	go run gateway/api/appserverapi.go -f ./gateway/api/etc/appServerApi-local.yaml
run-auth:
	go run ./gateway/services/authorization/rpc/auth.go -f ./gateway/services/authorization/rpc/etc/auth-local.yaml
run-product:
	go run ./gateway/services/products/rpc/products.go -f ./gateway/services/products/rpc/etc/products-local.yaml
run-order:
	go run ./gateway/services/orders/rpc/orders.go -f ./gateway/services/orders/rpc/etc/orders-local.yaml

# Docs
generate-docs-api:
	goctl api doc -dir . -o ./guide
