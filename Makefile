OUTPUT_DIR = bin/app
MAIN = cmd/main.go
.PHONY: all build clean run tidy

all: build

build:
	@echo "Building the application..."
	go build -o ${OUTPUT_DIR} $(MAIN)

run: build
	@echo "Running the application..."
	${OUTPUT_DIR}

clean:
	@echo "Cleaning up..."
	rm -rf ${OUTPUT_DIR}

tidy:
	@echo "Tidying up dependencies..."
	go mod tidy