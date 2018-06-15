# tcplatency

This repo contains a demo that shows the extra latency introduced by the Go runtime or std lib when doing simple TCP communications. 

In this demo, we have a server process listening on port 8081, a client sends a 8bytes message to the server every 100 millisecond. On receiving such a message from the client, the server immediately sends the same message back so the client can measure the RTT time of this Ping-Pong style communication between client & server. We repeat such Ping-Pong style communication for 100 times and measure the average RTT. This repo contains both C & Go implementation for comparison.
 
The RTT observed by using the C client + C server is significantly lower than the Go client + Go server:

|Host 1|Host 2|avg rtt (microseconds)|
|:-------:|:-------:|:---------------------------:|
|go-server|go-client|193|
|go-server|c-client|165|
|c-server|go-client|166|
|c-server|c-client|140|


