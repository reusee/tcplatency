all: go-client go-server c-server c-client

c-server: cserver/tcp-echo-server.c
	gcc -O3 cserver/tcp-echo-server.c -o c-server

c-client: cclient/client.c
	gcc -O3 cclient/client.c -o c-client

go-client:
	go build -o go-client github.com/lni/tcplatency/goclient

go-server:
	go build -o go-server github.com/lni/tcplatency/goserver

clean:
	rm -f go-client go-server c-server c-client
