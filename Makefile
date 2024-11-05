TEMPLE_CMD := templ generate
TEST_CMD := go test ./...
TIDY_CMD := go mod tidy

# Define variables 
GOARCH := amd64
GOOS := linux  
BINARY_NAME := loadmonitor

# Generate target
generate:
	$(TEMPLE_CMD)

build: generate
	go build -o $(BINARY_NAME) ./cmd/$(BINARY_NAME).go

clean:
	rm -f $(BINARY_NAME)

run: build  # Specify build as a dependency here
	./$(BINARY_NAME)

tidy: 
	$(TIDY_CMD)

test: 
	$(TEST_CMD)
