# BiblioExchange

## Setup

### Setup the K8S cluster
To setup the K8S cluster just run:
```
bash ./deploy.sh
```

### Migrations
To execute the database migrations run:
```
(cd migrations; goose postgres "postgres://postgres:admin@localhost:5432/library?sslmode=disable" up)
```

If you don`t have goose, [here](https://github.com/pressly/goose#install) there is instructions how to install it.

## Database Access
[Database tables diagram](https://dbdiagram.io/d/65897b4289dea6279984649d)

First make sure that minikube tunnel is active:

```
minikube tunnel --bind-address='127.0.0.1'
```

Then go to [PgAdmin Console](http://localhost:5050/browser/)  
- email: admin@admin.com  
- password: admin


After that register server:  
- host name: db  
- username: postgres  
- password: admin

## User Management Access

First make sure that minikube tunnel is active:

```
minikube tunnel --bind-address='127.0.0.1'
```

Then go to [Keycloak Console](http://localhost:8080/realms/master/protocol/openid-connect/auth?client_id=security-admin-console&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fadmin%2Fmaster%2Fconsole%2F&state=06ebf6eb-72bc-4da9-8684-d1803f60cf02&response_mode=fragment&response_type=code&scope=openid&nonce=58ccfaa8-1387-4e2a-8f28-1b42d61f61b7&code_challenge=Dd9ifIN9daS7n31jq0ziHjeoJaBrHDc_8OsqJjlHkLg&code_challenge_method=S256)  
- username: admin  
- password: admin

Create client
## Object Store Access

First make sure that minikube tunnel is active:

```
minikube tunnel --bind-address='127.0.0.1'
```

Then go to [Minio Console](http://localhost:9090/login)  
- username: minioadmin  
- password: minioadmin

Create bucket called bucket "bucket"
And Access Keys

## Run services

For now service must be run locally with:

```
go run main.go
```
OR below for rerecognition service
```
python3 main.py
```

The golang services have config file in config folder for configuration.
In the future configuration will be done by .env files and load environment function.
