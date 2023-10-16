
# Pokemon API - Golang v2

## Table of Contents

- [Overview](#overview)
- [Architecture Overview](#architecture-overview)
- [Key Dependencies](#key-dependencies)
- [Usage Guide](#usage-guide)
  - [Prerequisites](#prerequisites)
  - [Instructions](#instructions)
- [Project Structure](#project-structure)
- [Advanced Features](#advanced-features)
- [Future Enhancements](#future-enhancements)

## Overview

This project serves as an API written in Go that fetches Pokemon data...

## Architecture Overview

The application is structured based on the Hexagonal Architecture...

## Key Dependencies

- **Gin:** A lightweight web framework...

## Usage Guide

### Prerequisites:

Before diving in, ensure Docker and Docker Compose are installed...

### Instructions:

1. **Clone the Repository:**

   ```bash
   git clone <repository-url>
   ...

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

## Advanced Features

- **Logging:** Integrated logging utilities...

## Future Enhancements

- Expand the API to introduce features...

