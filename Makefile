all: test build

build: 
	echo "Building..."
	go build -o build/interval-merge main.go
	echo "Created Binary in folder: build"

run:
	go run main.go

test:
	echo "Testing the Application"
	go test example.com/interval-merge/merge