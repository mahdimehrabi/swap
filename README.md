# Project Setup and Documentation

## Setup

1. Copy the example environment variables file and configure your environment variables:
   ```sh
   cp env.example .env
   ```

2. Run the database migrations:
   ```sh
   make db-migrate-up
   ```

3. Seed the database:
   ```sh
   make db-seed
   ```

## Features

- **In-Memory Database**: Uses Redis for storing prices for fast access.
- **Cassandra Database**: Utilized for its scalability, suitable for non-complex relational data.
- **Arbitrary Precision Currency Types**: Created to ensure strong currency values and avoid floating point inaccuracies.
- **Worker Pool**: Implemented for fetching prices from the API in the background, including a rate limiter.
- **Error Handling**: Raises errors if a token price is not updated recently during a swap event.

## Considerations

- While ORMs can negatively impact performance, they are adequate for a test project.

## REST API Documentation

(Note: Swagger documentation was intended but could not be completed due to time constraints.)

### Swap (Get Price)

Create a transaction and get prices:
```sh
curl -X POST http://localhost:8080/swap/:userID/:srcCoinID/:destCoinID
```

### Commit Swap

Commit a transaction:
```sh
curl -X POST  curl -X POST http://localhost:8080/swap/commit/:transactionUUID
```

### User Management

#### Create User

```sh
curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password123"
    }'
```

#### Get User by ID

```sh
curl -X GET http://localhost:8080/users/1
```

#### Update User by ID

```sh
curl -X PUT http://localhost:8080/users/1 \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe Updated",
        "email": "john.updated@example.com",
        "password": "newpassword123"
    }'
```

#### Delete User by ID

```sh
curl -X DELETE http://localhost:8080/users/1
```

#### Get Users with Pagination

```sh
curl -X GET "http://localhost:8080/users?page=1&pageSize=10"
```