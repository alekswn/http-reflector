BIN_DIR ?= "BUILD"
EXEC_NAME ?= "http-reflector"

build-static:
	@(echo "-> Creating statically linked binary...")
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -a -o $(BIN_DIR)/$(EXEC_NAME)

run:
	go run main.go
