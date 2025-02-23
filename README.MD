# Auth API Service

A REST API service for user authentication and authorization using JWT tokens.

## Features

- User signup with email and password
- User signin with JWT token generation
- Token authorization and validation
- Token refresh mechanism
- Token expiry handling
- Secure password hashing
- MySQL database storage

## Prerequisites

## API Endpoints

- POST /api/auth/signup
    
- POST /api/auth/signin
- POST /api/auth/refresh
- GET /api/auth/check

## Configuration

Key configuration settings in `config/config.go`:


## Curl commands
Signup
```bash
curl --location 'http://localhost:8080/api/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@example.com",
    "password": "password123"
}'
```

Signin
```bash
curl --location 'http://localhost:8080/api/auth/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@example.com",
    "password": "password123"
}'
```

Refresh
```bash
curl --location --request POST 'http://localhost:8080/api/auth/refresh' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlR5cGUiOiJyZWZyZXNoIiwiZXhwIjoxNzQwOTE4ODUxLCJpYXQiOjE3NDAzMTQwNTF9.1S88hcKl45F2FUsMkwAkN0R8sLy061uyNbmJCkZOEBo'
``` 

Check
```bash
curl --location 'http://localhost:8080/api/auth/check' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlR5cGUiOiJhY2Nlc3MiLCJleHAiOjE3NDAzMTQzMjIsImlhdCI6MTc0MDMxNDI2Mn0.mRgoCq1JjkoKelTBmQb6xvUd5vmV8tjEtIn7r6UgnZo'
```

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->
