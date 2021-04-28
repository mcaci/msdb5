list:
	echo all, gen, briscola, briscola5, docker_clean, docker_run, staticcheck

all: gen briscola briscola5 docker_clean docker_run

gen: ./dom/briscola/briscola.proto
	protoc --proto_path=dom/briscola/ --go_out=. --go-grpc_out=. briscola.proto

briscola: ./dom/briscola/cmd/main.go ./dom/briscola/cmd/Dockerfile
	CGO_ENABLED=0 GOOS=linux go build -o briscola -a ./dom/briscola/cmd
	docker build -f ./dom/briscola/cmd/Dockerfile --force-rm -t briscola .
	rm briscola

briscola5: ./dom/briscola5/cmd/main.go ./dom/briscola5/cmd/Dockerfile
	CGO_ENABLED=0 GOOS=linux go build -o briscola5 -a ./dom/briscola5/cmd
	docker build -f ./dom/briscola5/cmd/Dockerfile -t briscola5 .
	rm briscola5

docker_clean:
	docker stop briscola briscola5
	docker rm briscola briscola5

docker_run:
	docker run -d -p 8081:8081 --name briscola briscola
	docker run -d -p 8082:8082 --name briscola5 briscola5

docker_restart:
	sudo systemctl start docker

staticcheck:
	staticcheck github.com/mcaci/msdb5/v2 github.com/mcaci/msdb5/v2/frw/... github.com/mcaci/msdb5/v2/app/... github.com/mcaci/msdb5/v2/dom/...