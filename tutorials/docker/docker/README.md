# Docker Workshop Part 1

## Sample Docker commands

```bash
docker ps -a     # see what docker containers are being run
docker images -a # see what docker images are already built
```

## Writing a simple Go web service

Create a file called `main.go` and copy and paste the following into it.

```go
package main

import (
	"log"
	"net/http"
	"fmt"
)



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello world\n")
	})
	log.Print("Running server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Creating a Dockerfile for the Go web service

Create a file called `Dockerfile` and copy and paste the following into it. For more Dockerfile instructions, refer [here](https://docs.docker.com/engine/reference/builder/).

```
# Grab image from Docker hub
FROM golang:1.7

MAINTAINER wchan@freewheel.tv

RUN mkdir example
COPY ./main.go example
WORKDIR example
RUN go build .

ENTRYPOINT ["./example"]
```

## Running the Go web service using Docker

```bash
docker build -t sample_program # build the docker image containing the Go web service
docker images                  # list all the images that are built in docker

docker run -d -p 8080:8080 --name dockertest sample_program # run the Go web service inside the Docker container
docker ps -a                                                # list all the docker containers

docker-machine inspect docker-workshop | grep IPAddress       # get the ip address of the Docker virtual machine
curl -XGET -v "http://{docker machine ip address here}:8080/" # make an http request to the web service
```
