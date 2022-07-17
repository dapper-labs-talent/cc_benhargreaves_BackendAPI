Docker is required to run this app

To build and run the application stack containing both the API and Postgres DB run the following command in the terminal from this directory:
> `docker-compose up --build -d`

The API listens on port `8080` 


Example Requests:

Get Users
``` http
GET /users HTTP/1.1
Host: localhost:8080
x-authentication-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28iLCJleHAiOjE2NTgwNTIxODF9.NC6GU3Ne1xpA12scLpJIrtTMkEQ9xkGOAT56DrUXOnc
```

Sign up
```http
POST /signup HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "email": "test@axiomzen.co",
  "password": "axiomzen",
  "firstName": "Alex",
  "lastName": "Zimmerman",
}
```

Login
```http
POST /login HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "email": "test@axiomzen.co",
  "password": "axiomzen"
}
```

Update name
```http
PUT /users HTTP/1.1
Host: localhost:8080
x-authentication-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28iLCJleHAiOjE2NTgwNDc0OTh9.nPJCiCWcHLc-mVesF0wuKb01UBhQpCjTa6sq0Hi1ZnI
Content-Type: application/json

{
  "lastName": "smith"
}
```