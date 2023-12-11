install:
	go mod download

build: install
	go build -o bin/workoutAnalyzer

run: build
	./bin/workoutAnalyzer

test:
	go test *.go -v