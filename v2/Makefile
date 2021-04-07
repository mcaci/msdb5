list:
	echo all, gen, briscola, briscola5, docker_run

all: gen briscola briscola5 docker_run

gen: ./dom/briscola/briscola.proto
	protoc --proto_path=dom/briscola/ --go_out=. --go-grpc_out=. briscola.proto

briscola: ./dom/briscola/cmd/main.go ./dom/briscola/cmd/Dockerfile
	CGO_ENABLED=0 GOOS=linux go build -o briscola -a ./dom/briscola/cmd
	docker build -f ./dom/briscola/cmd/Dockerfile --force-rm -t briscola .
	docker rm briscola
	rm briscola

briscola5: ./dom/briscola5/cmd/main.go ./dom/briscola5/cmd/Dockerfile
	CGO_ENABLED=0 GOOS=linux go build -o briscola5 -a ./dom/briscola5/cmd
	docker build -f ./dom/briscola5/cmd/Dockerfile -t briscola5 .
	docker rm briscola5
	rm briscola5

docker_run:
	docker run -d -p 8081:8081 --name briscola briscola
	docker run -d -p 8080:8080 --name briscola5 briscola5
