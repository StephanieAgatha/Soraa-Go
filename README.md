
# Sora Go

SDK Boilerplate for restful api

#### Installation

`
go get github.com/StephanieAgatha/Soraa-Go
`

#### Copy whole .env.example and create a new .env


#### Paseto & JWT expire token (Hour Format) ex :
```go
jsonToken := paseto.JSONToken{
		Issuer:     "Sora Project",
		Subject:    "Abrakadabra",
		Expiration: expire,
		IssuedAt:   now,
	}
```





## Got em Middleware

#### Dummy endpoint (middleware protected)

```http
  POST /new 
```

| Description                |
| :------------------------- |
| **Required**. Paseto / JWT Token |



#### Stack
- Gin
- JWT
- Paseto
- Zap (Logging)
- Redis Cache
 

