# Star Wars - Planet CRUD

Star Wars Planet CRUD &amp; API Integration

## Build Status

## How to run

run `docker-compose up`

## Env Variables

| ENV       | DESCRIPTION                                                                                       |
|-----------|---------------------------------------------------------------------------------------------------|
| MONGO_URL | Mongo Connection String: mongodb+srv://username:password@mongodb.net/?retryWrites=true&w=majority |
| SW_HOST   | Star wars API: https://swapi.dev                                                                  |

## Documentation and examples

- OpenAPI 
  - Location: resources/OpenAPI.json

- Requests collection

  ![image](https://user-images.githubusercontent.com/14155185/119331527-d49d6f80-bc5d-11eb-93e3-00e2801fe19d.png)

## Features

- Add Planet (Name, Terrain, Climate)
    - Planets cannot have the same name
- Find Planet with film appearances information
    - Find by name
    - Find by id
    - Find all
- Delete a planet

## Stack

- Java 16
- Spring Boot
- Spring Data Mongo
- Spring Cacheable
- Mongo Database
- Feign Client
- Docker

## Integrations

- Star Wars API
    - https://swapi.dev/about