generate-mocks:
	mockgen -source=./core/repository/product.go -destination=./shared/mock/repository/repository_mock.go -package repository
