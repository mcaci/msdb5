# msdb5

To launch it first run `make all`. Then run `go run main.go` and `go test ./...` for the unit tests.

Reference for the Makefile creation at this [link](https://makefiletutorial.com/)

## testing with curl

Terminal 1:
curl -XPOST -F 'playername=ai1' -F 'type=create' -F 'gamename=default' localhost:8080/start/ai1

Terminal 2:
curl -XPOST -F 'playername=ai2' -F 'type=create' -F 'gamename=default' localhost:8080/start/ai2
