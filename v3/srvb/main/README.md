# srvb

make build

make run TAG=0.0.1

curl -XPOST  -H "Content-Type: application/json" localhost:4000/create -d '{"name":"newgame"}'
