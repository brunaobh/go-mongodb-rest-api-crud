## A simple *golang app - REST API* CRUD connected to a NOSQL MongoDB Database and protected behind a NGINX proxy

## Introduction

Feel free to use, comment and contribute


## Instructions

Create a swarm ([Docker Swarm Cheatsheet](https://medium.com/@codingfriend/docker-swarm-cheatsheet-22665e3278b1){:target="_blank"}). You do this by running:

```bash
docker swarm init
```

To start the Docker services stack in a Docker Swarm cluster, I use the following command:

```bash
docker stack deploy -c docker-stack.yml myapp
```

## How to use

Check app health:

```bash
curl -k -H "Content-Type: application/json" http://<your ip:port>/health
```

Store new flight:

```bash
curl -k -d '{"flightNum":"LA8781", "airline":"TAM", "airport":"Confins", "status":"Confirmed", "expected":"20180713083500", "confirmed":"20180713083500"}' -H "Content-Type: application/json" -X POST http://<your ip:port>/flights
```

Get all flights:

```bash
curl -k -X GET http://<your ip:port>/flights
```
