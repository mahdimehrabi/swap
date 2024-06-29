## Setup 
```
cp env.example .env
```
config environment variables in .env 

run migrations
```
make db-migrate-up
```

run seeder
```
 make db-seed
```
### Extra features
### Used in memory db (redis) for storing prices
### Used cassandra db for best scalability because our entities do not have a complex relational
### Created arbitrary precision currency types for strong currency values and avoid floating point inaccuracies
### Implemented worker pool for getting price from API on the background along with rate limiter
### Seeder 
### Error for swap event if a token price is not updated recently

### Please consider
#### I know ORM has negative effect on performance, but I think for test project ORM is good enough.


## REST API Doc (I wanted to implement swagger too, but I changed my mind because I hadn't enough time on the weekend🥴 )
### Swap (get price)
create transaction and get prices
```
curl -X POST http://localhost:8080/swap/:userID/:srcCoinID/:destCoinID
``` 
### Commit swap 
```
curl -X POST http://localhost:8080/swap/:transactionUUID/commit
```


### User
Create User 
```
curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password123"
    }'
```

Get user with ID 

```
curl -X GET http://localhost:8080/users/1
```


Update user with id 
```
curl -X PUT http://localhost:8080/users/1 \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe Updated",
        "email": "john.updated@example.com",
        "password": "newpassword123"
    }'
```


Delete User with id
```
curl -X DELETE http://localhost:8080/users/1
```


Get users with pagination 
```
curl -X GET "http://localhost:8080/users?page=1&pageSize=10"
```