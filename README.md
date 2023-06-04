# Star Wars - Planet CRUD

Star Wars Planet CRUD &amp; API Integration

## Build Status

## How to run

run `docker-compose up`

run `make run`


## Env Variables

| ENV       | DESCRIPTION                                                                                       |
|-----------|---------------------------------------------------------------------------------------------------|
| MONGO_URL | Mongo Connection String: mongodb+srv://username:password@mongodb.net/?retryWrites=true&w=majority |
| SW_HOST   | Star wars API: https://swapi.dev                                                                  |

## Documentation and examples

- OpenAPI 
  - Location: resources/OpenAPI.json

## Features

- Add Planet (Name, Terrain, Climate)
    - Planets cannot have the same name
- Find Planet with film appearances information
    - Find by name
    - Find by id
    - Find all
- Delete a planet

## Stack

- Go 1.20
- Gin
- Mongo Database
- ZeroLog
- Docker

## Integrations

- Star Wars API
    - https://swapi.dev/about