build:
	go build -o server noun-guesser/cmd/server
build_docker:
	docker build -t nount-guesser:latest -f docker/Dockerfile .
