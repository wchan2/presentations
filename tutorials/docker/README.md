# Docker Workshop

## Goals

- Learn why Docker is important in the current distributed environments
- Learn how the essentials to work with Docker and common tools in its ecosystem
- Learn how to write a Dockerfile for a simple Go web service to create a Docker image and run a Docker container

## Why is Docker important?

Docker provides an abstraction for you environment and configurations.

## Setup

Note: This tutorial assumes that you are using the Mac OS with homebrew already installed. If you don't have homebrew, please follow [this guide](http://brew.sh/) to install it.

```bash
brew install docker
brew install docker-machine
brew install docker-compose
```

## Docker Machine

Docker machine is an environment that allows you to create virtual machines that run as Docker hosts. The `docker` command line is the client that connects to a Docker host to create images and start containers on them.

### Setting up the Docker machine

```bash
# 1. Create a Docker machine (VirtualBox virtual image) named, ts
docker-machine create --driver=virtualbox ts

# 2. Start the Docker machine
docker-machine start ts

# 3. Set up the Docker host environment variables
docker-machine env ts         # Check the environment variables to set up your `docker` command line client to connect to the docker host (docker-machine)
eval $(docker-machine env ts) # Automatically set up the necessary environment variables by evaluating the commands

# Optional
docker-machine --help     # lists all the commands available in docker-machine
docker-machine ls         # list all the docker machines that are created
docker-machine restart ts # restarts the 'ts' docker machine
docker-machine kill ts    # kills the 'ts' docker machine
docker-machine rm ts      # removes the 'ts' docker machine
```
