build:
	go build -o server noun-guesser/cmd/server
build_docker:
	docker build -t daniilty/wordle:latest -f docker/Dockerfile .
push_docker_amd64:
	docker buildx build --push --platform linux/amd64 -t daniilty/wordle:latest -f docker/Dockerfile .
