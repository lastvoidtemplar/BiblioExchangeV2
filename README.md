<style>
    p{
        margin-bottom:0
    }
</style>


# BiblioExchange

## Setup

To setup the K8S cluster just run:
```
bash ./deploy.sh
```

## Database Access

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

