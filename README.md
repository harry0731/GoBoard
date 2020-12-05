# GoBoard
A message board built in Golang

This project is for Go practicing and guided by [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) and [Test-driven Development of Go Web Applications with Gin](https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin) then modified by me.  

Install Dependency  
```sh
$ go mod download
```  

Run Database
```sh
$ docker-compose up -d
```
Remember to create a database named GoBoard in the PostgreSQL server

Run
```sh
$ go run main.go
```

Test
```sh
$ go test -v tests/
```

# Database

## Environments
This Compose file contains the following environment variables:

* `POSTGRES_USER` the default value is **postgres**
* `POSTGRES_PASSWORD` the default value is **changeme**
* `PGADMIN_PORT` the default value is **5050**
* `PGADMIN_DEFAULT_EMAIL` the default value is **pgadmin4@pgadmin.org**
* `PGADMIN_DEFAULT_PASSWORD` the default value is **admin**

## Access to postgres: 
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** changeme (as a default)

## Access to PgAdmin: 
* **URL:** `http://localhost:5050`
* **Username:** pgadmin4@pgadmin.org (as a default)
* **Password:** admin (as a default)

## Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`, by default: `postgres`
* **Password** as `POSTGRES_PASSWORD`, by default `changeme`

TODO:  
- [x] Display the list of all articles on the home page
- [x] Display a single article on its own page
- [x] Users register with a username and a password
- [x] Login
- [x] Log out
- [x] Create new articles (logged in users only)
- [x] Use PostgreSQL as database
- [ ] Mock GORM DB to test
- [ ] Deploy on cloud