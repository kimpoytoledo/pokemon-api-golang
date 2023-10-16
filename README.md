Certainly! Here's the complete and organized `README.md` for your project:

---

# Pokemon API - Golang

## Table of Contents

- [Overview](#overview)
- [Architecture Overview](#architecture-overview)
- [Project Structure](#project-structure)
- [Key Dependencies](#key-dependencies)
- [Usage Guide](#usage-guide)
  - [Prerequisites](#prerequisites)
  - [Instructions](#instructions)
- [Advanced Features](#advanced-features)
- [Future Enhancements](#future-enhancements)

## Overview

This project serves as an API written in Go that fetches Pokemon data. It uses a Redis cache to store and retrieve Pokemon data swiftly for subsequent requests, ensuring enhanced performance. The Hexagonal (or Ports and Adapters) design pattern underpins the architecture, emphasizing a clear separation of concerns and adaptability.

## Architecture Overview

The application adopts the Hexagonal Architecture, which segregates it into distinct sections:

- **Core Domain:** Contains essential business logic and entities.
- **Ports:** Interfaces that establish expected interactions and behaviors.
- **Adapters:** Implementations of the ports that facilitate interaction with external systems or services.

## Project Structure

```
pokemon-api-golang-v2/
│
├── adapter/                  # Adapters for external services
│   ├── redis/                # Redis adapter
│   └── pokeapi/              # PokeAPI adapter
│
├── api/                      # API-related code (endpoints, handlers)
│
├── cmd/                      # Main application entry point
│
├── core/                     # Core domain logic
│   ├── entity/               # Entity definitions
│   └── usecase/              # Core business logic
│
├── util/                     # Utilities (logging, error handling)
│   ├── logger/
│   └── errorhandler/
│
├── docker-compose.yml        # Docker Compose configuration
└── Dockerfile                # Dockerfile for the Go application
```

## Key Dependencies

- **Gin:** A lightweight web framework utilized for crafting the API.
- **Go Redis:** A Redis client tailored for Go, facilitating interaction with the Redis cache.
- **PokeAPI:** An external service supplying comprehensive Pokemon data.

## Usage Guide

### Prerequisites:

Before diving in, ensure Docker and Docker Compose are installed:

- [Docker installation guide](https://docs.docker.com/get-docker/)
- [Docker Compose installation guide](https://docs.docker.com/compose/install/)

### Instructions:

1. **Clone the Repository:**

   ```bash
   git clone <repository-url>
   cd pokemon-api-golang-v2
   ```

2. **Activate the Services:**

   Initiate the build and startup process using Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. **Engage with the API:**

   Once active, the API becomes accessible at:

   ```
   http://localhost:8080/pokemon/{pokemon_name}
   ```

   Replace `{pokemon_name}` with the name of the desired Pokemon (e.g., `bulbasaur`). The response will deliver the Pokemon data and specify the data source—either "Redis Cache" or "PokeAPI".

4. **Deactivate the Services:**

   To gracefully halt the services, use `CTRL+C` in the terminal. For a complete shutdown and removal of containers, execute:

   ```bash
   docker-compose down
   ```

## Advanced Features

- **Logging:** The integrated logging utilities provide insights into significant events and potential errors, aiding in troubleshooting and understanding the application's behavior.
- **Error Handling:** A dedicated utility for error handling ensures consistency and extensibility in managing and reporting errors.

## Future Enhancements

- Consider expanding the API to introduce features beyond just fetching Pokemon data, like listing all Pokemon, searching by type, and so on.

---
