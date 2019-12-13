# dinder-docker-compose

## To Run docker-compose
Simply clone the repo and in the root run `docker-compose up -d` followed by `docker-compose up`

## To run using docker build and run
Run command `docker build` followed by `docker run <name of image>`

## Services used
-Redis Cache Server
-MongoDB atlas

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
