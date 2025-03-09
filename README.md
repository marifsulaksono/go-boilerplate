# Go Boilerplate

A clean architecture scalable and maintainable Go boilerplate project using Echo, JWT, GORM, Redis, and more.

## Table of Contents

* [Overview](#overview)
* [Tech Stack](#tech-stack)
* [Getting Started](#getting-started)
* [Installation](#installation)
* [Running the Project](#running-the-project)
* [Contact](#contact)

## Overview

This project is a boilerplate template that provides a basic structure with a clean architecture implementation for building scalable and maintainable Go applications. It includes features such as authentication using JWT, database interaction using GORM, and temporary storage using Redis.

You can use this repository for your template project by click [use this template](https://github.com/new?template_name=go-boilerplate&template_owner=marifsulaksono)

## Tech Stack

* Go 1.23 (See [installation](https://go.dev/doc/install))
* Echo V4 (See [documentation](https://echo.labstack.com/docs))
* JWT V5 (See [documentation]([https://echo.labstack.com/docs](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)))
* Postgres/MySQL/SQLServer (See installation of [MySQL](https://dev.mysql.com/doc/mysql-getting-started/en/) | [Postgresql](https://www.postgresql.org/docs/current/tutorial-install.html) | [SQL Server](https://learn.microsoft.com/en-us/sql/database-engine/install-windows/install-sql-server?view=sql-server-ver16))
* GORM (See [documentation](https://gorm.io/docs/))
* Redis (See [documentation](https://redis.io/docs/latest/develop/))
* Viper (See [documentation](https://pkg.go.dev/github.com/dvln/viper))
* Logrus (See [documentation](https://pkg.go.dev/github.com/sirupsen/logrus))

## Folder Structure
```
go-echo-boilerplate/  
├── cmd/                 # Entry point for the application  
│   └── api/             # REST API starter  
├── internal/            # Internal application logic  
│   ├── api/             # REST API core logic  
│   │   ├── controller/  # Handles request & response processing  
│   │   ├── dto/         # Data transfer objects (request & response)  
│   │   ├── middleware/  # Custom middleware implementations  
│   │   ├── routes/      # API route definitions  
│   ├── config/          # Configuration & dependency injection  
│   ├── constants/       # Global constant variables  
│   ├── contract/        # Dependency injection contracts  
│   │   ├── common/      # Third-party dependencies  
│   │   ├── repository/  # Repository layer contracts  
│   │   └── service/     # Service layer contracts  
│   ├── migrations/      # Database migration files  
│   ├── model/           # Database models/entities  
│   ├── pkg/             # Utility functions & helpers  
│   ├── repository/      # Data access layer  
│   │   ├── interfaces/  # Repository interface definitions  
│   └── service/         # Business logic layer  
│   │   ├── interfaces/  # Service interface definitions  
├── logs/                # Application log files  
├── pkg/                 # Shared utilities  
└── .env                 # Environment variables
```

## Getting Started

### Installation

To install this project, clone the repository from GitHub:

* `git clone https://github.com/marifsulaksono/go-boilerplate.git`
* Copy file `.env.example` and rename to `.env`
  ```sh
  cp .env.example .env
  ```
* Adjust variable `.env` file according to the configuration in your local environment

### Running the Project

To run the project, use one of the following commands:

* `make run-api` (using Makefile)
* `go run cmd/api/main.go` (without Makefile)

### Using Docker

To build and run the project using Docker, use one of the following commands:

* `docker build -t go-boilerplate:1.0` (using Dockerfile)
* `docker compose up --build` (using Docker Compose)

## Contact
----------

For more information or to report issues, please contact me at:

* [LinkedIn](https://www.linkedin.com/in/marifsulaksono/)
* [Email](mailto:marifsulaksono@gmail.com)