BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: install-air
install-air:
	go install github.com/cosmtrek/air@latest

build:
	go build -o ${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: dev
dev: install-air
	air

run: build
	./${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: clean
clean:
	rm -rf ${BUILD_FOLDER}

