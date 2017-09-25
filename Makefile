OUTPUT_DIR=./_output

build: clean
	go build -o ${OUTPUT_DIR}/tictac ./cmd

product: clean
	env GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/tictac.linux ./cmd

test: clean
	@go test -v -race ./pkg/...
	
.PHONY: clean
clean:
	@: if [ -f ${OUTPUT_DIR} ] then rm -rf ${OUTPUT_DIR} fi
