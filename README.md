

```markdown
# Pokemon API - Golang

An API developed in Go that retrieves Pokemon data and leverages a Redis cache for enhanced performance. The architecture is built upon the Hexagonal (or Ports and Adapters) pattern for adaptability and clear separation of concerns.

## Table of Contents
- [Overview](#overview)
- [Architecture](#architecture)
- [Project Layout](#project-layout)
- [Dependencies](#dependencies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Running the Application](#running-the-application)
- [Features](#features)
- [Future Work](#future-work)

---

## Overview
This API fetches Pokemon data. It uses a Redis cache to store the data, ensuring that subsequent requests are served faster. The Hexagonal design emphasizes a distinct separation between the core application logic and external interfaces, such as APIs or databases.

---

## Architecture
The Hexagonal Architecture segregates the application into:

- **Core Domain:** Central business logic and entities.
- **Ports:** Interfaces that define expected behaviors.
- **Adapters:** Implementations of the ports, allowing interaction with external systems.


---

## Project Layout
```
pokemon-api-golang-v2/
├── adapter/                  
│   ├── redis/                # Redis adapter
│   └── pokeapi/              # PokeAPI adapter
│
├── api/                      # API endpoints and handlers
│
├── cmd/                      # Main application
│
├── core/                     # Core logic
│   ├── entity/               # Entity definitions
│   └── usecase/              # Business logic
│
├── util/                     # Utilities
│   ├── logger/
│   └── errorhandler/
│
├── docker-compose.yml        
└── Dockerfile                
```

---

## Dependencies
- **Gin:** Web framework for building the API.
- **Go Redis:** Redis client for Go.
- **PokeAPI:** External service providing Pokemon data.

---

## Getting Started

### Prerequisites
Ensure you have Docker and Docker Compose installed:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running the Application
1. Clone the repository:
```bash
git clone <repository-url>
cd pokemon-api-golang-v2
```
2. Use Docker Compose to build and run:
```bash
docker-compose up --build
```
3. Access the API:
```
http://localhost:8080/pokemon/{pokemon_name}
```
Replace `{pokemon_name}` with the desired Pokemon's name.

4. Shutdown:
```bash
docker-compose down
```



---

## Features
- **Logging:** Provides insights into app events and errors.
- **Error Handling:** Ensures consistent error responses.

---

## Future Work
- Expand API functionalities, such as listing all Pokemon or filtering by type.
```

