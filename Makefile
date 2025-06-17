run: 
	go run cmd/api/main.go

seed:
	go run cmd/api/main.go -seed

mock:
	@mkdir -p shared/mock
	@rm -rf shared/mock/*

	@echo "mock contract"
	@mockgen --source=./internal/contract/repository/repository.go -destination=./shared/mock/contract/repository/repository.go --package=mock_contract
	@mockgen --source=./internal/contract/service/service.go -destination=./shared/mock/contract/service/service.go --package=mock_contract

	@echo "mock repositories"
	@mockgen --source=./internal/repository/interfaces/user.go -destination=./shared/mock/repository/user.go --package=mock_repository

test:
	@go test -v -gcflags=all=-l -cover -coverpkg=./internal/... -coverprofile=./shared/coverage/cover.out ./...

coverage:
	@go tool cover -func=./shared/coverage/cover.out
	@go tool cover -html=./shared/coverage/cover.out  -o ./shared/coverage/cover.html