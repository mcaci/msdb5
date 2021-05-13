# test rabbitmq

## start 

```shell
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

# or 

sudo service rabbitmq-server start
```

## run

```shell
go run send.go
sudo rabbitmqctl list_queues
go run receive.go
```

## stop 

```shell
docker stop rabbitmq

# or 

sudo service rabbitmq-server start
```
