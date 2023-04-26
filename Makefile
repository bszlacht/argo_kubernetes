build:
	go build -o matrix_main cmd/matrix/main.go

run:
	go run cmd/matrix/main.go

clean:
	go clean
	rm matrix_main

test:
	go test ./... -count=1

test_integration:
	go test ./integration_tests --tags=integration -count=1
