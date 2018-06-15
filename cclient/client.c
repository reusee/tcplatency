#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h> 
#include <time.h>

#define BUFSIZE 1024

//
// copied from
// https://www.cs.cmu.edu/afs/cs/academic/class/15213-f99/www/class26/tcpclient.c
//

/* 
 * error - wrapper for perror
 */
void error(char *msg) {
    perror(msg);
    exit(0);
}

int64_t timespecDiff(struct timespec *timeA_p, struct timespec *timeB_p)
{
    return ((timeA_p->tv_sec * 1000000000) + timeA_p->tv_nsec) -
           ((timeB_p->tv_sec * 1000000000) + timeB_p->tv_nsec);
}

int main(int argc, char **argv) {
    int sockfd, portno, n;
    struct sockaddr_in serveraddr;
    struct hostent *server;
    char *hostname;
    char buf[8];

    /* check command line arguments */
    if (argc != 2) {
       fprintf(stderr,"usage: %s <hostname>\n", argv[0]);
       exit(0);
    }
    hostname = argv[1];
    portno = 8081;

    /* socket: create the socket */
    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) 
        error("ERROR opening socket");

    /* gethostbyname: get the server's DNS entry */
    server = gethostbyname(hostname);
    if (server == NULL) {
        fprintf(stderr,"ERROR, no such host as %s\n", hostname);
        exit(0);
    }

    /* build the server's Internet address */
    bzero((char *) &serveraddr, sizeof(serveraddr));
    serveraddr.sin_family = AF_INET;
    bcopy((char *)server->h_addr, 
	  (char *)&serveraddr.sin_addr.s_addr, server->h_length);
    serveraddr.sin_port = htons(portno);

    /* connect: create a connection with the server */
    if (connect(sockfd, (const struct sockaddr*)&serveraddr, sizeof(serveraddr)) < 0) 
      error("ERROR connecting");
    
    bzero(buf, 8);
    int64_t total = 0;
    struct timespec start, end;
    for(int i = 0; i < 100; i++) {
      clock_gettime(CLOCK_MONOTONIC, &start);
      /* send the message line to the server */
      n = write(sockfd, buf, 8);
      if (n < 0) 
        error("ERROR writing to socket");
      n = read(sockfd, buf, 8);
      clock_gettime(CLOCK_MONOTONIC, &end);
      if (n < 0) 
        error("ERROR reading from socket");
      total += timespecDiff(&end, &start);
      usleep(100000);
    }
    printf("avg latency %ld nanoseconds (%ld microseconds)\n",
      total/100, total/100000);
    close(sockfd);
    return 0;
}
