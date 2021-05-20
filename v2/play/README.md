# Play simulation

First run the server: from here `go run ../main.go`.
Then run this main.go twice, one on one terminal and the other on another terminal with the `--join` option can be done to do a play simulation of briscola in 2 players. This is useful to see a bit how the test goes.

## Testing with curl

Terminal 1:
curl -XPOST -F 'playername=ai1' -F 'type=create' -F 'gamename=default' localhost:8080/start/ai1

Terminal 2:
curl -XPOST -F 'playername=ai2' -F 'type=join' -F 'gamename=default' localhost:8080/start/ai2
