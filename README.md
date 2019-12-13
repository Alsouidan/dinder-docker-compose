# dinder-docker-compose

## To Run docker-compose
Simply clone the repo and in the root run `docker-compose up -d` followed by `docker-compose up`

## To run using docker build and run
### Redis
Run command `docker run --name redis -d redis`

### Server
Run command `docker build -t foo dinder && docker run -p 8081:8081 -it foo`

### Client
Run command `docker build -t foo dinder-frontend && docker run -p 3000:3000 -it foo`

## Services used
- Redis Cache Server
- MongoDB atlas

## Config file
A `sampleConfig.json` is provided in the server folder, replace the values by the correct ones to run correctly and rename the file to `config.json`

## Dependencies used
### Server Side:
- JWT authentication
- Gorilla MUX
- mongo-driver

### Client Side:
- axios
- jsonwebtoken
