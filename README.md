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
&nbsp;&nbsp;&nbsp;&nbsp;● email: admin@admin.com  
&nbsp;&nbsp;&nbsp;&nbsp;● password: a$$word


After that register server:  
&nbsp;&nbsp;&nbsp;&nbsp;● host name: db  
&nbsp;&nbsp;&nbsp;&nbsp;● username: postgres  
&nbsp;&nbsp;&nbsp;&nbsp;● password: a$$word


