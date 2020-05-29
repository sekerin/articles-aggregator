BUILD_DIR=./build

search-build:
	go build -o $(BUILD_DIR) github.com/sekerin/artticles-aggregator/internal/services/searchd

parser-build:
	go build -o $(BUILD_DIR) github.com/sekerin/artticles-aggregator/internal/services/parserd

search-run:
	go run github.com/sekerin/artticles-aggregator/internal/services/searchd

parser-run:
	go run github.com/sekerin/artticles-aggregator/internal/services/parserd
test:
	go test ./...