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
│   │      ├── v1/       # Controller V1 Group
│   │         ├── dto/   # Data transfer objects (request & response)
│   │   ├── middleware/  # Custom middleware implementations
│   │   ├── routes/      # API route definitions
│   │      ├── v1/       # Routes V1 Group
│   ├── config/          # Configuration & dependency injection
│   ├── constants/       # Global constant variables
│   ├── contract/        # Dependency injection contracts
│   │   ├── common/      # Third-party dependencies
│   │   ├── repository/  # Repository layer contracts
│   │   └── service/     # Service layer contracts
│   ├── migrations/      # Database migration files
│   ├── model/           # Database models/entities
│   ├── pkg/             # Utility functions & helpers
│   │   ├── helper/      # Global helper function
│   │   ├── utils/       # Global library used to utility
│   ├── repository/      # Data access layer
│   │   ├── interfaces/  # Repository interface definitions
│   └── service/         # Business logic layer
│   │   ├── interfaces/  # Service interface definitions
│   │   ├── test/        # Service unit test
├── logs/                # Application log files
├── seeder/              # Seeder
├── pkg/                 # Shared third party libraries
└── .env                 # Environment variables
└── Makefile             # Command Shortcut
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

* `make run` (using Makefile)
* `go run cmd/api/main.go` (without Makefile)

### Using Docker

To build and run the project using Docker, use one of the following commands:

* `docker build -t go-boilerplate:1.0` (using Dockerfile)
* `docker compose up --build` (using Docker Compose)

### Testing

To run your testing, make sure you have generate mock folder and prepared file ```./shared/coverage/cover.out``` to store testing logs. you can use this following command:

`make mock` (using Makefile)

after generate mock files, use this following command to run all your testings:

`make test` (using Makefile)

it will be store the testing logs, to generate html version, use this following command:

`make coverage` (using Makefile)

> [!NOTE]
> In this project, we used `bou.monkey` to patch unmocked function to reach perfect coverage, it will be need some configuration on your computer, please used with your needed or contact me for help.

## Contact
----------

For more information or to report issues, please contact me at:

* [LinkedIn](https://www.linkedin.com/in/marifsulaksono/)
* [Email](mailto:marifsulaksono@gmail.com)