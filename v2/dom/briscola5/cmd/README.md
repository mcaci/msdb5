# briscola5

Lists of commands to build, run and clean docker image

- CGO_ENABLED=0 GOOS=linux go build -o briscola5 -a .
- docker build -t briscola5 .
- docker run -d -p 8080:8080 --name briscola5 briscola5
- docker stop briscola5
- docker run briscola5

Flags `CGO_ENABLED=0` and `GOOS=linux` are necessary otherwise docker run doesn't work.

For testing http handlers read [this](https://blog.questionable.services/article/testing-http-handlers-go/).
