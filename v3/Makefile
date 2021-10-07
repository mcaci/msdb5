list:
	echo test staticcheck

test:
	go test -cover -race ./... --timeout 5s

staticcheck:
	staticcheck github.com/mcaci/msdb5/v3 github.com/mcaci/msdb5/v3/briscola/... github.com/mcaci/msdb5/v3/srvb/... github.com/mcaci/msdb5/v3/srvp/...