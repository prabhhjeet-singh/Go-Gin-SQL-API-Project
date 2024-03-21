A mock **Event Registration** project for practicing building API with Golang, Gin framework and SQL

## Following routes are implemented :- 
1. GET("/events")
2. GET("/events/:id")
3. POST("/events")
4. PUT("/events/:id")
5. DELETE("/events/:id")
6. POST("/signup")
7. POST("/login")

### In Process...
1. JWT Implementation
2. Authorization
3. Cancelling Registrations

## Example Usage

1. Get All Events <br> `GET http://localhost:8081/events`
2. Get Event by ID<br>`GET http://localhost:8081/events/id`
3. Create a New Event<br>`POST http://localhost:8081/events`
```http
POST http://localhost:8081/events 
Content-Type: application/json

{
    "name": "Test event",
    "description": "A test event",
    "location": "A test location",
    "dateTime": "2025-01-01T15:30:00.000Z"
}
```

4. Update Event<br>`PUT http://localhost:8081/events/1`
```http
PUT http://localhost:8081/events/1
Content-Type: application/json

{
    "name":"Updated test event",
    "description":"A test event",
    "location":"Test location (Updated)",
    "dateTime":"2025-01-01T15:30:00Z"
}
```

5. Delete Event<br>`DELETE http://localhost:8081/events/1`
   
6. Register User<br>`POST http://localhost:8081/signup`
```http
POST http://localhost:8081/signup
Content-Type: application/json

{
    "email":"test@test.com",
    "password":"password"
}
```
7. User Login<br>`POST http://localhost:8081/login`
```http
POST http://localhost:8081/login
Content-Type: application/json

{
    "email":"test@test.com",
    "password":"passwod"
}
```