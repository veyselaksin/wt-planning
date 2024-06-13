# Insider Case Study
## Introduction
This is a case study of a company called Insider. This case study goal is developing a system that automatically send message retrieved from the database, which have not yet been sent, every 2 minutes.

## Technologies
- Golang 1.22.3
- Fiber
- Gorm
- PostgreSQL
- Redis
- Docker
- Docker Compose

## How to run
1. Clone this repository
2. Create a `.env` file in the root directory and fill it with the following environment variables:
```
POSTGRES_DB_HOST=db
POSTGRES_DB_USERNAME=postgres
POSTGRES_DB_PASSWORD=postgres
POSTGRES_DB_NAME=postgres
POSTGRES_DB_PORT=5432
POSTGRES_DB_APP_NAME=case_study

APP_PORT=8000
APP_HOST=localhost
APP_DEBUG=true

WEBHOOK_URL=https://webhook.site/<your-webhook-id>

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_DB=0
```
3. Run the following command to start the application:
```
docker-compose up --build
```

## Run Test
To run the test, you can run the following command:
```
go test ./... -v
```

## API Documentation
You can find the API documentation on the following link when you run the application:
```
http://localhost:8000/v1/docs
```
