OUTPUT_DIR = bin/app
API_MAIN = cmd/api/main.go
.PHONY: all build clean run tidy

all: build

build:
	@echo "Building the application..."
	go build -o ${OUTPUT_DIR} $(API_MAIN)

run: build
	@echo "Running the application..."
	${OUTPUT_DIR}

clean:
	@echo "Cleaning up..."
	rm -rf ${OUTPUT_DIR}

tidy:
	@echo "Tidying up dependencies..."
	go mod tidy