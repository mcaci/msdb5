FROM golang as build-env

LABEL author="Michele Caci <michele.caci@gmail.com>"

WORKDIR /go/src

COPY ./ /go/src/github.com/mcaci/msdb5

RUN cd /go/src/github.com/mcaci/msdb5; \
    go get github.com/gorilla/websocket golang.org/x/text; \
    CGO_ENABLED=0 go build

FROM scratch

WORKDIR /app/github.com/mcaci/msdb5/frw/templates
COPY --from=build-env /go/src/github.com/mcaci/msdb5/frw/templates/msdb5.html ./
WORKDIR /app
COPY --from=build-env /go/src/github.com/mcaci/msdb5/msdb5 ./

EXPOSE 8080

ENTRYPOINT ["./msdb5", "-lang", "en"]
