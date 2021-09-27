.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o srvb
	docker build -t mcaci/msdb5/srvb:v${TAG} .
	rm srvb

.PHONY: run
run:
	docker run --rm --name srvb --detach --publish 4000:8080 mcaci/msdb5/srvb:v${TAG}