# briscola

Lists of commands to build, run and clean docker image

- CGO_ENABLED=0 GOOS=linux go build -o briscola -a .
- docker build -t briscola .
- docker run -d -p 8081:8081 --name briscola briscola
- docker stop briscola
- docker run briscola

Flags `CGO_ENABLED=0` and `GOOS=linux` are necessary otherwise docker run doesn't work.

For testing http handlers read [this](https://blog.questionable.services/article/testing-http-handlers-go/).
