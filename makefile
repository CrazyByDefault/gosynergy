all:
	go build -o ./hostd/host.o ./hostd/
	go build -o ./clientd/client.o ./clientd/